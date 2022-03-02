package algorithms

func BFS(adjMatrix map[int][]int) []int {
	list := make([]int, 0)
	vistited := []int{}

	for node, neighbors := range adjMatrix {
		for _, neighbor := range neighbors {
			if notIn(neighbor, vistited) && notIn(neighbor, list) {
				list = append(list, neighbor)
			}
			list = append(list, neighbor)

		}
		vistited = append(vistited, node)
	}
	return list
}

func notIn(value int, list []int) bool {
	for _, item := range list {
		if item == value {
			return false
		}
	}
	return true
}
