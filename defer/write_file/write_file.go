package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if err := writeToFile("dummy.txt", "It was the best of times, it was the worst of times"); err != nil {
		log.Fatalf("%v\n", err)
	}
}

func writeToFile(filename string, text string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.WriteString(file, text)
	if err != nil {
		return err
	}

	// hacky, but needed in case file.Close() itself raises an error
	return file.Close()
}