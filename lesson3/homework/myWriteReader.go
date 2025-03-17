package main

import (
	"fmt"
	"io"
	"os"
)

type WriteReader interface {
	Read(opts *InputOptions) ([]byte, error)
	Write(opts *OutputOptions, data []byte) error
}

type MyWriteReader struct{}

func NewMyWriteReader() *MyWriteReader {
	return &MyWriteReader{}
}

func (o *Options) Do(rw WriteReader) error {
	strings, err := rw.Read(&o.InputOptions)
	if err != nil {
		return err
	}

	err = rw.Write(&o.OutputOptions, strings)
	if err != nil {
		return err
	}
	return nil
}

func (w *MyWriteReader) Read(opts *InputOptions) ([]byte, error) {
	if opts.From == "" {
		return w.readFromStdIn(opts)
	} else {
		return w.readFromFile(opts)
	}
}

func (w *MyWriteReader) readFromFile(opts *InputOptions) ([]byte, error) {
	fileStat, err := os.Stat(opts.From)
	if err != nil {
		return nil, err
	}
	if fileStat.Size() < opts.Offset {
		return nil, fmt.Errorf("file %s is too small", opts.From)
	}

	file, err := os.OpenFile(opts.From, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	if opts.Offset != 0 {
		_, err = file.Seek(opts.Offset, 0)
		if err != nil {
			return nil, err
		}
	}

	var b []byte
	if opts.Limit > 0 {
		var n int
		b = make([]byte, opts.Limit)
		n, err = file.Read(b)
		if err != nil {
			return nil, err
		}
		b = b[:n]
	} else {
		b, err = io.ReadAll(file)
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}

func (w *MyWriteReader) readFromStdIn(opts *InputOptions) ([]byte, error) {
	discarded, err := io.CopyN(io.Discard, os.Stdin, opts.Offset)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("skipped %d bytes", opts.Offset-discarded)
	}

	// Если EOF достигнут раньше, чем мы пропустили offset байт
	if discarded < opts.Offset {
		return nil, fmt.Errorf("skipped %d bytes", opts.Offset-discarded)
	}

	var b []byte
	if opts.Limit > 0 {
		b = make([]byte, opts.Limit)
		n, err := io.ReadFull(os.Stdin, b)
		if err != nil {
			return nil, err
		}
		b = b[:n]
	} else {
		var err error
		b, err = io.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}

func (w *MyWriteReader) Write(opts *OutputOptions, data []byte) error {
	if opts.To != "" {
		err := w.writeToFile(opts, data)
		if err != nil {
			return err
		}
	} else {
		err := w.writeToStdOut(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *MyWriteReader) writeToFile(opts *OutputOptions, data []byte) error {
	a, err := os.Stat(opts.To)
	if a != nil {
		return fmt.Errorf("file %s already exists", opts.To)
	}

	file, err := os.Create(opts.To)
	if err != nil {
		return err
	}

	n, err := file.Write(data)
	if err != nil {
		return err
	}
	if n != len(data) {
		return fmt.Errorf("short write")
	}

	err = file.Close()
	return err
}

func (w *MyWriteReader) writeToStdOut(data []byte) error {
	n, err := os.Stdout.Write(data)
	if n != len(data) {
		return fmt.Errorf("short write")
	}
	return err
}
