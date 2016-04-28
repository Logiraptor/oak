package loader

import (
	"github.com/Logiraptor/oak/flow/parser"

	"go/types"
)

// App represents a flow program which has been loaded and
// type checked.
type App struct {
	Imports    map[string]string
	Entry      parser.ID
	Components []Component
	Flow       graph
}

// Component is a single process node within a flow program
type Component struct {
	Label   parser.ID
	Func    string
	Inputs  *types.Tuple
	Outputs *types.Tuple
}

// Component returns the component within the App with the given ID
func (a App) Component(label parser.ID) Component {
	for _, comp := range a.Components {
		if comp.Label == label {
			return comp
		}
	}
	panic("No such component: " + string(label))
}
