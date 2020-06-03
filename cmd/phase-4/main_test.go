package main

import (
	"testing"
	"fmt"
	"os"
	"bytes"

	testifyMocks "github.com/enmayo/refactoring-for-testability/internal/mocks/testify"
	goMockMocks "github.com/enmayo/refactoring-for-testability/internal/mocks/gomock"
	testify "github.com/stretchr/testify/mock"
	"github.com/golang/mock/gomock"
)

func TestWriteToFileWithTestify(t *testing.T) {
	// Arrange 1
	filename := "something.txt"
	dir := "/tmp"
	data := "I hope this works"
	var expectedFileMode os.FileMode = 0600

	testCases := []struct {
		message              string
		writeFileErr         error
	    readDirErr           error
		readDirResp          []os.FileInfo
		expectWriteFileFreq  int
		expectReadDirFreq    int
		expectedErr          error
	}{
		{
			message: "Successful Behavior",
			writeFileErr:nil,
			readDirErr:nil,
			readDirResp:nil,
			expectWriteFileFreq: 1,
			expectReadDirFreq: 1,
			expectedErr: nil,
		},
		{
			message: "Error on ioutil.ReadDir",
			writeFileErr:nil,
			readDirErr:fmt.Errorf("can't read dir"),
			readDirResp:nil,
			expectWriteFileFreq: 0,
			expectReadDirFreq: 1,
			expectedErr: fmt.Errorf("can't read dir"),
		},
		{
			message: "Error on ioutil.WriteFile",
			writeFileErr:fmt.Errorf("can't write file"),
			readDirErr:nil,
			readDirResp:nil,
			expectWriteFileFreq: 1,
			expectReadDirFreq: 1,
			expectedErr: fmt.Errorf("can't write file"),
		},
	}

	// Act and Assert
	for _, testCase := range testCases {
		t.Run(testCase.message, func(t *testing.T) {
			// Arrange 2
			mockIoutilpkg := &testifyMocks.IoutilPkg{}
			mockIoutilpkg.On("ReadDir", dir).Return(testCase.readDirResp, testCase.readDirErr).Times(testCase.expectReadDirFreq)
			
			// THERE HAS TO BE A BETTER WAY!
			if(testCase.expectWriteFileFreq > 0) {
				mockIoutilpkg.On("WriteFile",
					testify.MatchedBy(func(f string) bool { 
						return f == fmt.Sprintf("%s/%s", dir, filename)
					}),
					testify.MatchedBy(func(d []byte) bool { 
						return bytes.Compare(d, []byte(data)) == 0
					}),
					testify.MatchedBy(func(fileMode os.FileMode) bool { 
						return fileMode == expectedFileMode
					}),
				).Return(testCase.writeFileErr).Times(testCase.expectWriteFileFreq)
			}
			fileWriter := New(Config{mockIoutilpkg})

			// Act
			err := fileWriter.WriteToFile(data, dir, filename)

			// Assert
			if fmt.Sprint(err) != fmt.Sprint(testCase.expectedErr) {
				t.Logf("Expected error to be: %s but was: %s", testCase.expectedErr, err)
				t.FailNow()
			}
			mockIoutilpkg.AssertExpectations(t)
		})
	}
}

func TestWriteToFileWithGoMock(t *testing.T) {
	// Arrange 1
	filename := "something.txt"
	dir := "/tmp"
	data := "I hope this works"
	var expectedFileMode os.FileMode = 0600

	controller := gomock.NewController(t)
	defer controller.Finish()

	testCases := []struct {
		message              string
		writeFileErr         error
	    readDirErr           error
		readDirResp          []os.FileInfo
		expectWriteFileFreq  int
		expectReadDirFreq    int
		expectedErr          error
	}{
		{
			message: "Successful Behavior",
			writeFileErr:nil,
			readDirErr:nil,
			readDirResp:nil,
			expectWriteFileFreq: 1,
			expectReadDirFreq: 1,
			expectedErr: nil,
		},
		{
			message: "Error on ioutil.ReadDir",
			writeFileErr:nil,
			readDirErr:fmt.Errorf("can't read dir"),
			readDirResp:nil,
			expectWriteFileFreq: 0,
			expectReadDirFreq: 1,
			expectedErr: fmt.Errorf("can't read dir"),
		},
		{
			message: "Error on ioutil.WriteFile",
			writeFileErr:fmt.Errorf("can't write file"),
			readDirErr:nil,
			readDirResp:nil,
			expectWriteFileFreq: 1,
			expectReadDirFreq: 1,
			expectedErr: fmt.Errorf("can't write file"),
		},
	}

	// Act and Assert
	for _, testCase := range testCases {
		t.Run(testCase.message, func(t *testing.T) {
			// Arrange 2
			mockIoutilpkg := goMockMocks.NewMockIoutilPkg(controller)
			mockIoutilpkg.EXPECT().ReadDir(gomock.Eq(dir)).Return(testCase.readDirResp, testCase.readDirErr).Times(testCase.expectReadDirFreq)
			mockIoutilpkg.EXPECT().WriteFile(
				gomock.Eq(fmt.Sprintf("%s/%s", dir, filename)), 
				gomock.Eq([]byte(data)), 
				gomock.Eq(expectedFileMode),
			).Return(testCase.writeFileErr).Times(testCase.expectWriteFileFreq)

			fileWriter := New(Config{mockIoutilpkg})

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