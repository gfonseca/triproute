package graph

import (
	"reflect"
	"testing"
)

func TestNewEdge(t *testing.T) {
	type args struct {
		weight int
		point  *Vertex
	}
	tests := []struct {
		name string
		args args
		want *Edge
	}{
		{
			name: "Test make params = nil",
			args: args{0, nil},
			want: &Edge{0, nil},
		},
		{
			name: "Test make Weight = 2",
			args: args{2, nil},
			want: &Edge{2, nil},
		},
		{
			name: "Test make e.Point = Vertex",
			args: args{2, &Vertex{}},
			want: &Edge{2, &Vertex{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEdge(tt.args.weight, tt.args.point); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEdge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEdge_String(t *testing.T) {
	type fields struct {
		Weight int
		Point  *Vertex
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test for debug string",
			fields: fields{
				Weight: 12,
				Point:  nil,
			},
			want: "<Edge Weight: 12, Point: <nil>>",
		},
		{
			name: "Test for debug string with other fields",
			fields: fields{
				Weight: 12,
				Point:  &Vertex{"A", nil},
			},
			want: "<Edge Weight: 12, Point: <Vertex: \"A\", Edges: []>>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Edge{
				Weight: tt.fields.Weight,
				Point:  tt.fields.Point,
			}
			if got := e.String(); got != tt.want {
				t.Errorf("Edge.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

// <Edge Weight: 12, Point: <Vertex: "A", Edges: []>>
// <Edge Weight: 12, Point: <Vertex: "A" Edges: []>>
