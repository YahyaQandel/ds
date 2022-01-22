package algorithms

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Test unsorted array",
			arr:  []int{1, 100, 3, 453, 234, 12, 15, 10},
			want: []int{1, 3, 10, 12, 15, 100, 234, 453},
		},
		{
			name: "Test sorted array",
			arr:  []int{1, 10, 20, 34, 56, 100},
			want: []int{1, 10, 20, 34, 56, 100},
		},
		{
			name: "Test unsorted array of duplicates",
			arr:  []int{89, 10, 10, 34, 10, 200, 4},
			want: []int{4, 10, 10, 10, 34, 89, 200},
		},
		{
			name: "Test sorted array of duplicates",
			arr:  []int{10, 10, 34, 100, 200, 200, 200},
			want: []int{10, 10, 34, 100, 200, 200, 200},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Quicksort(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quicksort() = %v, want %v", got, tt.want)
			}
		})
	}
}
