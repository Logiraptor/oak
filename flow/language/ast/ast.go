package ast

import "github.com/Logiraptor/oak/flow/language/token"
import "github.com/Logiraptor/oak/flow/values"
import "strconv"

type Attrib interface{}

type Component struct {
	Name, Constructor string
	Args              []values.Value
}

type Port struct {
	Component string
	Port      string
}

type Connection struct {
	Source, Dest Port
}

type Pipeline struct {
	Components []Component
	Pipes      []Connection
}

func NewComponent(name, ctor, args Attrib) Component {
	return Component{
		Name:        string(name.(*token.Token).Lit),
		Constructor: string(ctor.(*token.Token).Lit),
		Args:        args.([]values.Value),
	}
}

func NewConnection(a, b, c, d Attrib) Connection {
	return Connection{
		Source: Port{
			Component: string(a.(*token.Token).Lit),
			Port:      string(b.(*token.Token).Lit),
		},
		Dest: Port{
			Component: string(c.(*token.Token).Lit),
			Port:      string(d.(*token.Token).Lit),
		},
	}
}

type Constant struct {
	val  values.Value
	dest Port
}

func NewConstant(a, b, c Attrib) Constant {
	return Constant{
		val: a.(values.Value),
		dest: Port{
			Component: string(b.(*token.Token).Lit),
			Port:      string(c.(*token.Token).Lit),
		},
	}
}

func AddComponent(pipeline, component Attrib) Pipeline {
	p := pipeline.(Pipeline)
	return Pipeline{
		Components: append(p.Components, component.(Component)),
		Pipes:      p.Pipes,
	}
}

func AddConnection(pipeline, pipe Attrib) Pipeline {
	p := pipeline.(Pipeline)
	return Pipeline{
		Components: p.Components,
		Pipes:      append(p.Pipes, pipe.(Connection)),
	}
}

var implicitConstants = 0

func AddConstant(pipeline, constant Attrib) Pipeline {
	implicitConstants++
	c := constant.(Constant)
	p := pipeline.(Pipeline)
	name := "$$implicitConst" + strconv.Itoa(implicitConstants)
	return Pipeline{
		Components: append(p.Components, Component{
			Args:        []values.Value{c.val},
			Name:        name,
			Constructor: "Constant",
		}),
		Pipes: append(p.Pipes, Connection{
			Dest:   c.dest,
			Source: Port{Component: name, Port: "Output"},
		}),
	}
}

func NewString(val Attrib) values.StringValue {
	buf := val.(*token.Token).Lit
	return values.StringValue(string(buf[1 : len(buf)-1]))
}

func NewInt(val Attrib) (values.IntValue, error) {
	i, err := strconv.Atoi(string(val.(*token.Token).Lit))
	return values.IntValue(i), err
}

type Flow struct {
	Frontend
	Pipeline
}

func NewFlow(frontend, pipeline Attrib) (Flow, error) {
	front, ok := frontend.(Frontend)
	if !ok {
		front = Frontend{
			Constructor: "cli",
		}
	}
	return Flow{
		Frontend: front,
		Pipeline: pipeline.(Pipeline),
	}, nil
}

type Frontend struct {
	Constructor string
}

func NewFrontend(ctor Attrib) (Frontend, error) {
	name := string(ctor.(*token.Token).Lit)
	return Frontend{
		Constructor: name,
	}, nil
}
