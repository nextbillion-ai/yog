package yog

import (
	"bufio"
	"io"
	"os"
)

func fileExists(filepath string) bool {
	if _, err := os.Stat(filepath); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

func fileSize(filePath string) int {
	fi, err := os.Stat(filePath)
	if err == nil {
		return int(fi.Size())
	}
	return -1
}

func bufferReaderFile(path string, bufferSize int) ([]byte, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	length := 0
	fileContent := make([]byte, 0)
	b := make([]byte, bufferSize)
	for {
		n, err := file.Read(b)
		length += n
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		fileContent = append(fileContent, b...)
	}
	if err != nil {
		return nil, err
	}
	return fileContent[:length], nil
}

func readByteByOffsetLength(fileName string, offset, length int64) ([]byte, error) {

	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	bufReader := bufio.NewReader(file)
	block := make([]byte, length)
	_, err = bufReader.Read(block)
	if err != nil {
		return nil, err
	}

	return block, nil
}
