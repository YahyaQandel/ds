package dc_problems

func sum_of(list []int) int {
	return sum(list)
}

func sum(list []int) int {
	if len(list) == 1 {
		return list[0]
	}
	return list[0] + sum(list[1:])
}

// another implementation

func sum_of_list_with_variable(list []int) int {
	return sum_with_variable(list)
}

func sum_with_variable(list []int) int {
	if len(list) == 0 {
		return 0
	}
	sm := sum_with_variable(list[1:])
	sm += list[0]
	return sm
}
