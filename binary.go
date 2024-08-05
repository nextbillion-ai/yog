package yog

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
)

func decodeChunk(input []byte) ([]uint32, []uint32, error) {
	var distances = make([]uint32, len(input)/8)
	var durations = make([]uint32, len(input)/8)

	for i := 0; i < len(input); i += 8 {
		buffer0 := bytes.NewBuffer(input[i : i+4])
		buffer1 := bytes.NewBuffer(input[i+4 : i+8])

		err := binary.Read(buffer0, binary.LittleEndian, &durations[i/8])
		if err != nil {
			return nil, nil, err
		}

		err = binary.Read(buffer1, binary.LittleEndian, &distances[i/8])
		if err != nil {
			return nil, nil, err
		}
	}
	return durations, distances, nil
}

func encodeChunk(durations, distances []uint32) ([]byte, error) {
	var binBuf bytes.Buffer

	for i := range durations {
		err := binary.Write(&binBuf, binary.LittleEndian, durations[i])
		if err != nil {
			return nil, err
		}
		err = binary.Write(&binBuf, binary.LittleEndian, distances[i])
		if err != nil {
			return nil, err
		}
	}

	return binBuf.Bytes(), nil
}

func decode(input []byte) (uint32, uint32, error) {
	buffer0 := bytes.NewBuffer(input[0:4])
	buffer1 := bytes.NewBuffer(input[4:8])
	var duration, distance uint32
	err := binary.Read(buffer0, binary.LittleEndian, &duration)
	if err != nil {
		return 0, 0, err
	}

	err = binary.Read(buffer1, binary.LittleEndian, &distance)
	if err != nil {
		return 0, 0, err
	}
	return duration, distance, nil
}

func encode(duration, distance uint32) ([]byte, error) {
	var binBuf bytes.Buffer
	err := binary.Write(&binBuf, binary.LittleEndian, duration)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&binBuf, binary.LittleEndian, distance)
	if err != nil {
		return nil, err
	}
	return binBuf.Bytes(), nil
}

// Encode MatrixInfo Result -> binary
// i = o * d.length + d
func EncodeNumber(data string) ([]byte, error) {

	var resp MatrixData
	err := json.Unmarshal([]byte(data), &resp)
	if err != nil {
		return nil, err
	}

	result := ""
	for _, row := range resp.Rows {
		rowStr := "|"
		for _, e := range row.Elements {
			rowStr = preByteConcat(rowStr, fmt.Sprintf("%v,%v|", e.Duration.Value, e.Distance.Value))
		}
		result = preByteConcat(result, rowStr)
		result = preByteConcat(result, "\n")
	}
	return []byte(result), nil
}

func preByteConcat(str1, str2 string) string {
	buf := []byte(str1)
	buf = append(buf, str2...)
	return string(buf)
}

// Encode MatrixInfo Result -> binary
// i = o * d.length + d
func Encode(data string) ([]byte, error) {

	var resp MatrixData
	err := json.Unmarshal([]byte(data), &resp)
	if err != nil {
		return nil, err
	}

	header, err := encode(uint32(len(resp.Rows)), uint32(len(resp.Rows[0].Elements)))
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
			chunk, err := encode(element.Duration.Value, element.Distance.Value)
			if err != nil {
				return nil, err
			}
			res = append(res, chunk...)
		}
	}
	return res, nil
}

// Decode binary -> MatrixInfo Result
func Decode(bin []byte) (string, error) {

	// check data
	if len(bin)%8 > 0 {
		return "", errors.New("illegal binary data format")
	}

	//read header
	sourceLength, destinationLength, err := decode(bin[0:8])
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
		duration, distance, err := decode(bin[i : i+8])
		if err != nil {
			return "", errors.New("binary decode failed")
		}

		n := uint32(i / 8)
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
