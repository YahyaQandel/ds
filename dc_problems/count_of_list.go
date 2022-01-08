package dc_problems

func count_of(list []int) int {
	return count(list, 0)
}

func count(list []int, cnt int) int {
	if len(list) == 0 {
		return cnt
	}
	cnt += 1
	return count(list[1:], cnt)
}
