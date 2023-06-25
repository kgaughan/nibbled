package main

import (
	"mime"
	"path/filepath"

	"github.com/armon/go-radix"
)

const (
	FILE     = '0'
	MENU     = '1'
	ERROR    = '3'
	ARCHIVE  = '5'
	SEARCH   = '7'
	TELNET   = '8'
	BINARY   = '9'
	MIRROR   = '+'
	GIF      = 'g'
	IMAGE    = 'I'
	INFO     = 'i'
	HTML     = 'h'
	DOCUMENT = 'd'
	AUDIO    = 's'
	VIDEO    = ';'
	CALENDAR = 'c'
	MIME     = 'M'
)

// This should be a pretty decent mapping of mimetypes to Gopher file types. I
// just gave /etc/mime.types a quick review on my machine and used my best
// judgement to map them to gopher's types. If nothing matches this list, it's
// treated as type '9'. This should probably be supplemented by something that
// checks the file extension.
var filetypes = map[string]byte{
	"text/":                               FILE,
	"application/gzip":                    ARCHIVE,
	"application/java-archive":            ARCHIVE,
	"application/rar":                     ARCHIVE,
	"application/zip":                     ARCHIVE,
	"application/x-cab":                   ARCHIVE,
	"application/x-cbr":                   ARCHIVE,
	"application/x-cbt":                   ARCHIVE,
	"application/x-cbz":                   ARCHIVE,
	"application/x-cpio":                  ARCHIVE,
	"application/x-gtar":                  ARCHIVE,
	"application/x-tar":                   ARCHIVE,
	"application/x-xz":                    ARCHIVE,
	"image/gif":                           GIF,
	"text/html":                           HTML,
	"application/xhtml+xml":               HTML,
	"image/":                              IMAGE,
	"application/msaccess":                DOCUMENT,
	"application/msword":                  DOCUMENT,
	"application/pdf":                     DOCUMENT,
	"application/postscript":              DOCUMENT,
	"application/rtf":                     DOCUMENT,
	"application/vnd.ms-excel":            DOCUMENT,
	"application/vnd.ms-powerpoint":       DOCUMENT,
	"application/vnd.ms-word":             DOCUMENT,
	"application/vnd.oasis.opendocument.": DOCUMENT,
	"application/vnd.openxmlformats-officedocument.": DOCUMENT,
	"application/x-dvi": DOCUMENT,
	"text/rtf":          DOCUMENT,
	"audio/":            AUDIO,
	"video/":            VIDEO,
	"text/calendar":     CALENDAR,
	"text/x-vcalendar":  CALENDAR,
	"application/mbox":  MIME,
	"message/":          MIME,
	"multipart/":        MIME,
}

var ftPrefixes *radix.Tree

func init() {
	// I need to do this rather than just passing `filetypes` into
	// radix.NewFromMap, because the type checker fails with 'cannot use
	// map[string]string literal (type map[string]string) as type
	// map[string]interface {} in argument to radix.NewFromMap', which is
	// unfortunate.
	ftPrefixes = radix.New()
	for prefix, ft := range filetypes {
		ftPrefixes.Insert(prefix, ft)
	}
}

func filenameToGopherType(filename string) byte {
	mimetype := mime.TypeByExtension(filepath.Ext(filename))
	if _, ft, ok := ftPrefixes.LongestPrefix(mimetype); ok {
		return ft.(byte)
	}
	return BINARY
}
