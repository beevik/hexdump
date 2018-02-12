package hexdump

import (
	"fmt"
	"io"
)

// Format describes an output format for hex-dumped data.
type Format uint8

// All possible hex-dump formats.
const (
	FormatGo Format = iota
)

// Dump a slice to a writer using the requested format.
func Dump(b []byte, format Format, w io.Writer) error {
	switch format {
	case FormatGo:
		return dumpGo(b, w)
	default:
		return nil
	}
}

var b = []byte{
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
}

func dumpGo(b []byte, w io.Writer) error {
	ww := newWriter(w)

	ww.Printf("var b = []byte{\n")

	for i := 0; i < len(b); i += 8 {
		r := i + 8
		if r > len(b) {
			r = len(b)
		}

		ww.Printf("\t")

		var j int
		for j = i; j < r-1; j++ {
			ww.Printf("0x%02x, ", b[j])
		}
		ww.Printf("0x%02x,\n", b[j])
	}

	ww.Printf("}\n")

	return ww.err
}

type writer struct {
	w   io.Writer
	err error
}

func newWriter(w io.Writer) *writer {
	return &writer{w: w}
}

func (w *writer) Printf(format string, a ...interface{}) {
	if w.err != nil {
		return
	}
	_, w.err = fmt.Fprintf(w.w, format, a...)
}
