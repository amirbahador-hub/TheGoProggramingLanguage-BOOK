package main

import "fmt"

func binarySearch(arr []int, value int, defaultIndex int) int {
	mid := len(arr) / 2
	if arr[mid] == value {
		return mid + defaultIndex
	} else if value < arr[mid] {
		return binarySearch(arr[:mid], value, defaultIndex)
	} else {
		defaultIndex += mid
		return binarySearch(arr[mid:], value, defaultIndex)
	}
}

func main() {
	arr := []int{1, 3, 5, 8, 11, 13, 14, 19, 27}
	fmt.Println(binarySearch(arr, 1, 0))
	fmt.Println(binarySearch(arr, 5, 0))
	fmt.Println(binarySearch(arr, 8, 0))
	fmt.Println(binarySearch(arr, 11, 0))
	fmt.Println(binarySearch(arr, 13, 0))
	fmt.Println(binarySearch(arr, 19, 0))
	fmt.Println(binarySearch(arr, 27, 0))
}
