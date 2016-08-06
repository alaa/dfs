package splitter

import (
	"bytes"
	"fmt"
	"os"
)

type part struct {
	data []byte
}

func WriteParts(parts []part) error {
	var buf bytes.Buffer
	for i, part := range parts {
		f, err := newFile(fmt.Sprintf("%d", i))
		defer f.Close()
		if err != nil {
			return err
		}
		buf.Write(part.data)
		buf.WriteTo(f)
	}
	return nil
}

func Split(bytes []byte, chunkSize int) []part {
	var parts []part
	for i := 0; i <= len(bytes); i += chunkSize {
		chunk := bytes[i:min(i+chunkSize, len(bytes))]
		parts = append(parts, part{data: chunk})
	}
	return parts
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
