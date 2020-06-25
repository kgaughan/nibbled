package main

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func getFileType(entry os.FileInfo) string {
	if entry.IsDir() {
		return "1"
	}
	return filenameToGopherType(entry.Name())
}

func listDirectory(out io.Writer, localPath string, selector string) error {
	if entries, err := ioutil.ReadDir(localPath); err != nil {
		return err
	} else {
		for _, entry := range entries {
			name := entry.Name()
			if name == "gophermap" || strings.HasPrefix(name, ".") {
				continue
			}
			filetype := getFileType(entry)
			newSelector := filepath.Join(selector, name)
			if err := writeLine(out, filetype, name, newSelector, hostname, port); err != nil {
				return err
			}
		}
		return nil
	}
}
