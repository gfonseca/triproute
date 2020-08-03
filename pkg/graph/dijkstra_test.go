package graph

import (
	"reflect"
	"testing"
)

func flatResponse(vtxs []*Vertex) []string {
	out := make([]string, 0)

	for _, v := range vtxs {
		out = append(out, v.Name)
	}

	return out
}

func TestDijkstra(t *testing.T) {
	type args struct {
		graph map[string]*Vertex
		start string
		end   string
	}
	tests := []struct {
		name      string
		args      args
		want      []string
		want1     int64
		wantErr   bool
		makeGraph func() map[string]*Vertex
	}{
		{
			name: "Testing invalide destiny",
			args: args{
				start: "GRU",
				end:   "RTS",
			},
			wantErr: true,
			makeGraph: func() map[string]*Vertex {
				gru := NewVertex("GRU", []*Edge{})
				scl := NewVertex("SCL", []*Edge{})

				return map[string]*Vertex{
					"GRU": gru,
					"SCL": scl,
				}
			},
			want:  []string{},
			want1: 0,
		},
		{
			name: "Testing invalide origin",
			args: args{
				start: "DRF",
				end:   "GRU",
			},
			wantErr: true,
			makeGraph: func() map[string]*Vertex {
				gru := NewVertex("GRU", []*Edge{})
				scl := NewVertex("SCL", []*Edge{})

				return map[string]*Vertex{
					"GRU": gru,
					"SCL": scl,
				}
			},
			want:  []string{},
			want1: 0,
		},
		{
			name: "Testing invalide route",
			args: args{
				start: "SCL",
				end:   "GRU",
			},
			wantErr: true,
			makeGraph: func() map[string]*Vertex {
				gru := NewVertex("GRU", []*Edge{})
				scl := NewVertex("SCL", []*Edge{})

				return map[string]*Vertex{
					"GRU": gru,
					"SCL": scl,
				}
			},
			want:  []string{},
			want1: 0,
		},
		{
			name: "Testing basic route",
			args: args{
				start: "SCL",
				end:   "GRU",
			},
			wantErr: false,
			makeGraph: func() map[string]*Vertex {
				gru := NewVertex("GRU", []*Edge{})
				scl := NewVertex("SCL", []*Edge{})

				scl.AddEdge(&Edge{Weight: 20, Point: gru})

				return map[string]*Vertex{
					"GRU": gru,
					"SCL": scl,
				}
			},
			want:  []string{"SCL", "GRU"},
			want1: 20,
		},
		{
			name: "Testing basic route",
			args: args{
				start: "SCL",
				end:   "GRU",
			},
			wantErr: false,
			makeGraph: func() map[string]*Vertex {
				gru := NewVertex("GRU", []*Edge{})
				scl := NewVertex("SCL", []*Edge{})
				orl := NewVertex("ORL", []*Edge{})
				brc := NewVertex("BRC", []*Edge{})

				gru.AddEdge(&Edge{Weight: 12, Point: scl})
				scl.AddEdge(&Edge{Weight: 20, Point: brc})
				brc.AddEdge(&Edge{Weight: 30, Point: gru})
				brc.AddEdge(&Edge{Weight: 5, Point: orl})
				orl.AddEdge(&Edge{Weight: 4, Point: gru})
				orl.AddEdge(&Edge{Weight: 10, Point: scl})

				return map[string]*Vertex{
					"GRU": gru,
					"SCL": scl,
					"ORL": orl,
					"BRC": brc,
				}
			},
			want:  []string{"SCL", "BRC", "ORL", "GRU"},
			want1: 29,
		},
	}
	for _, tt := range tests {

		graph := tt.makeGraph()
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Dijkstra(graph, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dijkstra() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(flatResponse(got), tt.want) {
				t.Errorf("Dijkstra():path got = %v, want %v", got, tt.want)
			}

			if got1 != tt.want1 {
				t.Errorf("Dijkstra():dist got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
