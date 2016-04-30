package loader

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestAddNode(t *testing.T) {
	g := new(graph)
	a1 := g.AddNode("a")
	a2 := g.AddNode("a")
	assert.True(t, a1 == a2, "Multiple calls return the same node")
}

func TestTopologicalSort(t *testing.T) {
	g := new(graph)
	g.AddEdge("a", "b", "a->b")
	g.AddEdge("b", "d", "b->d")
	g.AddEdge("a", "c", "a->c")
	g.AddEdge("c", "d", "c->d")
	sorted := g.TopologicalSort("a")
	assert.Len(t, sorted, 4)
	assert.Equal(t, g.AddNode("a"), sorted[0])
	b := g.AddNode("b")
	c := g.AddNode("c")
	d := g.AddNode("d")
	assert.True(t,
		(sorted[1] == b && ((sorted[2] == c && sorted[3] == d) || (sorted[2] == d && sorted[3] == c))) ||
			(sorted[1] == c && ((sorted[2] == b && sorted[3] == d) || (sorted[2] == d && sorted[3] == b))),
		"dependency order is preserved")
}
