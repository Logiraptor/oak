package main

import (
	"fmt"
	"os"

	"context"

	"github.com/Logiraptor/oak/flow/language/ast"
	"github.com/Logiraptor/oak/flow/language/interpreter"
	"github.com/Logiraptor/oak/flow/language/lexer"
	"github.com/Logiraptor/oak/flow/language/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s %s\n", os.Args[0], "INPUTFILE")
		os.Exit(1)
	}
	inputFile := os.Args[1]
	l, err := lexer.NewLexerFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p := parser.NewParser()
	prog, err := p.Parse(l)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pipeline := interpreter.Interp(prog.(ast.Pipeline))

	err = pipeline.Verify()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pipeline.Run(context.Background())
}
