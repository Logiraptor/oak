package pipeline

import (
	"fmt"

	"github.com/Logiraptor/oak/flow/values"
)

type Pipe struct {
	Source, Dest Token
}

type Pipeline struct {
	Components []Component
	Pipes      []Pipe
}

func (p *Pipeline) Verify() error {
	var specifiedInputs = make(map[Token]struct{})

	for _, pipe := range p.Pipes {
		if sourcePort, ok := findPortByName(p.Components, pipe.Source); ok {
			if destPort, ok := findPortByName(p.Components, pipe.Dest); ok {

				specifiedInputs[destPort.Name] = struct{}{}

				if !values.EqualTypes(sourcePort.Type, destPort.Type) {
					return fmt.Errorf("Type error: port %s expects type %s, but port %s produces type %s",
						sourcePort.Name, values.TypeToString(sourcePort.Type),
						destPort.Name, values.TypeToString(destPort.Type))
				}
				continue
			}
		}
	}

	for _, component := range p.Components {
		for _, input := range component.InputPorts {
			if _, ok := specifiedInputs[input.Name]; !ok {
				return fmt.Errorf("Port %s has not had its input specified", input.Name)
			}
		}
	}

	return nil
}

func findNewName(usedNames map[string]struct{}, orig string) string {
	if _, ok := usedNames[orig]; ok {
		return findNewName(usedNames, orig+"'")
	}
	usedNames[orig] = struct{}{}
	return orig
}

func findPortByName(components []Component, name Token) (Port, bool) {
	for _, component := range components {
		for _, port := range component.InputPorts {
			if port.Name == name {
				return port, true
			}
		}
		for _, port := range component.OutputPorts {
			if port.Name == name {
				return port, true
			}
		}
	}

	return Port{}, false
}
