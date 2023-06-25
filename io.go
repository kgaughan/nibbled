package main

import (
	"fmt"
	"io"
)

func write(out io.Writer, obj fmt.Stringer) (int, error) {
	return out.Write([]byte(obj.String())) //nolint:wrapcheck
}

func writeError(out io.Writer, desc string) error {
	if _, err := write(out, Entry{Type: ERROR, Name: desc}); err != nil {
		return err
	}
	_, err := out.Write([]byte(".\r\n"))
	return err //nolint:wrapcheck
}
