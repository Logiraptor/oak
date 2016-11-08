package interpreter

import (
	"fmt"

	"github.com/Logiraptor/oak/core"
	"github.com/Logiraptor/oak/flow/language/ast"
	"github.com/Logiraptor/oak/flow/pipeline"
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
	default:
		panic(fmt.Sprintf("undefined constructor %s", component.Constructor))
	}
}
