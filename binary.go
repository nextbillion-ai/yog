package yog

import (
	"bytes"
	"encoding/binary"
)

type Uint32BinarySerializer struct {
}

func (b Uint32BinarySerializer) decodeChunk(input []byte) ([]uint32, []uint32, error) {
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

func (b Uint32BinarySerializer) encodeChunk(durations, distances []int32) ([]byte, error) {
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

func (b Uint32BinarySerializer) decode(input []byte) (uint32, uint32, error) {
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

func (b Uint32BinarySerializer) encode(duration, distance int32) ([]byte, error) {
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

type Int32BinarySerializer struct {
}

func (b Int32BinarySerializer) decodeChunk(input []byte) ([]int32, []int32, error) {
	var distances = make([]int32, len(input)/8)
	var durations = make([]int32, len(input)/8)

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

func (b Int32BinarySerializer) encodeChunk(durations, distances []int32) ([]byte, error) {
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

func (b Int32BinarySerializer) decode(input []byte) (int32, int32, error) {
	buffer0 := bytes.NewBuffer(input[0:4])
	buffer1 := bytes.NewBuffer(input[4:8])
	var duration, distance int32
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

func (b Int32BinarySerializer) encode(duration, distance int32) ([]byte, error) {
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
