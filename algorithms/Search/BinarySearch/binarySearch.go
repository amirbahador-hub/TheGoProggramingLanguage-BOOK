package main

import "fmt"

func binarySearch(arr []int, value int, defaultIndex ...int) int {
	var defaultIndexValue int
	if len(defaultIndex) == 0 {
		defaultIndexValue = 0
	} else {
		defaultIndexValue = defaultIndex[0]
	}
	mid := len(arr) / 2
	if arr[mid] == value {
		return mid + defaultIndexValue
	} else if value < arr[mid] {
		return binarySearch(arr[:mid], value, defaultIndexValue)
	} else {
		defaultIndexValue += mid
		return binarySearch(arr[mid:], value, defaultIndexValue)
	}
}

func main() {
	arr := []int{1, 3, 5, 8, 11, 13, 14, 19, 27}
	fmt.Println(binarySearch(arr, 1))
	fmt.Println(binarySearch(arr, 5))
	fmt.Println(binarySearch(arr, 8))
	fmt.Println(binarySearch(arr, 11))
	fmt.Println(binarySearch(arr, 13))
	fmt.Println(binarySearch(arr, 19))
	fmt.Println(binarySearch(arr, 27))
}
