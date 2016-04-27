package flow

import "go/types"

type ComponentID string

type App struct {
	Imports    map[string]string
	Entry      ComponentID
	Components []Component
	Flow       Graph
}

type Component struct {
	Label   ComponentID
	Func    string
	Inputs  *types.Tuple
	Outputs *types.Tuple
}

func (a App) Component(label ComponentID) Component {
	for _, comp := range a.Components {
		if comp.Label == label {
			return comp
		}
	}
	panic("No such component: " + string(label))
}
