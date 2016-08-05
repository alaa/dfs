package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var buf bytes.Buffer
	chunkSize := 100

	bytes, err := ioutil.ReadFile("./test")
	if err != nil {
		log.Panic("could not read file")
	}

	for i := 0; i <= len(bytes); i += chunkSize {
		part := bytes[i:min(i+chunkSize, len(bytes))]
		partName := fmt.Sprintf("part-%d", i)

		f, err := newFile(partName)
		defer f.Close()
		if err != nil {
			log.Printf("Could not create part file %s", part)
		}

		buf.Write(part)
		buf.WriteTo(f)
	}
}

func newFile(filename string) (*os.File, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
