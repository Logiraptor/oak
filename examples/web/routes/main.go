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
	eq := core.Eq()
	cond := core.Cond()
	path1 := core.Constant(values.StringValue("/bye"))
	path2 := core.Constant(values.StringValue("/hello"))

	message1 := core.Constant(values.StringValue("Goodbye!"))
	message2 := core.Constant(values.StringValue("Hello!"))

	success := core.Constant(values.IntValue(200))
	responder := web.HTTPResponder()
	request := web.HTTPRequest()
	path := core.FieldAccessor("Path")
	p := pipeline.Pipeline{
		Components: []pipeline.Component{
			request, responder,
			cond, path1, path2, path,
			message1, message2,
			success, eq,
		},
		Pipes: []pipeline.Pipe{

			// Path Extraction
			pipeline.Pipe{Source: request.PortTokenByName("Output"), Dest: path.PortTokenByName("Input")},

			// Condition
			pipeline.Pipe{Source: path.PortTokenByName("Output"), Dest: eq.PortTokenByName("LHS")},
			pipeline.Pipe{Source: path1.PortTokenByName("Output"), Dest: eq.PortTokenByName("RHS")},
			pipeline.Pipe{Source: eq.PortTokenByName("Output"), Dest: cond.PortTokenByName("Condition")},

			// Message options
			pipeline.Pipe{Source: message1.PortTokenByName("Output"), Dest: cond.PortTokenByName("Consequence")},
			pipeline.Pipe{Source: message2.PortTokenByName("Output"), Dest: cond.PortTokenByName("Alternative")},

			// Final output
			pipeline.Pipe{Source: cond.PortTokenByName("Output"), Dest: responder.PortTokenByName("Body")},
			pipeline.Pipe{Source: success.PortTokenByName("Output"), Dest: responder.PortTokenByName("Status")},
		},
	}

	f, err := os.Create("debug")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = p.WriteToDot(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = p.Verify()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	frontend.Serve(":8080", p)
}
