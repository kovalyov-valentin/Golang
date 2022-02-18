package main

import (
	"fmt"
	"io"
	"bytes"
)

func Copy(dst Writer, src Reader) (written int64, err error) {
	var buf bytes.Buffer
	buf.Write([]byte("test"))
	return
}