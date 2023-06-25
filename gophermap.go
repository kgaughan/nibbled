package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func loadGopherMap(out io.Writer, localPath, selector string) error {
	f, err := os.Open(filepath.Join(localPath, "gophermap"))
	if err != nil {
		return fmt.Errorf("could not open gophermap: %w", err)
	}
	defer f.Close()
	return processGopherMap(f, out, localPath, selector)
}

func processGopherMap(in io.Reader, out io.Writer, localPath, selector string) error {
	reader := bufio.NewReader(in)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("cannot read from gophermap: %w", err)
		}
		if bytes.HasPrefix(line, []byte("#")) || bytes.HasPrefix(line, []byte("!")) {
			continue
		}
		if bytes.HasPrefix(line, []byte(".")) {
			if _, err := out.Write([]byte(".\r\n")); err != nil {
				return err //nolint:wrapcheck
			}
			break
		}
		if bytes.HasPrefix(line, []byte("*")) {
			catalogue, err := listDirectory(localPath, selector)
			if err == nil {
				_, err = write(out, catalogue)
			}
			return fmt.Errorf("could not list directory: %w", err)
		}
		if !bytes.ContainsRune(line, '\t') {
			if _, err := out.Write([]byte("i")); err != nil {
				return err //nolint:wrapcheck
			}
		}
		if _, err := out.Write(line); err != nil {
			return err //nolint:wrapcheck
		}
		if _, err := out.Write([]byte("\r\n")); err != nil {
			return err //nolint:wrapcheck
		}
	}
	return nil
}
