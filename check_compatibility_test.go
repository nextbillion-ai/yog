package yog

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// Use only public types/methods and the legacy wire encoding so the same test
// can be run unchanged against v0.1.12 as well as the optimized implementation.
func TestCheckReloadCompatibility(t *testing.T) {
	for _, version := range []int{Uint32Binary, Int32Binary} {
		for _, chunkSize := range []int{1, 2, 4} {
			t.Run(fmt.Sprintf("version=%d/chunk=%d", version, chunkSize), func(t *testing.T) {
				dir := t.TempDir()
				var records bytes.Buffer
				distances := []int32{-1, 0, -2147483648, 2147483647}
				if err := binary.Write(&records, binary.LittleEndian, []int32{1, 4}); err != nil {
					t.Fatal(err)
				}
				for i, distance := range distances {
					if err := binary.Write(&records, binary.LittleEndian, []int32{int32(100 + i), distance}); err != nil {
						t.Fatal(err)
					}
				}
				if err := os.WriteFile(filepath.Join(dir, "0"), records.Bytes(), 0600); err != nil {
					t.Fatal(err)
				}
				writeMeta := func(destinations []int, snapshot bool) {
					t.Helper()
					item := IndexItem{Origin: []int{0}, Destination: destinations}
					meta := TaskMeta{Version: version, Index: map[string]IndexItem{}}
					if snapshot {
						item.Offset = append([]int(nil), destinations...)
						meta.MatrixInfo = MatrixInfo{OriginCount: 1, DestinationCount: 4}
					}
					meta.Index["0"] = item
					data, err := json.Marshal(meta)
					if err != nil {
						t.Fatal(err)
					}
					if err := os.WriteFile(filepath.Join(dir, META_FILE_NAME), data, 0600); err != nil {
						t.Fatal(err)
					}
				}
				writeMeta([]int{0, 1, 2, 3}, false)
				y := New("check-compatibility", dir, "")
				if err := y.ReloadMeta(); err != nil {
					t.Fatal(err)
				}
				for stage, destinations := range [][]int{{3, 2, 1, 0}, {0, 1, 2, 3}, {3, 2, 1, 0}} {
					writeMeta(destinations, true)
					if err := y.Check(); err != nil {
						t.Fatal(err)
					}
					wantDurations, wantDistances := make([]int32, 4), make([]int32, 4)
					for position, destination := range destinations {
						wantDurations[destination] = int32(100 + position)
						wantDistances[destination] = distances[position]
					}
					// Check must make the new offsets readable without another
					// ReloadMeta call to rebuild the acceleration state.
					y.Reset()
					var gotDurations, gotDistances []int32
					for read := 0; read < 4/chunkSize; read++ {
						durations, distances, err := y.ReadChunk(chunkSize)
						if err != nil {
							t.Fatal(err)
						}
						gotDurations = append(gotDurations, durations...)
						gotDistances = append(gotDistances, distances...)
					}
					if !reflect.DeepEqual(gotDurations, wantDurations) || !reflect.DeepEqual(gotDistances, wantDistances) {
						t.Errorf("stage=%d: got %v/%v; want %v/%v", stage, gotDurations, gotDistances, wantDurations, wantDistances)
					}
					if _, _, err := y.ReadChunk(chunkSize); err != io.EOF {
						t.Fatalf("expected EOF after reloaded matrix, got %v", err)
					}
				}
			})
		}
	}
}
