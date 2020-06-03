package main

import (
	"log"
	"fmt"
)

func main() {
	fileWriter := New(Config{NewIoutilPkg()})
	fileWriter.WriteToFile("Hello World", "/tmp", "message-phase-3-approach-2.txt")
}

type Config struct {
	IoutilPkg IoutilPkg
}

type FileWriter interface {
	WriteToFile(data, dir, location string) error
}

type FileWriterImpl struct {
	ioutilPkg IoutilPkg
}

func New(c Config) FileWriterImpl{
	return FileWriterImpl{c.IoutilPkg}
}

func (fwi *FileWriterImpl) WriteToFile(data, dir, location string) error {
	_, err := fwi.ioutilPkg.ReadDir(dir)
	if err != nil {
		log.Println("Sorry:", err)
		return err
	}
	outputFile := fmt.Sprintf("%s/%s", dir, location)
	err = fwi.ioutilPkg.WriteFile(outputFile, []byte(data), 0600)
	if err != nil {
		log.Println("Sorry:", err)
		return err
	}
	return nil
}
