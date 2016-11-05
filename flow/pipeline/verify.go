package pipeline

import (
	"fmt"

	"github.com/Logiraptor/oak/flow/values"
)

type Pipe struct {
	Source, Dest values.Token
}

type Pipeline struct {
	Components []Component
	Pipes      []Pipe
}

func (p *Pipeline) Verify() error {
	var specifiedInputs = make(map[values.Token]struct{})
	var derivedGenericTypes = make(values.TypeEnv)

	for _, pipe := range p.Pipes {
		if sourcePort, ok := findPortByName(p.Components, pipe.Source); ok {
			if destPort, ok := findPortByName(p.Components, pipe.Dest); ok {
				specifiedInputs[destPort.Name] = struct{}{}

				_, unifyable := values.UnifyType(derivedGenericTypes, sourcePort.Type, destPort.Type)
				if !unifyable {
					return fmt.Errorf("Type error: port %s produces type %s, but port %s expects type %s",
						sourcePort.Name.Name, values.TypeToString(sourcePort.Type),
						destPort.Name.Name, values.TypeToString(destPort.Type))
				}
			}
		}
	}

	for _, component := range p.Components {
		for _, input := range component.InputPorts {
			if _, ok := specifiedInputs[input.Name]; !ok {
				return fmt.Errorf("Port %s has not had its input specified", input.Name.Name)
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

func findPortByName(components []Component, name values.Token) (Port, bool) {
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
