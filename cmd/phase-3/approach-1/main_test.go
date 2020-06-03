package main

import (
	"testing"
	"fmt"
	"os"
)

func TestWriteToFile(t *testing.T) {
	// Arrange 1
	filename := "something.txt"
	dir := "/tmp"
	data := "I hope this works"

	testCases := []struct {
		message            string
		ioutil_ReadDir     ReadDirFunc
		ioutil_WriteFile   WriteFileFunc
		expectedErr        error
	}{
		{
			message: "Successful Behavior",
			ioutil_ReadDir: func(dirname string) ([]os.FileInfo, error) {
				return nil, nil
			},
			ioutil_WriteFile: func(filename string, data []byte, perm os.FileMode) error {
				return nil
			},
			expectedErr: nil,
		},
		{
			message: "Error on ioutil.ReadDir",
			ioutil_ReadDir: func(dirname string) ([]os.FileInfo, error) {
				return nil, fmt.Errorf("can't read dir")
			},
			ioutil_WriteFile: func(filename string, data []byte, perm os.FileMode) error {
				return nil
			},
			expectedErr: fmt.Errorf("can't read dir"),
		},
		{
			message: "Error on ioutil.WriteFile",
			ioutil_ReadDir: func(dirname string) ([]os.FileInfo, error) {
				return nil, nil
			},
			ioutil_WriteFile: func(filename string, data []byte, perm os.FileMode) error {
				return fmt.Errorf("can't write file")
			},
			expectedErr: fmt.Errorf("can't write file"),
		},
	}

	// Act and Assert
	for _, testCase := range testCases {
		t.Run(testCase.message, func(t *testing.T) {
			// Arrange 2
			fileWriter := New(Config{WriteFile:testCase.ioutil_WriteFile, ReadDir:testCase.ioutil_ReadDir})

			// Act
			err := fileWriter.WriteToFile(data, dir, filename)

			// Assert
			if fmt.Sprint(err) != fmt.Sprint(testCase.expectedErr) {
				t.Logf("Expected error to be: %s but was: %s", testCase.expectedErr, err)
				t.FailNow()
			}
		})
	}
}
