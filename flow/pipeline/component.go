package pipeline

import (
	"github.com/Logiraptor/oak/flow/values"
)

type Emitter interface {
	Emit(Token, values.Value)
}

type Port struct {
	Name Token
	Type values.Type
}

type Component struct {
	InputPorts  []Port
	OutputPorts []Port
	Invoke      func(values.RecordValue, Emitter)
}

func (c Component) PortTokenByName(name string) Token {
	for _, port := range c.InputPorts {
		if port.Name.Name == name {
			return port.Name
		}
	}
	for _, port := range c.OutputPorts {
		if port.Name.Name == name {
			return port.Name
		}
	}
	panic("No such token: " + name)
}
