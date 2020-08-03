package graph

import "fmt"

// Vertex reperesntes a graqph vertex
type Vertex struct {
	Name  string
	Edges []*Edge
}

// NewVertex builds and configure a new Vertex
func NewVertex(name string, edges []*Edge) *Vertex {
	return &Vertex{
		Name:  name,
		Edges: edges,
	}
}

// String return a debug string for the vertex object
func (v Vertex) String() string {
	return fmt.Sprintf("<Vertex: \"%s\", Edges: %s>", v.Name, v.Edges)
}

// AddEdge append a new edge to vertex
func (v *Vertex) AddEdge(edge *Edge) error {
	if edge == nil {
		return fmt.Errorf("Canot add a \"nil\" Edge to %s", v.Name)
	}

	v.Edges = append(v.Edges, edge)
	return nil
}

// <Vertex: "A", Edges: []>>
