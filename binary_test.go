package yog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sedecodeChunk(t *testing.T) {
	durations := []uint32{1, 2, 4, 512, 5142, 123, 14}
	distances := []uint32{5142, 123, 14, 1, 2, 4, 512}

	binary, err := encodeChunk(durations, distances)
	durations1, distances1, err := decodeChunk(binary)
	if err != nil {
		return
	}

	assert.Equal(t, durations, durations1)
	assert.Equal(t, distances, distances1)
}
