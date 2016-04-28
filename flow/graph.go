package flow

type graph struct {
	Nodes []*node
}

type node struct {
	Label    ComponentID
	Children map[string]*node
}

func newGraph(app App, conf map[ComponentID]ComponentID) graph {
	output := graph{}

	for a, b := range conf {
		compA := app.Component(a)
		for i := 0; i < compA.Outputs.Len(); i++ {
			v := compA.Outputs.At(i)
			output.AddEdge(a, b, v.Name())
		}
	}

	return output
}

func (f *graph) AddNode(label ComponentID) *node {
	for _, n := range f.Nodes {
		if n.Label == label {
			return n
		}
	}
	n := &node{
		Label:    label,
		Children: make(map[string]*node),
	}
	f.Nodes = append(f.Nodes, n)
	return n
}

func (f *graph) AddEdge(from, to ComponentID, label string) {
	fromNode := f.AddNode(from)
	toNode := f.AddNode(to)
	fromNode.Children[label] = toNode
}

func (f graph) TopologicalSort(start ComponentID) []*node {
	var current = f.AddNode(start)
	return topologicalSort(current, map[ComponentID]bool{start: true})
}

func topologicalSort(root *node, visited map[ComponentID]bool) []*node {
	var output = []*node{root}
	for _, child := range root.Children {
		if visited[child.Label] {
			continue
		}
		visited[child.Label] = true

		output = append(output, topologicalSort(child, visited)...)
	}
	return output
}
