package controller

import (
	"bytes"
	"fmt"
	"triproute/pkg/graph"
)

// FormatResponse output the route in a readable format
func FormatResponse(out []*graph.Vertex, cost int64) string {
	var response bytes.Buffer

	for _, v := range out {
		response.WriteString(fmt.Sprintf("%s - ", v.Name))
	}

	response.WriteString(fmt.Sprintf("%d", cost))

	return response.String()
}
