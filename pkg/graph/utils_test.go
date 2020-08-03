package graph

import (
	"os"
	"reflect"
	"testing"
	"triproute/pkg/test"
)

const tempFile = "/tmp/testefile.csv"

func TestFromCSV(t *testing.T) {
	os.Remove(tempFile)
	type args struct {
	}
	tests := []struct {
		name        string
		args        args
		want        map[string]*Vertex
		mockFile    func(path string) test.MockFile
		wantErr     bool
		fileContent string
		invalidFile bool
	}{
		{
			name: "Load csv from File",
			want: map[string]*Vertex{
				"GRU": &Vertex{
					Name:  "GRU",
					Edges: []*Edge{&Edge{Weight: 12, Point: &Vertex{Name: "BRC", Edges: nil}}},
				},
				"BRC": &Vertex{
					Name:  "BRC",
					Edges: nil,
				},
			},
			fileContent: "GRU,BRC,12",
			invalidFile: false,
		},
		{
			name: "Load one vertice two destinations",
			want: map[string]*Vertex{
				"GRU": &Vertex{
					Name: "GRU",
					Edges: []*Edge{
						&Edge{Weight: 12, Point: &Vertex{Name: "BRC", Edges: nil}},
						&Edge{Weight: 33, Point: &Vertex{Name: "SCL", Edges: nil}},
					},
				},
				"BRC": &Vertex{
					Name:  "BRC",
					Edges: nil,
				},
				"SCL": &Vertex{
					Name:  "SCL",
					Edges: nil,
				},
			},
			fileContent: "GRU,BRC,12\nGRU,SCL,33",
			invalidFile: false,
		},
		{
			name: "Load two vertices two destinations",
			want: map[string]*Vertex{
				"GRU": &Vertex{
					Name: "GRU",
					Edges: []*Edge{
						&Edge{Weight: 12, Point: &Vertex{Name: "SCL", Edges: nil}},
					},
				},
				"BRC": &Vertex{
					Name: "BRC",
					Edges: []*Edge{
						&Edge{Weight: 33, Point: &Vertex{Name: "SCL", Edges: nil}},
					},
				},
				"SCL": &Vertex{
					Name:  "SCL",
					Edges: nil,
				},
			},
			fileContent: "GRU,SCL,12\nBRC,SCL,33",
			invalidFile: false,
		},
		{
			name: "Load csv from File 2 edges",
			want: map[string]*Vertex{
				"GRU": &Vertex{
					Name: "GRU",
					Edges: []*Edge{&Edge{Weight: 12, Point: &Vertex{Name: "BRC", Edges: []*Edge{
						&Edge{
							Weight: 45,
							Point:  &Vertex{Name: "SCL", Edges: nil},
						},
					},
					}}},
				},
				"BRC": &Vertex{
					Name: "BRC",
					Edges: []*Edge{
						&Edge{
							Weight: 45,
							Point:  &Vertex{Name: "SCL", Edges: nil},
						},
					},
				},
				"SCL": &Vertex{Name: "SCL", Edges: nil},
			},
			fileContent: "GRU,BRC,12\nBRC,SCL,45",
			invalidFile: false,
		},
		{
			name:        "Test csv not exits",
			wantErr:     true,
			invalidFile: true,
		},
		{
			name:        "Test invalid csv format",
			wantErr:     true,
			invalidFile: false,
			fileContent: "GRU,BRC,invalidstring",
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
			got, err := FromCSV(testFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}
