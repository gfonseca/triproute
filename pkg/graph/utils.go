package graph

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// FromCSV transform an csv file in a Vertex/Edges structure called Graph
func FromCSV(filePath string) (map[string]*Vertex, error) {
	var vtxb, vtxa *Vertex
	vtxs := make(map[string]*Vertex)
	csvfile, err := os.Open(filePath)
	defer csvfile.Close()

	if err != nil {
		return nil, fmt.Errorf("Couldn't open the csv file: %s", err)
	}

	r := csv.NewReader(csvfile)

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("Couldn't read the csv file: %s", err)
		}

		va := record[0]
		vb := record[1]
		weight, err := strconv.Atoi(record[2])

		if err != nil {
			return nil, fmt.Errorf("Failed to parse edge weight: %s", err)
		}

		if val, ok := vtxs[va]; ok {
			vtxa = val
		} else {
			vtxa = NewVertex(va, nil)
			vtxs[va] = vtxa
		}

		if val, ok := vtxs[vb]; ok {
			vtxb = val
		} else {
			vtxb = NewVertex(vb, nil)
			vtxs[vb] = vtxb
		}

		edge := NewEdge(weight, vtxb)

		vtxa.AddEdge(edge)
	}

	return vtxs, nil
}
