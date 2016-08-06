package splitter

import (
	"bytes"
	"os"

	"github.com/twinj/uuid"
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

func WriteParts(Parts [][]byte) ([]string, error) {
	var buf bytes.Buffer
	var ids []string
	for _, part := range Parts {
		id := uuid.NewV4().String()
		ids = append(ids, id)
		f, err := newFile("parts/" + id)
		defer f.Close()
		if err != nil {
			return nil, err
		}
		buf.Write(part)
		buf.WriteTo(f)
	}
	return ids, nil
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
