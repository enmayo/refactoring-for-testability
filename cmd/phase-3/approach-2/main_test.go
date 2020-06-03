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
		message        string
		ioutilpkg      IoutilPkg
		expectedErr    error
	}{
		{
			message: "Successful Behavior",
			ioutilpkg: &mockIoutilPkg{writeFileErr:nil, readDirErr:nil, readDirResp:nil},
			expectedErr: nil,
		},
		{
			message: "Error on ioutil.ReadDir",
			ioutilpkg: &mockIoutilPkg{writeFileErr:nil, readDirErr:fmt.Errorf("can't read dir"), readDirResp:nil},
			expectedErr: fmt.Errorf("can't read dir"),
		},
		{
			message: "Error on ioutil.WriteFile",
			ioutilpkg: &mockIoutilPkg{writeFileErr:fmt.Errorf("can't write file"), readDirErr:nil, readDirResp:nil},
			expectedErr: fmt.Errorf("can't write file"),
		},
	}

	// Act and Assert
	for _, testCase := range testCases {
		t.Run(testCase.message, func(t *testing.T) {
			// Arrange 2
			fileWriter := New(Config{testCase.ioutilpkg})

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

type mockIoutilPkg struct {
	writeFileErr error
	readDirErr error
	readDirResp []os.FileInfo
}

func (mioutilpkg *mockIoutilPkg) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return mioutilpkg.writeFileErr
}

func (mioutilpkg *mockIoutilPkg) ReadDir(dirname string) ([]os.FileInfo, error) {
	return mioutilpkg.readDirResp, mioutilpkg.readDirErr
}