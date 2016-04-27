package main

import (
	"fmt"
	"os"

	"github.com/Logiraptor/oak/flow"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("USAGE: %s FILE\n", os.Args[0])
		return
	}

	conf, err := flow.ParseYAMLFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	app, errs := flow.Load(conf)
	if errs != nil {
		for _, err := range errs {
			fmt.Println(err.Error())
		}
		return
	}

	flow.WriteFlowApp(app)
}
