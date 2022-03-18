package algorithms

import "testing"

func TestDijkstra(t *testing.T) {
	type args struct {
		adjMatrix map[int][]int
		target    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "search the cheapest path",
			args: args{adjMatrix: map[int][]int{1: []int{2, 3, 4}, 2: []int{7, 8}, 3: []int{5, 6}, 4: []int{6}}, target: 8},
			want: true,
		},
		{
			name: "search for non existing item in graph",
			args: args{adjMatrix: map[int][]int{1: []int{2, 3, 4}, 2: []int{7, 8}, 3: []int{5, 6}, 4: []int{6}}, target: 0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BFS(tt.args.adjMatrix, tt.args.target); got != tt.want {
				t.Errorf("Dijkstra() = %v, want %v", got, tt.want)
			}
		})
	}
}
