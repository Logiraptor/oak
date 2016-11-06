package main

import (
	"fmt"
	"os"

	"github.com/Logiraptor/oak/core"
	"github.com/Logiraptor/oak/core/web"
	frontend "github.com/Logiraptor/oak/flow/frontends/web"
	"github.com/Logiraptor/oak/flow/pipeline"
	"github.com/Logiraptor/oak/flow/values"
)

func main() {
	output := core.Constant(values.StringValue("Hello, World Wide Web!"))
	success := core.Constant(values.IntValue(200))
	responder := web.HTTPResponder()
	p := pipeline.Pipeline{
		Components: []pipeline.Component{
			responder,
			output,
			success,
		},
		Pipes: []pipeline.Pipe{
			{Source: output.PortTokenByName("Output"), Dest: responder.PortTokenByName("Body")},
			{Source: success.PortTokenByName("Output"), Dest: responder.PortTokenByName("Status")},
		},
	}

	err := p.Verify()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	frontend.Serve(":8080", p)
}
