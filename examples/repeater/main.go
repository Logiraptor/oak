package main

import (
	"fmt"

	"os"

	"time"

	"github.com/Logiraptor/oak/core"
	"github.com/Logiraptor/oak/flow/pipeline"
	"github.com/Logiraptor/oak/flow/values"
)

/*

Repeater(5s)        ->
Const(Hello, World) ->
Const(Hello, World) -> Cond() -> Logger()

*/

func main() {
	cond := core.Cond()
	constant := core.Constant(values.StringValue("Hello, World!"))
	constant2 := core.Constant(values.StringValue("Hello, World!"))
	repeater := core.Repeater(time.Second * 1)
	logger := core.Logger()
	var p = pipeline.Pipeline{
		Components: []pipeline.Component{
			cond,
			constant,
			repeater,
			logger,
			constant2,
		},
		Pipes: []pipeline.Pipe{
			{
				Source: repeater.PortTokenByName("Output"),
				Dest:   cond.PortTokenByName("Condition"),
			},
			{
				Source: constant.PortTokenByName("Output"),
				Dest:   cond.PortTokenByName("Consequence"),
			},
			{
				Source: constant2.PortTokenByName("Output"),
				Dest:   cond.PortTokenByName("Alternative"),
			},
			{
				Source: cond.PortTokenByName("Output"),
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
