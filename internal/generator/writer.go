package generator

import (
	gofmt "go/format"
	"io"
)

var lineBreak = []byte("\n")

type writer struct {
	internal io.Writer
}

func (w *writer) Write(b []byte) (int, error) {
	formatted, err := gofmt.Source(b)
	if err != nil {
		return 0, err
	}

	n1, err := w.internal.Write(lineBreak)
	if err != nil {
		return 0, err
	}

	n2, err := w.internal.Write(formatted)
	if err != nil {
		return 0, err
	}

	n3, err := w.internal.Write(lineBreak)
	if err != nil {
		return 0, err
	}

	return n1 + n2 + n3, nil
}
