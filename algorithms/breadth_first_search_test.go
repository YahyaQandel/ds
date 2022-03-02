package algorithms

import (
	"fmt"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	adjMatrix := make(map[int][]int)
	// check example at example.MD
	adjMatrix[1] = []int{2, 3, 4}
	adjMatrix[2] = []int{7, 8}
	adjMatrix[3] = []int{5, 6}
	adjMatrix[4] = []int{6}
	list := BFS(adjMatrix)
	fmt.Print(list)
}
