package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

func loadGopherMap(out io.Writer, localPath string, selector string) error {
	f, err := os.Open(filepath.Join(localPath, "gophermap"))
	if err != nil {
		return err
	}
	defer f.Close()
	return processGopherMap(f, out, localPath, selector)
}

func processGopherMap(in io.Reader, out io.Writer, localPath string, selector string) error {
	reader := bufio.NewReader(in)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if bytes.HasPrefix(line, []byte("#")) || bytes.HasPrefix(line, []byte("!")) {
			continue
		}
		if bytes.HasPrefix(line, []byte(".")) {
			out.Write([]byte(".\r\n"))
			break
		}
		if bytes.HasPrefix(line, []byte("*")) {
			if catalogue, err := listDirectory(localPath, selector); err != nil {
				return err
			} else {
				_, err := write(out, catalogue)
				return err
			}
		}
		if !bytes.ContainsRune(line, '\t') {
			out.Write([]byte("i"))
		}
		out.Write(line)
		out.Write([]byte("\r\n"))
	}
	return nil
}
