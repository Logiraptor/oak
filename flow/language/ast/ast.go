package ast

import "github.com/Logiraptor/oak/flow/language/token"

type Attrib interface{}

type Component struct {
	Name, Constructor string
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

func NewComponent(name, ctor Attrib) Component {
	return Component{
		Name:        string(name.(*token.Token).Lit),
		Constructor: string(ctor.(*token.Token).Lit),
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
