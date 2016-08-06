package main

import (
	"io/ioutil"
	"log"

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
}
