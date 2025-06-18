package yog

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_sedecodeChunk(t *testing.T) {
	durations := []int32{-1, 2, 4, 512, 5142, 123, 14}
	distances := []int32{5142, 123, 14, 1, 2, 4, 512}

	var binaryc Int32BinarySerializer
	binary, err := binaryc.encodeChunk(durations, distances)
	durations1, distances1, err := binaryc.decodeChunk(binary)
	if err != nil {
		return
	}

	assert.Equal(t, durations, durations1)
	assert.Equal(t, distances, distances1)
}

func Test_random_sedecodeChunk(t *testing.T) {

	for i := 0; i < 10; i++ {

		var distances, durations []int32
		for ii := 0; ii < 10; ii++ {
			durations = append(durations, int32(rand.Intn(2*1000000)-1000000))
			distances = append(distances, int32(rand.Intn(2*1000000)-1000000))
		}

		var serializer Int32BinarySerializer
		binary, err := serializer.encodeChunk(durations, distances)
		durations1, distances1, err := serializer.decodeChunk(binary)
		if err != nil {
			return
		}
		assert.Equal(t, durations, durations1)
		assert.Equal(t, distances, distances1)
	}

}
