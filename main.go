package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"./journal"
	"./splitter"
)

func main() {
	chunkSize := (1024 * 1) // 1 KB

	bytes, err := ioutil.ReadFile("./test")
	if err != nil {
		log.Panic("could not read file")
	}

	parts := splitter.Split(bytes, chunkSize)
	ids, err := splitter.WriteParts(parts)
	if err != nil {
		log.Println(err)
	}
	metadata := journal.Register("test", ids)
	fmt.Println("%v", metadata)

	// Merge all parts again
	var file [][]byte
	for _, c := range metadata.Parts {
		p, err := ioutil.ReadFile("parts/" + c)
		if err != nil {
			log.Panic("could not read file")
		}
		file = append(file, p)
	}
	f := splitter.MergeParts(&file)
	fmt.Printf("assembled file size is %d \n", len(*f))

	fh, err := os.Create("assembled/" + metadata.Filename)
	if err != nil {
		log.Printf("could not create assembled file with err: %s", err)
	}
	fh.Write(*f)
}
