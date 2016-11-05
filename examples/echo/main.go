package main

import (
	"fmt"

	"os"

	"github.com/Logiraptor/oak/core"
	"github.com/Logiraptor/oak/flow/pipeline"
)

func main() {
	input := core.StdinLines()
	logger := core.Logger()
	var p = pipeline.Pipeline{
		Components: []pipeline.Component{
			input,
			logger,
		},
		Pipes: []pipeline.Pipe{
			{
				Source: input.PortTokenByName("Output"),
				Dest:   logger.PortTokenByName("Input"),
			},
		},
	}

	err := p.Verify()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p.Run()
}
