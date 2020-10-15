package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func (sp spaceEraser) Read(p []byte) (int, error) {
	n, err := sp.r.Read(p)
	count := 0
	for i := 0; i < n; i++ {
		if p[i] != 32 {
			p[count] = p[i]
			count++
		}
	}
	return count, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}
