package flow

type Graph struct {
	Nodes []*Node
}

type Node struct {
	Label    ComponentID
	Children map[string]*Node
}

func NewGraph(app App, conf map[ComponentID]ComponentID) Graph {
	output := Graph{}

	for a, b := range conf {
		compA := app.Component(a)
		for i := 0; i < compA.Outputs.Len(); i++ {
			v := compA.Outputs.At(i)
			output.AddEdge(a, b, v.Name())
		}
	}

	return output
}

func (f *Graph) AddNode(label ComponentID) *Node {
	for _, n := range f.Nodes {
		if n.Label == label {
			return n
		}
	}
	n := &Node{
		Label:    label,
		Children: make(map[string]*Node),
	}
	f.Nodes = append(f.Nodes, n)
	return n
}

func (f *Graph) AddEdge(from, to ComponentID, label string) {
	fromNode := f.AddNode(from)
	toNode := f.AddNode(to)
	fromNode.Children[label] = toNode
}

func (f Graph) TopologicalSort(start ComponentID) []*Node {
	var current = f.AddNode(start)
	return topologicalSort(current, map[ComponentID]bool{start: true})
}

func topologicalSort(root *Node, visited map[ComponentID]bool) []*Node {
	var output = []*Node{root}
	for _, child := range root.Children {
		if visited[child.Label] {
			continue
		}
		visited[child.Label] = true

		output = append(output, topologicalSort(child, visited)...)
	}
	return output
}
