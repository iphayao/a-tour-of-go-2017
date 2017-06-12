package main

import "golang.org/x/tour/reader"

type MyReader struct {

}

// TODO: Add a Reader([]byte)(int, error) method to MyReader

func main() {
	reader.Validate(MyReader{})
}