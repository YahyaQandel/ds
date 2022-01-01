package algorithms

func BinarySearh(arr []int, value int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if value == arr[mid] {
			return mid
		} else if value > arr[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
