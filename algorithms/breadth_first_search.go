package algorithms

import (
	datastructres "github.com/YahyaQandel/ds/datastructures"
)

func BFS(adjMatrix map[int][]int, target int) bool {
	queue := datastructres.Queue{}
	queue.Enqueue(1)
	listToWalkTo := []int{}
	visited := []int{}
	for queue.Size() > 0 {
		currentNode := queue.Dequeue()
		if currentNode == target {
			return true
		}
		if notIn(currentNode, visited) {
			for i := 0; i < len(adjMatrix[currentNode]); i++ {
				queue.Enqueue(adjMatrix[currentNode][i])
			}
			listToWalkTo = append(listToWalkTo, currentNode)
			visited = append(visited, currentNode)
		}
	}
	return false
}

func notIn(value int, list []int) bool {
	for _, item := range list {
		if item == value {
			return false
		}
	}
	return true
}
