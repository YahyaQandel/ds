package dc_problems

func findMax(arr []int) int {
	return getMax(arr, -9999999999)
}

func getMax(arr []int, max int) int {
	if max < arr[0] {
		max = arr[0]
	}
	if len(arr) == 1 {
		return max
	}
	return getMax(arr[1:], max)
}
