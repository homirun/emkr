package cmd

import (
	"flag"
	"fmt"
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

func Flag() {
	k := flag.String("k", "", "key")
	v := flag.String("v", "", "value")
	flag.Parse()
	fmt.Printf("%s: %s\n", *k, *v)
}
