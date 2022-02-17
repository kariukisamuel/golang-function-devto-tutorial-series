package main

import "io"

type BadReader struct {
	err error
}

func main() {
	var r io.Reader
}

func (br *BadReader) Read(p []byte) (int, error) {

}
