package interpreter

import (
	"fmt"

	"time"

	"github.com/Logiraptor/oak/core"
	"github.com/Logiraptor/oak/flow/language/ast"
	"github.com/Logiraptor/oak/flow/pipeline"
	"github.com/Logiraptor/oak/flow/values"
)

func Interp(p ast.Pipeline) pipeline.Pipeline {
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
	return output
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
		fmt.Println(value, []byte(value))
		d, err := time.ParseDuration(value)
		if err != nil {
			panic(fmt.Sprintf("Invalid duration passed to repeater: %s", err.Error()))
		}
		return core.Repeater(d)
	default:
		panic(fmt.Sprintf("undefined constructor %s", component.Constructor))
	}
}
