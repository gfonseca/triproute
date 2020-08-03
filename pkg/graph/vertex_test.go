package graph

import (
	"reflect"
	"testing"
)

func TestNewVertex(t *testing.T) {
	type args struct {
		name  string
		edges []*Edge
	}
	tests := []struct {
		name string
		args args
		want *Vertex
	}{
		{
			name: "Test make name empty",
			args: args{"", nil},
			want: &Vertex{"", nil},
		},
		{
			name: "Test make name A",
			args: args{"A", nil},
			want: &Vertex{"A", nil},
		},
		{
			name: "Test make new Edge",
			args: args{"A", []*Edge{&Edge{}}},
			want: &Vertex{
				"A",
				[]*Edge{&Edge{}},
			},
		},
		{
			name: "Test make many edges",
			args: args{"A", []*Edge{&Edge{}, &Edge{}}},
			want: &Vertex{
				"A",
				[]*Edge{
					&Edge{},
					&Edge{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVertex(tt.args.name, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEdge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVertex_AddEdge(t *testing.T) {
	type fields struct {
		Name  string
		Edges []*Edge
	}

	type args struct {
		edge *Edge
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    int
	}{
		{
			name:    "Adding 1 Edge",
			fields:  fields{"A", nil},
			args:    args{&Edge{}},
			wantErr: false,
			want:    1,
		},
		{
			name:    "Adding 2 Edge",
			fields:  fields{"A", []*Edge{&Edge{}}},
			args:    args{&Edge{}},
			wantErr: false,
			want:    2,
		},
		{
			name:    "Fail adding  Edge",
			fields:  fields{"A", []*Edge{}},
			args:    args{nil},
			wantErr: true,
			want:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vertex{
				Name:  tt.fields.Name,
				Edges: tt.fields.Edges,
			}

			if err := v.AddEdge(tt.args.edge); (err != nil) != tt.wantErr {
				t.Errorf("Vertex.AddEdge() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got := len(v.Edges); got != tt.want {
				t.Errorf("len(v.Edges) = %v, want %d", got, tt.want)
			}
		})
	}
}

func TestVertex_String(t *testing.T) {
	type fields struct {
		Name  string
		Edges []*Edge
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Check for debug string",
			fields: fields{"VtxA", nil},
			want:   "<Vertex: \"VtxA\", Edges: []>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vertex{
				Name:  tt.fields.Name,
				Edges: tt.fields.Edges,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Vertex.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
