package main

import "fmt"

func rotate(arr []int, shift int) {
	shift = shift % len(arr)
	var temp = make([]int, shift)
	copy(temp, arr[:shift])
	copy(arr, arr[shift:])
	copy(arr[(len(arr)-shift):], temp)
}

func main() {
	arr := []int{2, 3, 5, 8, 7}
	rotate(arr, 2)
	fmt.Println(arr)
}
