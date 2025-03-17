package main

import (
	"fmt"
	"os"
)

func main() {
	opts := ParseFlags()

	err := opts.validate()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	wr := NewMyWriteReader()
	err = opts.Do(wr)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
