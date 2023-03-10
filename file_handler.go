package filehandler

import (
	"fmt"
	"io/ioutil"
	"os"
)

func OpenFile(filepath string) (*os.File, error) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	return file, nil
}

func CloseFile(file *os.File) error {
	err := file.Close()
	if err != nil {
		return fmt.Errorf("failed to close file: %v", err)
	}
	return nil
}

func ReadFile(filepath string) ([]byte, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file info: %v", err)
	}

	var data []byte
	if fileInfo.Size() == 0 {
		data = []byte{}
	} else {
		data, err = ioutil.ReadFile(filepath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %v", err)
		}
	}

	return data, nil
}

func WriteToFile(file *os.File, data []byte) error {
	_, err := file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	return nil
}
