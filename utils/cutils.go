package utils

import (
	"io"
	"golang.org/x/text/encoding/charmap"
	"bytes"
	"golang.org/x/text/transform"
	"bufio"
)

func CopyWithEncoding(f io.Reader, charmap *charmap.Charmap) bytes.Buffer {
	var buffer bytes.Buffer
	r := transform.NewReader(f, charmap.NewDecoder())
	sc := bufio.NewScanner(r)

	for sc.Scan() {
		buffer.WriteString(sc.Text())
		buffer.WriteString("\n")
	}

	return buffer
}