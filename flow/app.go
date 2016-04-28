package flow

import "go/types"

// ComponentID is a unique identifier for a flow process node
type ComponentID string

// App represents a flow program which has been loaded and
// type checked.
type App struct {
	Imports    map[string]string
	Entry      ComponentID
	Components []Component
	Flow       graph
}

// Component is a single process node within a flow program
type Component struct {
	Label   ComponentID
	Func    string
	Inputs  *types.Tuple
	Outputs *types.Tuple
}

// Component returns the component within the App with the given ID
func (a App) Component(label ComponentID) Component {
	for _, comp := range a.Components {
		if comp.Label == label {
			return comp
		}
	}
	panic("No such component: " + string(label))
}
