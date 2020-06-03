
package main

import (
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

type WriteFileFunc func(filename string, data []byte, perm os.FileMode) error
type ReadDirFunc func(dirname string) ([]os.FileInfo, error)

func main() {
	fileWriter := New(Config{WriteFile:ioutil.WriteFile, ReadDir:ioutil.ReadDir})
	fileWriter.WriteToFile("Hello World", "/tmp", "message-phase-3-approach-1.txt")
}

type FileWriter interface {
	WriteToFile(data, dir, location string) error
}

type Config struct {
	WriteFile WriteFileFunc
	ReadDir ReadDirFunc
}

type FileWriterImpl struct {
	WriteFile WriteFileFunc
	ReadDir ReadDirFunc
}

func New(c Config) FileWriterImpl{
	return FileWriterImpl{WriteFile:c.WriteFile, ReadDir: c.ReadDir}
}

func (fwi *FileWriterImpl) WriteToFile(data, dir, location string) error {
	_, err := fwi.ReadDir(dir)
	if err != nil {
		log.Println("Sorry:", err)
		return err
	}
	outputFile := fmt.Sprintf("%s/%s", dir, location)
	err = fwi.WriteFile(outputFile, []byte(data), 0600)
	if err != nil {
		log.Println("Sorry:", err)
		return err
	}
	return nil
}
