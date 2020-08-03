package repository

import (
	"fmt"
	"os"

	"triproute/pkg/graph"
)

// Repository is an data access object
type Repository struct {
	DbFile string
	Graph  map[string]*graph.Vertex
}

// NewRepository build and configure Repository instance
func NewRepository(dbFile string) (*Repository, error) {
	r := &Repository{DbFile: dbFile}
	err := r.LoadDb()

	if err != nil {
		return nil, err
	}

	return r, nil
}

// LoadDb load data from csv file
func (r *Repository) LoadDb() error {
	var err error
	r.Graph, err = graph.FromCSV(r.DbFile)

	if err != nil {
		return err
	}
	return nil
}

// FindRoute find the cheapest route
func (r Repository) FindRoute(start, end string) ([]*graph.Vertex, int64, error) {
	return graph.Dijkstra(r.Graph, start, end)
}

// InsertRoute write a new conection to DataBase
func (r *Repository) InsertRoute(start, end string, weight int64) error {
	fileHandler, _ := os.OpenFile(r.DbFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	_, err := fileHandler.WriteString(fmt.Sprintf("\n%s,%s,%d", start, end, weight))

	if err != nil {
		return fmt.Errorf("Failed to update database file: %s", err)
	}

	err = r.LoadDb()
	if err != nil {
		return fmt.Errorf("Failed to parse database file: %s", err)
	}

	return nil
}

func (r Repository) String() string {
	return fmt.Sprintf("<Repository file: %s, graph: %s>", r.DbFile, r.Graph)
}