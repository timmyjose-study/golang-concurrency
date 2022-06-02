package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := writeToFile("dummy.txt", "It was the best of times, it was the worst of times"); err != nil {
		log.Fatalf("%v\n", err)
	}

	if err := copyFile("dummy.txt", "dummy_copy.txt"); err != nil {
		log.Fatalf("%v\n", err)
	}
}

func copyFile(source string, destination string) error {
	src, err := os.Open(source)
	if err != nil {
		return err
	}

	defer src.Close()

	dst, err := os.Create(destination)
	if err != nil {
		return err
	}

	defer dst.Close()

	n, err := io.Copy(dst, src)
	if err != nil {
		return err
	}

	fmt.Printf("Copied %d bytes from %s to %s\n", n, source, destination)

	if err := src.Close(); src != nil {
		return err
	}

	return dst.Close()
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

	return file.Close()
}