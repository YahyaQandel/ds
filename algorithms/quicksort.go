package algorithms

func Quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	less := []int{}
	greater := []int{}
	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		}
	}
	for _, v := range arr[1:] {
		if v > pivot {
			greater = append(greater, v)
		}
	}
	res := []int{}
	res = append(res, Quicksort(less)...)
	res = append(res, pivot)
	res = append(res, Quicksort(greater)...)
	return res
}
