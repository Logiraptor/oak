package main

import (
	"fmt"
	"log"
	"os"

	"flag"

	"context"

	"net/http"

	"github.com/Logiraptor/oak/flow/language/ast"
	"github.com/Logiraptor/oak/flow/language/interpreter"
	"github.com/Logiraptor/oak/flow/language/lexer"
	"github.com/Logiraptor/oak/flow/language/parser"
	ppipeline "github.com/Logiraptor/oak/flow/pipeline"
)

type Config struct {
	Debug bool
}

func main() {
	var c = Config{}
	flag.BoolVar(&c.Debug, "debug", false, "Start a debug server on port 9999")
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Printf("usage: %s %s\n", args[0], "INPUTFILE")
		flag.Usage()
		os.Exit(1)
	}
	inputFile := args[1]
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

	frontend, pipeline, err := interpreter.Load(prog.(ast.Flow))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := context.Background()

	if c.Debug {
		log.Println("Debug server listening on port 9999")

		dv := &debugView{p: pipeline, listeners: make(chan chan message), messages: make(chan message)}
		go dv.Start()

		ctx = context.WithValue(ctx, ppipeline.InspectorKey, ppipeline.Inspector(dv))

		s := http.Server{
			Handler: dv,
			Addr:    "localhost:9999",
		}
		go s.ListenAndServe()
	}

	err = frontend.Start(ctx, pipeline)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
