package algorithms

import (
	"reflect"
	"testing"
)

type test struct {
	input []int
	value int
	want  int
}

func TestBs(t *testing.T) {
	sampleArray := []int{1, 4, 6, 8, 10, 12}
	tests := []test{
		{input: sampleArray, value: 1, want: 0},
		{input: sampleArray, value: 4, want: 1},
		{input: sampleArray, value: 6, want: 2},
		{input: sampleArray, value: 8, want: 3},
		{input: sampleArray, value: 10, want: 4},
		{input: sampleArray, value: 12, want: 5},
		{input: sampleArray, value: 200, want: -1},
	}
	for _, tc := range tests {
		got := BinarySearh(tc.input, tc.value)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
