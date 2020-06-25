package main

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func writeLine(out io.Writer, fields ...string) error {
	_, err := out.Write([]byte(strings.Join(fields, "\t") + "\r\n"))
	return err
}

func getFileType(entry os.FileInfo) string {
	if entry.IsDir() {
		return "1"
	}
	return filenameToGopherType(entry.Name())
}

func listDirectory(path string, out io.Writer) error {
	if entries, err := ioutil.ReadDir(path); err != nil {
		return err
	} else {
		for _, entry := range entries {
			filetype := getFileType(entry)
			if err := writeLine(out, filetype, entry.Name(), hostname, port); err != nil {
				return err
			}
		}
		return nil
	}
}
