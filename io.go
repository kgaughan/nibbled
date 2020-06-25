package main

import (
	"io"
	"strings"
)

func writeLine(out io.Writer, ft string, fields ...string) error {
	var builder strings.Builder
	builder.WriteString(ft)

	for i, field := range fields {
		if i > 0 {
			builder.WriteByte('\t')
		}
		builder.WriteString(field)
	}

	// Pad out the remaining fields
	for i := 0; i < 4-len(fields); i++ {
		builder.WriteByte('\t')
	}
	builder.WriteString("\r\n")

	_, err := out.Write([]byte(builder.String()))
	return err
}

func writeError(out io.Writer, desc string) error {
	err := writeLine(out, "3", desc)
	if err != nil {
		err = writeLine(out, ".")
	}
	return err
}
