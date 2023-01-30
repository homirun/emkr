package cmd

import (
	"flag"
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

func Flag() (*string, *string) {
	k := flag.String("k", "", "key")
	v := flag.String("v", "", "value")
	flag.Parse()
	return k, v
}
