package main

import (
	"fmt"
)

func binarySearh(arr []int, value int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if value == arr[mid] {
			return mid
		} else if value > arr[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return -1
}

func main() {
	fmt.Println(binarySearh([]int{10, 20, 30, 40, 50}, 10))
	fmt.Println(binarySearh([]int{10, 20, 30, 40, 50}, 20))
	fmt.Println(binarySearh([]int{10, 20, 30, 40, 50}, 30))
	fmt.Println(binarySearh([]int{10, 20, 30, 40, 50}, 40))
	fmt.Println(binarySearh([]int{10, 20, 30, 40, 50}, 50))
	fmt.Println(binarySearh([]int{10, 20, 30, 40, 50}, 60))
}
