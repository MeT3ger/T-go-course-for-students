package main

import (
	"flag"
	"fmt"
)

type InputOptions struct {
	From      string
	Offset    int64
	Limit     int64
	BlockSize int64
	Conv      string
}

type OutputOptions struct {
	To string
}

type Options struct {
	InputOptions
	OutputOptions
}

func ParseFlags() *Options {
	var opts Options

	flag.StringVar(&opts.From, "from", "", "input filename (stdin by default)")
	flag.StringVar(&opts.To, "to", "", "output filename (stdout by default)")
	flag.Int64Var(&opts.Offset, "offset", 0, "count of copied bytes")
	flag.Int64Var(&opts.Limit, "limit", 0, "max count of copied bytes")
	flag.Int64Var(&opts.BlockSize, "block-size", 1, "copy block size")
	flag.StringVar(&opts.Conv, "conv", "", "copied bytes conv type")

	flag.Parse()

	return &opts
}

func (o *Options) validate() error {
	if o.Offset < 0 {
		return fmt.Errorf("invalid offset %d", o.Offset)
	}
	if o.Limit < 0 {
		return fmt.Errorf("invalid limit %d", o.Limit)
	}
	if o.BlockSize < 1 {
		return fmt.Errorf("invalid block size %d", o.BlockSize)
	}
	return nil
}
