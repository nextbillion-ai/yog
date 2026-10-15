package yog

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

// The MVRO loader depends on this interface and signed values, including -1.
var _ interface {
	Load() error
	ReadChunk(int) ([]int32, []int32, error)
} = (*Yog)(nil)

type fixtureStorage map[string][]byte

func (s fixtureStorage) DownloadSingleFile(_ context.Context, name, path string) error {
	data, ok := s[name]
	if !ok {
		return fmt.Errorf("file not exist")
	}
	return os.WriteFile(path+name, data, 0600)
}

func fixtureValues(o, d int) (int32, int32) {
	values := []int32{0, -1, -2147483648, 2147483647, 123, 65536}
	return values[(o*6+d)%len(values)], int32(o*1000 + d)
}

// Write fixtures with the original binary.Write wire contract, independent of Encode.
func fixtureYog(t *testing.T, version int, unsorted, failed bool) *Yog {
	t.Helper()
	meta := TaskMeta{Version: version, Index: map[string]IndexItem{}}
	storage := fixtureStorage{}
	for i, origins := range [][]int{{0, 2}, {1, 3}} {
		for j, destinations := range [][]int{{0, 2, 4}, {1, 3, 5}} {
			if unsorted {
				origins = []int{origins[1], origins[0]}
				destinations = []int{destinations[2], destinations[0], destinations[1]}
			}
			name := fmt.Sprintf("%d-%d", i, j)
			meta.Index[name] = IndexItem{Origin: origins, Destination: destinations}
			var buf bytes.Buffer
			words := []int32{int32(len(origins)), int32(len(destinations))}
			for _, o := range origins {
				for _, d := range destinations {
					duration, distance := fixtureValues(o, d)
					words = append(words, duration, distance)
				}
			}
			if err := binary.Write(&buf, binary.LittleEndian, words); err != nil {
				t.Fatal(err)
			}
			if !failed || name != "1-1" {
				storage[name] = buf.Bytes()
			}
		}
	}
	data, err := json.Marshal(meta)
	if err != nil {
		t.Fatal(err)
	}
	storage[META_FILE_NAME] = data
	y := New("", t.TempDir(), "")
	y.storage = storage
	if err := y.Load(); err != nil {
		t.Fatal(err)
	}
	return y
}

func TestReadChunkCompatibility(t *testing.T) {
	for _, version := range []int{Uint32Binary, Int32Binary} {
		for _, unsorted := range []bool{false, true} {
			for _, failed := range []bool{false, true} {
				t.Run(fmt.Sprintf("v%d/unsorted=%v/failed=%v", version, unsorted, failed), func(t *testing.T) {
					y := fixtureYog(t, version, unsorted, failed)
					if got := y.MatrixInfo(); got != (MatrixInfo{4, 6}) {
						t.Fatalf("matrix info: %+v", got)
					}
					chunkSizes := []int{1, 5, 6, 7, 100}
					if unsorted {
						// Legacy partial reads of nonmonotonic offsets may panic.
						// This optimization leaves that behavior unchanged.
						chunkSizes = []int{24, 100}
					}
					for _, chunkSize := range chunkSizes {
						y.Reset()
						offset := 0
						for {
							durations, distances, err := y.ReadChunk(chunkSize)
							if err == io.EOF {
								break
							}
							if err != nil {
								t.Fatal(err)
							}
							for i := range durations {
								o, d := (offset+i)/6, (offset+i)%6
								duration, distance := fixtureValues(o, d)
								if failed && o%2 == 1 && d%2 == 1 {
									duration, distance = 0, 0
								}
								if durations[i] != duration || distances[i] != distance {
									t.Fatalf("chunk %d cell %d,%d: %d,%d; want %d,%d", chunkSize, o, d, durations[i], distances[i], duration, distance)
								}
								singleDuration, singleDistance, err := y.Read(o, d)
								if err != nil || singleDuration != duration || singleDistance != distance {
									t.Fatalf("Read(%d,%d): %d,%d,%v", o, d, singleDuration, singleDistance, err)
								}
							}
							offset += len(durations)
						}
						if offset != 24 {
							t.Fatalf("read %d cells", offset)
						}
					}
					if err := y.ReloadMeta(); err != nil {
						t.Fatal(err)
					}
					if err := y.ReloadMeta(); err != nil {
						t.Fatal(err)
					}
					for name, item := range y.taskMeta.Index {
						if len(item.Offset) != len(item.Origin)*len(item.Destination) {
							t.Fatalf("duplicated offsets for %s", name)
						}
					}
					y.Reset()
					if _, _, err := y.ReadChunk(24); err != nil {
						t.Fatal(err)
					}
				})
			}
		}
	}
}

