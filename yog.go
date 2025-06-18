package yog

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	Uint32Binary int = iota
	Int32Binary
)

type Yog struct {
	taskID     string // 方舟接口返回的
	path       string // 我们把 gsg 文件下载到哪里
	apiKey     string // google key
	offset     int    // 记录读数据进度
	outputAddr string
	taskMeta   TaskMeta
	storage    ICloudStorage
}

func removeGsorS3Prefix(outputAddr string) string {
	outputAddr = strings.TrimPrefix(outputAddr, "gs://")
	outputAddr = strings.TrimPrefix(outputAddr, "s3://")
	return outputAddr
}

func NewInternal(taskID string, path string, outputAddr string, apiKeyPath string) *Yog {
	outputAddrWithoutPrefix := removeGsorS3Prefix(outputAddr)
	arrays := strings.Split(outputAddrWithoutPrefix, "/")
	return &Yog{
		taskID:  taskID,
		path:    path,
		apiKey:  apiKeyPath,
		offset:  0,
		storage: NewGCSClient(arrays[0], strings.Join(arrays[1:], "/"), apiKeyPath, outputAddr),
	}
}

func New(taskID string, path string, outputAddr string) *Yog {
	if path != "" && path[len(path)-1] != '/' {
		path = path + "/"
	}
	outputAddr = "https://" + removeGsorS3Prefix(outputAddr)
	return &Yog{
		taskID:  taskID,
		path:    path,
		apiKey:  "",
		offset:  0,
		storage: NewHTTPGetClient(outputAddr),
	}
}

// Load download file & check meta
func (y *Yog) Load() error {

	if y.path[len(y.path)-1] != '/' {
		y.path = y.path + "/"
	}

	err := os.MkdirAll(y.path, os.ModePerm)
	if err != nil {
		return err
	}

	ctx := context.TODO()

	// download meta
	err = y.storage.DownloadSingleFile(ctx, META_FILE_NAME, y.path)
	if err != nil {
		return fmt.Errorf("meta file download failed filename:%v error:%v", META_FILE_NAME, err)
	}

	// download chunk by meta content
	metaContent, err := bufferReaderFile(y.path+META_FILE_NAME, 10240)
	if err != nil {
		return err
	}

	if len(metaContent) == 0 || metaContent[0] != '{' {
		return errors.New("meta file not exist in remote storage")
	}

	// meta.json unmarshal TaskMeta
	err = json.Unmarshal(metaContent, &y.taskMeta)
	if err != nil {
		return fmt.Errorf("meta file illegal file format %v", err)
	}

	for chunkFileName, chunkInfo := range y.taskMeta.Index {
		err := y.storage.DownloadSingleFile(ctx, chunkFileName, y.path)
		if err != nil {
			if err.Error() == "file not exist" {
				chunkInfo.Status = FFailed
				y.taskMeta.Index[chunkFileName] = chunkInfo
				continue
			}
			return fmt.Errorf("chunk file download failed filename:%v error:%v", chunkFileName, err)
		}
		chunkInfo.Status = FSucceeded
		y.taskMeta.Index[chunkFileName] = chunkInfo
	}

	return y.ReloadMeta()
}

func (y *Yog) Check() error {

	// check meta.json exist
	metaFilePath := y.path + "/" + META_FILE_NAME
	if !fileExists(metaFilePath) {
		return errors.New("meta file does not exist")
	}

	// download chunk by meta content
	metaContent, err := bufferReaderFile(y.path+META_FILE_NAME, 10240)
	if err != nil {
		return err
	}

	if len(metaContent) == 0 || metaContent[0] != '{' {
		return errors.New("meta file not exist in remote storage")
	}

	// meta.json unmarshal TaskMeta
	err = json.Unmarshal(metaContent, &y.taskMeta)
	if err != nil {
		return fmt.Errorf("meta file illegal file format %v", err)
	}

	// check meta data
	if len(y.taskMeta.Index) == 0 {
		return errors.New("meta file does not include index")
	}

	for chunkName, chunkInfo := range y.taskMeta.Index {
		chunkPath := y.path + "/" + chunkName
		if !fileExists(chunkPath) {
			chunkInfo.Status = FFailed
		} else {
			chunkInfo.Status = FSucceeded
		}
		y.taskMeta.Index[chunkName] = chunkInfo
	}

	return nil
}

