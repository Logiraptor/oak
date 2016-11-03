package pipeline

import (
	"github.com/Logiraptor/oak/flow/values"
)

type Port struct {
	Name string
	Type values.Type
}

type Component struct {
	InputPorts  []Port
	OutputPorts []Port
	Invoke      func([]values.Value) []values.Value
}
