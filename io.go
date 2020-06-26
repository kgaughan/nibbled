package main

import (
	"fmt"
	"io"
)

func write(out io.Writer, obj fmt.Stringer) (int, error) {
	return out.Write([]byte(obj.String()))
}

func writeError(out io.Writer, desc string) {
	write(out, Entry{Type: ERROR, Name: desc})
	out.Write([]byte(".\r\n"))
}