// ReloadMeta download file & check meta
func (y *Yog) ReloadMeta() error {

	if len(y.taskMeta.Index) == 0 {
		err := y.Check()
		if err != nil {
			return err
		}
	}

	oc, dc := 0, 0
	for _, data := range y.taskMeta.Index {
		data.Offset = []int{}
		for _, o := range data.Origin {
			if oc < o {
				oc = o
			}
		}
		for _, d := range data.Destination {
			if dc < d {
				dc = d
			}
		}
	}

	y.taskMeta.MatrixInfo.OriginCount = oc + 1
	y.taskMeta.MatrixInfo.DestinationCount = dc + 1

	for i, data := range y.taskMeta.Index {
		for _, o := range data.Origin {
			for _, d := range data.Destination {
				data.Offset = append(data.Offset, o*y.taskMeta.MatrixInfo.DestinationCount+d)
			}
		}
		y.taskMeta.Index[i] = data
	}

	// check meta data
	if len(y.taskMeta.Index) == 0 {
		return errors.New("meta file does not include index")
	}

	return nil
}

type page struct {
	status        FileState
	file          string
	offset        int64
	length        int64
	originalIndex IndexItem
	start         int
	end           int
	version       int // 记录版本
}

func (p *page) read() ([]byte, error) {

	if p.status == FFailed {
		return make([]byte, p.length), nil
	}

	file, err := os.OpenFile(p.file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = file.Seek(p.offset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	bufReader := bufio.NewReader(file)
	block := make([]byte, p.length)
	_, err = bufReader.Read(block)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (y *Yog) Reset() {
	y.offset = 0
}

// ReadChunk read a chunk of data in order
// i = o * d.length + d
func (y *Yog) ReadChunk(chunkSize int) (durations []int32, distances []int32, err error) {

	if y.offset+chunkSize > y.taskMeta.MatrixInfo.OriginCount*y.taskMeta.MatrixInfo.DestinationCount {
		chunkSize = y.taskMeta.MatrixInfo.OriginCount*y.taskMeta.MatrixInfo.DestinationCount - y.offset
		if chunkSize <= 0 {
			return nil, nil, io.EOF
		}
	}

	// read file chunk
	var pages = make([]page, 0)
	for indexName, indexData := range y.taskMeta.Index {

		startOffset, endOffset := findSubset(y.offset, y.offset+chunkSize-1, indexData.Offset)
		if startOffset < 0 || endOffset < 0 {
			continue
		}

		pages = append(pages, page{
			file:          y.path + "/" + indexName,
			offset:        int64(startOffset*8 + HEADER_LENGTH),
			length:        int64(endOffset-startOffset+1) * 8,
			originalIndex: indexData,
			start:         startOffset,
			end:           endOffset,
			status:        indexData.Status,
			version:       y.taskMeta.version,
		})
	}

	// allocate space
	durations = make([]int32, chunkSize)
	distances = make([]int32, chunkSize)

	// fill the data in order
	for _, p := range pages {
		data, err := p.read()
		if err != nil {
			return nil, nil, fmt.Errorf("read chunk failed %v", err)
		}

		var chunkDurations, chunkDistances []int32
		switch y.taskMeta.version {
		case Int32Binary:
			var serializer Int32BinarySerializer
			chunkDurations, chunkDistances, err = serializer.decodeChunk(data)
			if err != nil {
				return nil, nil, fmt.Errorf("decode binary chunk failed %v", err)
			}
		case Uint32Binary:
			var serializer Uint32BinarySerializer
			uintChunkDurations, uintchunkDistances, err := serializer.decodeChunk(data)
			if err != nil {
				return nil, nil, fmt.Errorf("decode binary chunk failed %v", err)
			}
			for _, duration := range uintChunkDurations {
				chunkDurations = append(chunkDurations, int32(duration))
			}
			for _, distance := range uintchunkDistances {
				chunkDistances = append(chunkDistances, int32(distance))
			}
		}

		for i := 0; i < p.end-p.start+1; i++ {
			ir := p.originalIndex.Offset[p.start+i] - y.offset
			durations[ir] = chunkDurations[i]
			distances[ir] = chunkDistances[i]
		}
	}

	y.offset += chunkSize
	return durations, distances, err
}

func findSubset(min, max int, array []int) (int, int) {
	begin, end := -1, -1

	for i, v := range array {
		if begin == -1 && v >= min && v <= max {
			begin = i
		}

		if v >= min && v <= max {
			end = i
		}
	}

	return begin, end
}

// Read the data with the specified index
func (y *Yog) Read(o, d int) (duration int32, distance int32, err error) {

	var p page
	oindex, dindex := -1, -1
	for indexName, indexData := range y.taskMeta.Index {
		for i, oo := range indexData.Origin {
			if o == oo {
				oindex = i
			}
		}
		for i, dd := range indexData.Destination {
			if d == dd {
				dindex = i
			}
		}

		if oindex >= 0 && dindex >= 0 {
			p.file = y.path + "/" + indexName
			p.offset = int64(oindex*len(indexData.Destination)+dindex)*8 + HEADER_LENGTH
			p.length = 8
			p.status = indexData.Status
			break
		}
		oindex, dindex = -1, -1
	}
	data, err := p.read()
	if err != nil {
		return 0, 0, fmt.Errorf("read chunk failed %v", err)
	}

	switch y.taskMeta.version {
	case Int32Binary:
		var serializer Int32BinarySerializer
		return serializer.decode(data)
	case Uint32Binary:
		var serializer Uint32BinarySerializer
		uDuration, uDistance, err := serializer.decode(data)
		return int32(uDuration), int32(uDistance), err
	default:
		return 0, 0, fmt.Errorf("version shouldn`t be %v", y.taskMeta.version)
	}
}

// MatrixInfo get MatrixInfo from meta
func (y *Yog) MatrixInfo() MatrixInfo {
	return y.taskMeta.MatrixInfo
}

// Clear remove task file
func (y *Yog) Clear() error {
	for name := range y.taskMeta.Index {
		chunkFilePath := y.path + name
		err := os.Remove(chunkFilePath)
		if err != nil {
			return fmt.Errorf("remove chunk file failed %v", err)
		}
	}

	indexFilePath := y.path + META_FILE_NAME
	err := os.Remove(indexFilePath)
	if err != nil {
		return fmt.Errorf("remove index file failed %v", err)
	}
	return nil
}

// encode MatrixInfo Result -> binary i = o * d.length + d
func Encode(data string) ([]byte, error) {

	var serializer Int32BinarySerializer

	var resp MatrixData
	err := json.Unmarshal([]byte(data), &resp)
	if err != nil {
		return nil, err
	}

	header, err := serializer.encode(int32(len(resp.Rows)), int32(len(resp.Rows[0].Elements)))
	if err != nil {
		return nil, err
	}

	// 将 MatrixData.Rows 转化成 binary
	res := make([]byte, 0)

	// add header
	res = append(res, header...)

	// source index
	for _, row := range resp.Rows {
		// destination index
		for _, element := range row.Elements {
			chunk, err := serializer.encode(element.Duration.Value, element.Distance.Value)
			if err != nil {
				return nil, err
			}
			res = append(res, chunk...)
		}
	}
	return res, nil
}

// decode binary -> MatrixInfo Result
func Decode(bin []byte) (string, error) {

	var serializer Int32BinarySerializer

	// check data
	if len(bin)%8 > 0 {
		return "", errors.New("illegal binary data format")
	}

	//read header
	sourceLength, destinationLength, err := serializer.decode(bin[0:8])
	if err != nil {
		return "", errors.New("binary header decode failed")
	}
	bin = bin[8:]

	// resize result
	var m MatrixData
	m.Rows = make([]MatrixRow, sourceLength)
	for i := range m.Rows {
		m.Rows[i].Elements = make([]MatrixElement, destinationLength)
	}

	// decode
	for i := 0; i < len(bin); i = i + 8 {
		duration, distance, err := serializer.decode(bin[i : i+8])
		if err != nil {
			return "", errors.New("binary decode failed")
		}

		n := int32(i / 8)
		var d = n % destinationLength
		var s = (n - d) / destinationLength
		//fmt.Println(fmt.Sprintf("s=%v, d=%v", s, d))

		m.Rows[s].Elements[d] = MatrixElement{
			Distance: Value{distance},
			Duration: Value{duration},
		}
	}

	marshal, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(marshal), nil
}
