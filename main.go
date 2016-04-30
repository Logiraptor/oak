package main

import (
	"fmt"
	"os"

	"github.com/Logiraptor/oak/flow/codegen"
	"github.com/Logiraptor/oak/flow/loader"
	"github.com/Logiraptor/oak/flow/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("USAGE: %s FILE\n", os.Args[0])
		return
	}

	conf, err := parser.ParseFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	app, errs := loader.Load(conf)
	if errs != nil {
		for _, err := range errs {
			fmt.Println(err.Error())
		}
		return
	}

	codegen.WriteFlowApp(app)
}
