package pipeline

import (
	"fmt"

	"github.com/Logiraptor/oak/flow/values"
	"github.com/kr/pretty"
)

type Pipe struct {
	SourcePort      string
	DestPort        string
	SourceComponent int
	DestComponent   int
}

type Pipeline struct {
	Components []Component
	Pipes      []Pipe
}

func (p *Pipeline) Verify() error {
	pretty.Print(p)
	p.alphaConvert()
	pretty.Print(p)

	var specifiedInputs = make(map[string]struct{})

	for _, pipe := range p.Pipes {
		sourceComponent := p.Components[pipe.SourceComponent]
		destComponent := p.Components[pipe.DestComponent]

		if sourcePort, ok := findPortByName(sourceComponent.OutputPorts, pipe.SourcePort); ok {
			if destPort, ok := findPortByName(destComponent.InputPorts, pipe.DestPort); ok {

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

func (p *Pipeline) alphaConvert() {
	var usedNames = make(map[string]struct{})
	for i, component := range p.Components {
		for inputIndex, input := range component.InputPorts {
			newName := findNewName(usedNames, input.Name)
			renamePipes(i, input.Name, newName, p.Pipes)
			component.InputPorts[inputIndex].Name = newName
		}

		for outputIndex, output := range component.OutputPorts {
			newName := findNewName(usedNames, output.Name)
			renamePipes(i, output.Name, newName, p.Pipes)
			component.OutputPorts[outputIndex].Name = newName
		}
	}
}

func renamePipes(i int, name, newName string, pipes []Pipe) {
	for j, pipe := range pipes {
		if pipe.DestComponent == i && pipe.DestPort == name {
			pipes[j].DestPort = newName
		}
		if pipe.SourceComponent == i && pipe.SourcePort == name {
			pipes[j].SourcePort = newName
		}
	}
}

func findNewName(usedNames map[string]struct{}, orig string) string {
	if _, ok := usedNames[orig]; ok {
		return findNewName(usedNames, orig+"'")
	}
	usedNames[orig] = struct{}{}
	return orig
}

func findPortByName(ports []Port, name string) (Port, bool) {
	for _, port := range ports {
		if port.Name == name {
			return port, true
		}
	}
	return Port{}, false
}
