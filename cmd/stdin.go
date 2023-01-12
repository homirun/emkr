package cmd

import (
	"io"
	"os"
)

func Pipe() ([]byte, error) {
	r := os.Stdin
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//func File() ([]byte, error) {
//	return
//}
