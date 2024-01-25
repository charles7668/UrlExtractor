package fileutil

import (
	"fmt"
	"io"
	"os"
)

// CheckFileExist checks if the file exist
func CheckFileExist(path string) bool {
	// Check if the file path is valid
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

// CloseWithHandleError close file with error handling
func CloseWithHandleError(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("Error closing file : ", err)
	}
}

// ReadFile reads file and return the content as slice of string
func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer CloseWithHandleError(file)

	var lines []string
	var line string
	// read the file line by line
	for {
		_, err := fmt.Fscanf(file, "%s\n", &line)
		if err != nil {
			if err.Error() == "unexpected newline" {
				continue
			}
			if err == io.EOF {
				break
			}
			return []string{}, err
		}
		lines = append(lines, line)
	}
	return lines, nil
}
