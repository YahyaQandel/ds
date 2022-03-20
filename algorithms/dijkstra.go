package algorithms

import "reflect"

const INF = 99999999999

func Dijkstra(graph map[string]map[string]int, start string, target string) (int, []string) {
	/*
		Dijkstra steps:
		- find lowest/cheapest node to the start (ex: x)
		- loop through its neighbours and get their costs by (cost to x) + cost to neighbour
		- update the neigjbour cost if it is lower than the saved cost and update its parent to be x
		- process node as done
		- repeat again the steps until nodes
	*/
	processedNodes := make([]string, 0)
	costs, parents := getGraphNodesCostsAndParents(graph, start)
	// parents := make(map[string]string)
	for {

		nearestNode := getNearestNodeNotProcessed(costs, processedNodes)
		if nearestNode == "" {
			break
		}
		nearestNodeCost := costs[nearestNode]
		for neighbour, cost := range graph[nearestNode] {
			newNeighbourNodeCost := nearestNodeCost + cost
			if newNeighbourNodeCost < costs[neighbour] {
				costs[neighbour] = newNeighbourNodeCost
				parents[neighbour] = nearestNode
			}
		}
		processedNodes = append(processedNodes, nearestNode)
	}
	path := []string{}
	indexNode := target
	path = append(path, indexNode)
	for {
		indexParent := parents[indexNode]
		path = append(path, indexParent)
		if indexParent == start {
			break
		}
		indexNode = indexParent
	}
	return costs[target], reverse(path)
}

func getNearestNodeNotProcessed(costs map[string]int, processedNodes []string) string {
	lowestCost := INF
	lowestCostNode := ""
	for node, cost := range costs {
		if cost < int(lowestCost) && nodeNotIn(node, processedNodes) {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

func nodeNotIn(value string, list []string) bool {
	for _, item := range list {
		if item == value {
			return false
		}
	}
	return true
}

func getGraphNodesCostsAndParents(graph map[string]map[string]int, startNode string) (map[string]int, map[string]string) {

	costs := make(map[string]int)
	parents := make(map[string]string)

	// set all nodes costs to inf
	for node, _ := range graph {
		if node != startNode {
			costs[node] = INF
		}
	}
	for neighbour, cost := range graph[startNode] {
		costs[neighbour] = cost
	}
	for node, neighbours := range graph {
		for neighbour, _ := range neighbours {
			parents[neighbour] = node
		}
	}
	return costs, parents
}

func reverse(s []string) []string {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	return s
}
