package repository

import (
	"fmt"
	"reflect"
	"testing"
	"triproute/pkg/graph"
	"triproute/pkg/test"
)

const tempFile = "/tmp/test.csv"

func TestNewRepository(t *testing.T) {
	type args struct {
		dbFile string
	}
	tests := []struct {
		name        string
		args        args
		want        *Repository
		invalidFile bool
		fileContent string
		wantErr     bool
	}{
		{
			name: "Repository from dbFile",
			args: args{dbFile: tempFile},
			want: &Repository{
				DbFile: tempFile,
				Graph: map[string]*graph.Vertex{
					"GRU": &graph.Vertex{Name: "GRU", Edges: []*graph.Edge{&graph.Edge{Weight: 12, Point: &graph.Vertex{Name: "BRC", Edges: nil}}}},
					"BRC": &graph.Vertex{Name: "BRC", Edges: nil},
				},
			},
			invalidFile: false,
			wantErr:     false,
			fileContent: "GRU,BRC,12\n",
		},
		{
			name:        "Repository invalid dbFile",
			args:        args{dbFile: tempFile},
			want:        nil,
			invalidFile: true,
			wantErr:     true,
			fileContent: "",
		},
	}
	for _, tt := range tests {
		testFile := tempFile

		if tt.invalidFile == true {
			testFile = "/invalidpath/invalidfile.csv"
		} else {
			f := test.NewMockFile(testFile, tt.fileContent)
			defer f.ClearFile()
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRepository(testFile)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_String(t *testing.T) {

	type fields struct {
		DbFile string
	}
	tests := []struct {
		fields      fields
		name        string
		want        string
		fileContent string
	}{
		{
			name:        "Testing debug string",
			fields:      fields{DbFile: tempFile},
			want:        "<Repository file: /tmp/test.csv, graph: map[BRC:<Vertex: \"BRC\", Edges: []> GRU:<Vertex: \"GRU\", Edges: [<Edge Weight: 12, Point: <Vertex: \"BRC\", Edges: []>>]>]>",
			fileContent: "GRU,BRC,12",
		},
	}
	for _, tt := range tests {
		f := test.NewMockFile(tempFile, tt.fileContent)
		defer f.ClearFile()

		t.Run(tt.name, func(t *testing.T) {

			r, _ := NewRepository(tempFile)

			if got := r.String(); got != tt.want {
				t.Errorf("Repository.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_FindRoute(t *testing.T) {
	type fields struct {
		Graph map[string]*graph.Vertex
	}
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*graph.Vertex
		want1   int64
		wantErr bool
	}{
		{
			name: "Find basic route",
			args: args{start: "BRC", end: "SCL"},
			fields: fields{
				Graph: map[string]*graph.Vertex{
					"BRC": &graph.Vertex{Name: "BRC", Edges: []*graph.Edge{&graph.Edge{Weight: 12, Point: &graph.Vertex{Name: "SCL", Edges: nil}}}},
					"SCL": &graph.Vertex{Name: "SCL", Edges: nil},
				},
			},
			want:    []*graph.Vertex{&graph.Vertex{Name: "BRC", Edges: []*graph.Edge{&graph.Edge{Weight: 12, Point: &graph.Vertex{Name: "SCL", Edges: nil}}}}, &graph.Vertex{Name: "SCL", Edges: nil}},
			want1:   12,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{
				DbFile: tempFile,
				Graph:  tt.fields.Graph,
			}
			got, got1, err := r.FindRoute(tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.FindRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.FindRoute() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Repository.FindRoute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRepository_InsertRoute(t *testing.T) {
	type fields struct {
		DbFile string
	}
	type args struct {
		start  string
		end    string
		weight int64
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      bool
		fileContent  string
		invalidFile  bool
		wantGraphLen int
	}{
		{
			name:         "Insert simple route",
			fields:       fields{DbFile: tempFile},
			args:         args{start: "BRC", end: "CGD", weight: int64(37)},
			wantErr:      false,
			wantGraphLen: 2,
			invalidFile:  false,
		},
		{
			name:         "Wrong databse input",
			fields:       fields{DbFile: tempFile},
			args:         args{start: "BRC,\",", end: "CGD", weight: int64(37)},
			wantErr:      true,
			invalidFile:  false,
			wantGraphLen: 0,
		},
	}
	for _, tt := range tests {
		testFile := tempFile
		if tt.invalidFile == true {
			testFile = "/invalidpath/invalidfile.csv"
		} else {
			f := test.NewMockFile(tempFile, tt.fileContent)
			defer f.ClearFile()
		}

		t.Run(tt.name, func(t *testing.T) {
			r, errw := NewRepository(testFile)
			if errw != nil {
				fmt.Println(errw)
			}

			if err := r.InsertRoute(tt.args.start, tt.args.end, tt.args.weight); (err != nil) != tt.wantErr {
				t.Errorf("Repository.InsertRoute() error = %v, wantErr %v", err, tt.wantErr)
			}

			if len(r.Graph) != tt.wantGraphLen {
				t.Errorf("Repository.InsertRoute(), len(r.Graph) = %v, got %v", tt.wantGraphLen, len(r.Graph))
			}
		})

	}
}
