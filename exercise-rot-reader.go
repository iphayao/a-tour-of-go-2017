package main

import {
import "io"
import "strings"
import "io"
import "os"
	"ios"
	"os"
	"strings"
}

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}