package interpreter

import (
	"fmt"

	"time"

	"context"

	"github.com/Logiraptor/oak/core"
	"github.com/Logiraptor/oak/core/web"
	wf "github.com/Logiraptor/oak/flow/frontends/web"
	"github.com/Logiraptor/oak/flow/language/ast"
	"github.com/Logiraptor/oak/flow/pipeline"
	"github.com/Logiraptor/oak/flow/values"
)

type Frontend struct {
	Start func(context.Context, pipeline.Pipeline) error
}

func Load(p ast.Flow) (Frontend, pipeline.Pipeline, error) {
	var components = make(map[string]pipeline.Component)
	var output pipeline.Pipeline
	for _, component := range p.Components {
		c := interpComponent(component)
		components[component.Name] = c
		output.Components = append(output.Components, c)
	}

	for _, pipe := range p.Pipes {
		sourceComponent := components[pipe.Source.Component]
		destComponent := components[pipe.Dest.Component]

		output.Pipes = append(output.Pipes, pipeline.Pipe{
			Dest:   destComponent.PortTokenByName(pipe.Dest.Port),
			Source: sourceComponent.PortTokenByName(pipe.Source.Port),
		})
	}

	err := output.Verify()
	if err != nil {
		return Frontend{}, pipeline.Pipeline{}, err
	}

	frontend, err := constructFrontend(p.Frontend)
	if err != nil {
		return Frontend{}, pipeline.Pipeline{}, err
	}

	return frontend, output, nil
}

func constructFrontend(f ast.Frontend) (Frontend, error) {
	switch f.Constructor {
	case "cli":
		return Frontend{
			Start: func(ctx context.Context, p pipeline.Pipeline) error {
				p.Run(ctx)
				return nil
			},
		}, nil
	case "web":
		return Frontend{
			Start: func(ctx context.Context, p pipeline.Pipeline) error {
				wf.Serve(":8080", ctx, p)
				return nil
			},
		}, nil
	}
	return Frontend{}, fmt.Errorf("undefined frontend constructor: %s", f.Constructor)
}

func interpComponent(component ast.Component) pipeline.Component {
	switch component.Constructor {
	case "Logger":
		return core.Logger()
	case "StdinLines":
		return core.StdinLines()
	case "Cond":
		return core.Cond()
	case "Constant":
		return core.Constant(component.Args[0])
	case "Repeater":
		value := string(component.Args[0].(values.StringValue))
		d, err := time.ParseDuration(value)
		if err != nil {
			panic(fmt.Sprintf("Invalid duration passed to repeater: %s", err.Error()))
		}
		return core.Repeater(d)
	case "HTTPResponder":
		return web.HTTPResponder()
	case "HTTPRequest":
		return web.HTTPRequest()
	case "FieldAccessor":
		return core.FieldAccessor(string(component.Args[0].(values.StringValue)))
	case "Eq":
		return core.Eq()
	default:
		panic(fmt.Sprintf("undefined constructor %s", component.Constructor))
	}
}
