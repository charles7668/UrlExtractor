package fileutil

import (
	"os"
	"testing"
)

func TestCheckFileExist(t *testing.T) {
	existFile := "existFile.txt"
	nonExistFile := "nonExistFile.txt"
	tempDir := t.TempDir()

	// create the existFile and keep nonExistFile non exist
	file, err := os.Create(tempDir + "/" + existFile)
	if err != nil {
		t.Error(err)
	}
	defer CloseWithHandleError(file)

	// check
	if !CheckFileExist(tempDir + "/" + existFile) {
		t.Error(existFile + " File should exist")
	}

	if CheckFileExist(tempDir + "/" + nonExistFile) {
		t.Error(nonExistFile + " File should not exist")
	}

	// now create the nonExistFile and check again
	fileNonExist, err := os.Create(tempDir + "/" + nonExistFile)
	if err != nil {
		t.Error(err)
	}
	defer CloseWithHandleError(fileNonExist)

	// check nonExistFile is exist
	if !CheckFileExist(tempDir + "/" + nonExistFile) {
		t.Error("in second test , " + nonExistFile + " File should exist")
	}
}