func TestEncodingWireCompatibility(t *testing.T) {
	data := MatrixData{Rows: make([]MatrixRow, 4)}
	var golden bytes.Buffer
	if err := binary.Write(&golden, binary.LittleEndian, []int32{4, 6}); err != nil {
		t.Fatal(err)
	}
	for o := range data.Rows {
		for d := 0; d < 6; d++ {
			duration, distance := fixtureValues(o, d)
			data.Rows[o].Elements = append(data.Rows[o].Elements, MatrixElement{Duration: Value{duration}, Distance: Value{distance}})
			if err := binary.Write(&golden, binary.LittleEndian, []int32{duration, distance}); err != nil {
				t.Fatal(err)
			}
		}
	}
	raw, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	encoded, err := Encode(string(raw))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(encoded, golden.Bytes()) {
		t.Fatal("wire format changed")
	}
	decoded, err := Decode(golden.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	var got MatrixData
	if err := json.Unmarshal([]byte(decoded), &got); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, data) {
		t.Fatal("decoded matrix differs")
	}
}

// Retain the original short-read padding behavior; validation is out of scope.
func TestShortPageRetainsLegacyPadding(t *testing.T) {
	path := t.TempDir() + "/page"
	if err := os.WriteFile(path, []byte{1, 2, 3}, 0600); err != nil {
		t.Fatal(err)
	}
	p := page{file: path, length: 8, status: FSucceeded}
	got, err := p.read()
	if err != nil || !bytes.Equal(got, []byte{1, 2, 3, 0, 0, 0, 0, 0}) {
		t.Fatalf("legacy short read changed: %v, %v", got, err)
	}
}

func TestReadLargePage(t *testing.T) {
	path := t.TempDir() + "/page"
	want := bytes.Repeat([]byte{1, 2, 3, 4, 5, 6, 7, 8}, 8192)
	if err := os.WriteFile(path, append(make([]byte, 8), want...), 0600); err != nil {
		t.Fatal(err)
	}
	p := page{file: path, offset: 8, length: int64(len(want)), status: FSucceeded}
	got, err := p.read()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, want) {
		t.Fatal("incomplete page read")
	}
}

func TestSerializerWireCompatibility(t *testing.T) {
	durations := []int32{-1, 0, -2147483648, 2147483647}
	distances := []int32{0, 2147483647, -1, -2147483648}
	var golden bytes.Buffer
	for i := range durations {
		if err := binary.Write(&golden, binary.LittleEndian, []int32{durations[i], distances[i]}); err != nil {
			t.Fatal(err)
		}
	}
	for _, encoder := range []interface {
		encodeChunk([]int32, []int32) ([]byte, error)
	}{Int32BinarySerializer{}, Uint32BinarySerializer{}} {
		got, err := encoder.encodeChunk(durations, distances)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(got, golden.Bytes()) {
			t.Fatal("serializer wire format changed")
		}
	}
	gotDurations, gotDistances, err := (Int32BinarySerializer{}).decodeChunk(golden.Bytes())
	if err != nil || !reflect.DeepEqual(gotDurations, durations) || !reflect.DeepEqual(gotDistances, distances) {
		t.Fatal("signed serializer changed", err)
	}
	unsignedDurations, unsignedDistances, err := (Uint32BinarySerializer{}).decodeChunk(golden.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	for i := range durations {
		if unsignedDurations[i] != uint32(durations[i]) || unsignedDistances[i] != uint32(distances[i]) {
			t.Fatal("unsigned serializer changed")
		}
	}
}
