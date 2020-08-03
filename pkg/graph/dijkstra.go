package graph

import (
	"fmt"
)

type vertexregister struct {
	vertex        *Vertex
	previous      *Vertex
	shortDistance int64
}

// PathTable is used to store graph distances
type PathTable map[string]*vertexregister

func inMap(key string, m map[string]*Vertex) bool {
	if _, ok := m[key]; ok {
		return true
	}

	return false
}

func getMinorVertex(table PathTable, visited map[string]*Vertex) *Vertex {
	var minor *vertexregister
	for k, v := range table {

		if inMap(k, visited) {
			continue
		}

		if minor == nil {
			minor = v
			continue
		}

		if v.shortDistance < minor.shortDistance {
			minor = v
		}
	}
	return minor.vertex
}

func filterVisited(visited map[string]*Vertex, edges []*Edge) []*Edge {
	out := make([]*Edge, 0)

	for _, e := range edges {
		name := e.Point.Name
		if !inMap(name, visited) {
			out = append(out, e)
		}
	}

	return out
}

func maxInt64() int64 {
	return 1<<63 - 1
}

func reverseArr(input []*Vertex) []*Vertex {
	if len(input) == 0 {
		return input
	}

	return append(reverseArr(input[1:]), input[0])
}

func rewind(pTable PathTable, start, end string) ([]*Vertex, error) {
	out := make([]*Vertex, 0)
	endVertex := pTable[end]
	out = append(out, endVertex.vertex)
	for endVertex.vertex.Name != start {
		if endVertex.previous == nil {
			return nil, fmt.Errorf("No route found to this input")
		}
		endVertex = pTable[endVertex.previous.Name]
		out = append(out, endVertex.vertex)
	}

	return reverseArr(out), nil
}

// Dijkstra algorith calculate short path between two vertexs
func Dijkstra(graph map[string]*Vertex, start, end string) ([]*Vertex, int64, error) {

	if !inMap(start, graph) {
		return nil, 0, fmt.Errorf("Invalid origin")
	}

	if !inMap(end, graph) {
		return nil, 0, fmt.Errorf("Invalid destiny: %s", end)
	}

	MaxInt := maxInt64()
	unvisited := make(map[string]*Vertex, 0)
	visited := make(map[string]*Vertex, 0)
	pTable := make(PathTable)

	for k, v := range graph {
		var sdistance int64

		if k == start {
			sdistance = 0
		} else {
			sdistance = MaxInt
		}

		pTable[k] = &vertexregister{vertex: v, shortDistance: sdistance, previous: nil}
		unvisited[v.Name] = v
	}

	current := start
	u := graph[current]
	for len(unvisited) > 0 {
		neightbors := u.Edges

		for _, n := range filterVisited(visited, neightbors) {
			tableRegisterA := pTable[current]
			tableRegisterB := pTable[n.Point.Name]
			newDistance := tableRegisterA.shortDistance + int64(n.Weight)

			if newDistance < tableRegisterB.shortDistance {
				tableRegisterB.shortDistance = newDistance
				tableRegisterB.previous = tableRegisterA.vertex
			}
		}

		u = getMinorVertex(pTable, visited)
		current = u.Name
		delete(unvisited, current)
		visited[current] = u

	}

	out, err := rewind(pTable, start, end)

	if err != nil {
		return nil, 0, err

	}
	return out, pTable[end].shortDistance, nil
}
