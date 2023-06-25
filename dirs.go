package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Catalogue []*Entry

func (c Catalogue) String() string {
	var builder strings.Builder
	for _, entry := range c {
		builder.WriteString(entry.String())
	}
	builder.WriteString(".\r\n")
	return builder.String()
}

func getFileType(entry os.DirEntry) byte {
	if entry.IsDir() {
		return MENU
	}
	return filenameToGopherType(entry.Name())
}

func listDirectory(localPath, selector string) (Catalogue, error) {
	entries, err := os.ReadDir(localPath)
	if err != nil {
		return nil, fmt.Errorf("cannot list directory: %w", err)
	}
	return parseDirectory(entries, selector), nil
}

func parseDirectory(entries []os.DirEntry, selector string) Catalogue {
	result := make(Catalogue, 0, 10)
	for _, entry := range entries {
		name := entry.Name()
		if name == "gophermap" || strings.HasPrefix(name, ".") {
			continue
		}
		result = append(result, &Entry{
			Type:     getFileType(entry),
			Name:     name,
			Selector: filepath.Join(selector, name),
			Host:     hostname,
			Port:     port,
		})
	}
	return result
}
