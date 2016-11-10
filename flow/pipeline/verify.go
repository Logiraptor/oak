package pipeline

import (
	"fmt"
	"io"

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
				result, unifyable := values.UnifyType(derivedGenericTypes, sourcePort.Type, destPort.Type)
				if !unifyable {
					return fmt.Errorf("Type error: port %s produces type %s, but port %s expects type %s",
						sourcePort.Name.Name, values.TypeToString(sourcePort.Type),
						destPort.Name.Name, values.TypeToString(destPort.Type))
				}
				fmt.Println("Unified as", values.TypeToString(result))
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

type errWriter struct {
	err error
	w   io.Writer
}

func (e *errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	n, err := e.w.Write(buf)
	e.err = err
	return n, err
}

func (p *Pipeline) WriteToDot(w io.Writer) error {
	we := &errWriter{w: w}
	io.WriteString(we, "digraph {")
	for _, pipe := range p.Pipes {
		var sourceIndex int
		var sourceType values.Type
		var destIndex int
		var destType values.Type
		for i, component := range p.Components {
			for _, port := range component.InputPorts {
				if port.Name == pipe.Source {
					sourceIndex = i
					sourceType = port.Type
				}
				if port.Name == pipe.Dest {
					destIndex = i
					destType = port.Type
				}
			}
			for _, port := range component.OutputPorts {
				if port.Name == pipe.Source {
					sourceIndex = i
					sourceType = port.Type
				}
				if port.Name == pipe.Dest {
					destIndex = i
					destType = port.Type
				}
			}
		}

		fmt.Fprintf(we, "%q -> %q [headlabel = %q taillabel = %q];",
			p.Components[sourceIndex].Name.Name,
			p.Components[destIndex].Name.Name,

			values.TypeToString(destType),
			values.TypeToString(sourceType))
	}
	io.WriteString(we, "}")
	return we.err
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
