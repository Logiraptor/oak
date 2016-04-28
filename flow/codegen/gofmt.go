package codegen

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
)

func formatFile(file string) string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "tmp.go", file, parser.ParseComments)
	if err != nil {
		panic("gofmt: unable to parse file: " + err.Error())
	}

	buf := new(bytes.Buffer)
	err = format.Node(buf, fset, f)
	if err != nil {
		panic("gofmt: unable to format file: " + err.Error())
	}

	return buf.String()
}
