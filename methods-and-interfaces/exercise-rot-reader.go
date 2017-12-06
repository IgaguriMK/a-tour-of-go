package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)

	for i := 0; i < n; i++ {
		ch := b[i]
		switch {
		case ch < 'A':
		case ch <= 'Z':
			b[i] = (ch-'A'+13)%26 + 'A'
		case ch < 'a':
		case ch <= 'z':
			b[i] = (ch-'a'+13)%26 + 'a'
		}
	}

	if err == io.EOF {
		return n, io.EOF
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
