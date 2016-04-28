package codegen

import (
	"bytes"
	"fmt"
	"go/types"

	"github.com/Logiraptor/oak/flow/internal/templates"
	"github.com/Logiraptor/oak/flow/loader"
	"github.com/Logiraptor/oak/flow/parser"
)

type templateStep struct {
	LHS  []string
	Func parser.ID
	Args []string
}

type templateParams struct {
	App   loader.App
	Steps []templateStep
}

// WriteFlowApp generates a Go program that implements the given flow
// program and prints it to stdout.
func WriteFlowApp(app loader.App) {

	order := app.Flow.TopologicalSort(app.Entry)

	var steps = make([]templateStep, len(order))

	steps[0] = templateStep{
		LHS:  names(app.Component(order[1].Label).Inputs),
		Func: order[0].Label,
		Args: nil,
	}

	for i, node := range order[1 : len(order)-1] {
		steps[i+1].Func = node.Label
		steps[i+1].Args = steps[i].LHS
		steps[i+1].LHS = names(app.Component(order[i+2].Label).Inputs)
	}

	steps[len(steps)-1] = templateStep{
		LHS:  nil,
		Func: order[len(order)-1].Label,
		Args: steps[len(steps)-2].LHS,
	}

	buf := new(bytes.Buffer)
	templates.FlowApp(buf, templateParams{
		App:   app,
		Steps: steps,
	})
	fmt.Println(formatFile(buf.String()))
}

func names(in *types.Tuple) []string {
	out := make([]string, in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i] = in.At(i).Name()
	}
	return out
}
