package algorithms

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	type args struct {
		graph  map[string]map[string]int
		start  string
		target string
	}
	tests := []struct {
		name     string
		args     args
		wantCost int
		wantPath []string
	}{
		{
			name: "search the cheapest path to item",
			args: args{graph: map[string]map[string]int{
				"start": {
					"a": 6,
					"b": 2,
				},
				"a": {
					"fin": 1,
				},
				"b": {
					"a":   3,
					"fin": 5,
				},
				"fin": {},
			},
				start:  "start",
				target: "fin"},
			wantCost: 6,
			wantPath: []string{"start", "b", "a", "fin"},
		},
		{
			name: "search the cheapest path to get the piano",
			args: args{graph: map[string]map[string]int{
				"book": {
					"lp":     5,
					"poster": 0,
				},
				"lp": {
					"bass guitar": 15,
					"drums":       20,
				},
				"poster": {
					"bass guitar": 30,
					"drums":       35,
				},
				"bass guitar": {
					"piano": 20,
				},
				"drums": {
					"piano": 10,
				},
				"piano": {},
			},
				start:  "book",
				target: "piano"},
			wantCost: 35,
			wantPath: []string{"book", "lp", "drums", "piano"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCost, gotPath := Dijkstra(tt.args.graph, tt.args.start, tt.args.target)
			if gotCost != tt.wantCost {
				t.Errorf("Dijkstra() = %v, wanted cost %v", gotCost, tt.wantCost)
			}
			if !reflect.DeepEqual(gotPath, tt.wantPath) {
				t.Errorf("Dijkstra() = %v, wanted path %v", gotPath, tt.wantPath)
			}
		})
	}
}
