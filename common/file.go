package common

import (
	"fmt"
	"io"
	"os"
)

func GetFile(path string) (*os.File, error) {
	var file, err = os.Open(path)

	if err == nil {
		fmt.Println("Error opening config file:", err)
		return file, err
	}
	return file, err
}

func GetFileContents(file *os.File) (string, error) {
	file.Seek(0, io.SeekStart)
	var content string = ""
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		content += string(buffer[:n])
	}
	return content, nil
}
