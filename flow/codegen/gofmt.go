package codegen

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"strings"
)

func formatString(in string) string {
	buf := new(bytes.Buffer)
	err := formatFile(buf, strings.NewReader(in))
	if err != nil {
		panic("unable to format: " + err.Error())
	}
	return buf.String()
}

func formatFile(out io.Writer, file io.Reader) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "tmp.go", file, parser.ParseComments)
	if err != nil {
		return err
	}

	return format.Node(out, fset, f)
}
