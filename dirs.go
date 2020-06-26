package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Catalogue []Entry

func (c Catalogue) String() string {
	var builder strings.Builder
	for _, entry := range c {
		builder.WriteString(entry.String())
	}
	builder.WriteString(".\r\n")
	return builder.String()
}

func getFileType(entry os.FileInfo) byte {
	if entry.IsDir() {
		return MENU
	}
	return filenameToGopherType(entry.Name())
}

func listDirectory(localPath string, selector string) (Catalogue, error) {
	entries, err := ioutil.ReadDir(localPath)
	if err != nil {
		return nil, err
	}
	return parseDirectory(entries, selector), nil
}

func parseDirectory(entries []os.FileInfo, selector string) Catalogue {
	var result Catalogue
	for _, entry := range entries {
		name := entry.Name()
		if name == "gophermap" || strings.HasPrefix(name, ".") {
			continue
		}
		result = append(result, Entry{
			Type:     getFileType(entry),
			Name:     name,
			Selector: filepath.Join(selector, name),
			Host:     hostname,
			Port:     port,
		})
	}
	return result
}
