package main

import "fmt"

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		pointer := i
		for j := i + 1; j < len(arr); j++ {
			if arr[pointer] > arr[j] { //check if next value is smaller than this value, pointer point to next value
				pointer = j
			}
		}
		arr[i], arr[pointer] = arr[pointer], arr[i] // swap arr index i and index pointer ti sort values
	}
	return arr
}

func main() {
	arr := []int{6, 5, 3, 7, 8}
	fmt.Println(SelectionSort(arr))
}
