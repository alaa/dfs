package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"./splitter"
)

func main() {
	chunkSize := (1024 * 10)

	bytes, err := ioutil.ReadFile("./test")
	if err != nil {
		log.Panic("could not read file")
	}

	parts := splitter.Split(bytes, chunkSize)
	if err = splitter.WriteParts(parts); err != nil {
		log.Println(err)
	}

	p1, err := ioutil.ReadFile("./0")
	if err != nil {
		log.Panic("could not read file")
	}
	p2, err := ioutil.ReadFile("./1")
	if err != nil {
		log.Panic("could not read file")
	}

	var file [][]byte
	file = append(file, p1)
	file = append(file, p2)
	f := splitter.MergeParts(file)
	fmt.Printf("assembled file size is %d \n", len(f))

	fh, err := os.Create("assembled")
	if err != nil {
		log.Printf("could not create assembled file with err: %s", err)
	}
	fh.Write(f)
}
