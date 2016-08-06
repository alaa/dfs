package journal

type Metadata struct {
	Filename string
	Parts    []string
}

func Register(filename string, parts []string) *Metadata {
	return &Metadata{
		Filename: filename,
		Parts:    parts,
	}
}
