package main

import (
	"mime"
	"path/filepath"

	"github.com/armon/go-radix"
)

// This should be a pretty decent mapping of mimetypes to Gopher file types. I
// just gave /etc/mime.types a quick review on my machine and used my best
// judgement to map them to gopher's types. If nothing matches this list, it's
// treated as type '9'. This should probably be supplemented by something that
// checks the file extension.
var filetypes = map[string]string{
	"text/":                               "0",
	"application/gzip":                    "5",
	"application/java-archive":            "5",
	"application/rar":                     "5",
	"application/zip":                     "5",
	"application/x-cab":                   "5",
	"application/x-cbr":                   "5",
	"application/x-cbt":                   "5",
	"application/x-cbz":                   "5",
	"application/x-cpio":                  "5",
	"application/x-gtar":                  "5",
	"application/x-tar":                   "5",
	"application/x-xz":                    "5",
	"image/gif":                           "g",
	"text/html":                           "h",
	"application/xhtml+xml":               "h",
	"image/":                              "I",
	"application/msaccess":                "d",
	"application/msword":                  "d",
	"application/pdf":                     "d",
	"application/postscript":              "d",
	"application/rtf":                     "d",
	"application/vnd.ms-excel":            "d",
	"application/vnd.ms-powerpoint":       "d",
	"application/vnd.ms-word":             "d",
	"application/vnd.oasis.opendocument.": "d",
	"application/vnd.openxmlformats-officedocument.": "d",
	"application/x-dvi": "d",
	"text/rtf":          "d",
	"audio/":            "s",
	"video/":            ";",
	"text/calendar":     "c",
	"text/x-vcalendar":  "c",
	"application/mbox":  "M",
	"message/":          "M",
	"multipart/":        "M",
}

var ftPrefixes *radix.Tree

func init() {
	// I need to do this rather than just passing `filetypes` into
	// radix.NewFromMap, because the type checker fails with 'cannot use
	// map[string]string literal (type map[string]string) as type
	// map[string]interface {} in argument to radix.NewFromMap', which is
	// unfortunate.
	for prefix, ft := range filetypes {
		ftPrefixes.Insert(prefix, ft)
	}
}

func filenameToGopherType(filename string) string {
	mimetype := mime.TypeByExtension(filepath.Ext(filename))
	if _, ft, ok := ftPrefixes.LongestPrefix(mimetype); ok {
		return ft.(string)
	} else {
		return "9"
	}
}
