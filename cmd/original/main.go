package main

import (
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	WriteToFile("Hello World", "/tmp", "message.txt")
}

func WriteToFile(data, dir, location string) error {
	_, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println("Sorry:", err)
		return err
	}
	outputFile := fmt.Sprintf("%s/%s", dir, location)
	err = ioutil.WriteFile(outputFile, []byte(data), 0600)
	if err != nil {
		log.Println("Sorry:", err)
		return err
	}
	return nil
}
