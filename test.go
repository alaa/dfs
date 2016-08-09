package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("/tmp/dat")
	check(err)
	reader := bufio.NewReader(f)
	buf := make([]byte, 100) // 1KB

	for {
		_, err := reader.Read(buf)
		if err == io.EOF {
			log.Println("Reached the end of file")
			break
		} else if err != nil {
			log.Fatalf("Could not read file chunk: %s", err)
		}

		stdout := bufio.NewWriter(os.Stdout)
		defer stdout.Flush()
		stdout.Write(buf)
		log.Println("---- Chunk")
	}

	defer f.Close()
}
