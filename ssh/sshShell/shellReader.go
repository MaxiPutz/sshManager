package sshshell

import (
	"bytes"
	"io"
)

type shellReader struct {
	writer *socketWrite
	reader io.Reader
}

func (sr shellReader) Read(b []byte) (n int, err error) {

	io.Copy(sr.writer, bytes.NewReader([]byte("pwd\n")))

	io.Copy(sr.writer, sr.reader)
	return len(b), nil
}
