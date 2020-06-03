package main

import (
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

type WriteFileFunc func(filename string, data []byte, perm os.FileMode) error
type ReadDirFunc func(dirname string) ([]os.FileInfo, error)

var (
	ioutil_WriteFile WriteFileFunc = ioutil.WriteFile
	ioutil_ReadDir ReadDirFunc = ioutil.ReadDir
)

func main() {
	WriteToFile("Hello World", "/tmp", "message-phase-1-and-2.txt")
}

func WriteToFile(data, dir, location string) error {
	_, err := ioutil_ReadDir(dir)
	if err != nil {
		log.Println("Sorry:", err)
		return err
	}
	outputFile := fmt.Sprintf("%s/%s", dir, location)
	err = ioutil_WriteFile(outputFile, []byte(data), 0600)
	if err != nil {
		log.Println("Sorry:", err)
		return err
	}
	return nil
}
