package fileutil

import (
	"os"
)

// CheckPathValid checks if the file path is valid
func CheckPathValid(path string) bool {
	// Check if the file path is valid
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return true
		}
		return false
	}
	return true
}
