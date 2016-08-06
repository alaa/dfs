package splitter

import (
	"bytes"
	"fmt"
	"os"
)

func MergeParts(Parts [][]byte) []byte {
	var file []byte
	for _, part := range Parts {
		for _, data := range part {
			file = append(file, data)
		}
	}
	return file
}

func WriteParts(Parts [][]byte) error {
	var buf bytes.Buffer
	for i, part := range Parts {
		f, err := newFile(fmt.Sprintf("%d", i))
		defer f.Close()
		if err != nil {
			return err
		}
		buf.Write(part)
		buf.WriteTo(f)
	}
	return nil
}

func Split(bytes []byte, chunkSize int) [][]byte {
	var Parts [][]byte
	for i := 0; i <= len(bytes); i += chunkSize {
		chunk := bytes[i:min(i+chunkSize, len(bytes))]
		Parts = append(Parts, chunk)
	}
	return Parts
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
