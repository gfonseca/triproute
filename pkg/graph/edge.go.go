package graph

import "fmt"

// Edge represents an graph Edge
type Edge struct {
	Weight int
	Point  *Vertex
}

// NewEdge Build and configure new Edge
func NewEdge(weight int, point *Vertex) *Edge {
	return &Edge{
		Weight: weight,
		Point:  point,
	}
}

func (e Edge) String() string {
	return fmt.Sprintf("<Edge Weight: %d, Point: %v>", e.Weight, e.Point)
}
