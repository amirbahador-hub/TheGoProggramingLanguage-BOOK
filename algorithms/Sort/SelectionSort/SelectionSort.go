package main

import (
	"fmt"
	"time"
)

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
	totalStart := time.Now()
	arr := []int{33, 98, 74, 3, 13, 55, 20, 10, 2, 0, 14, 43, 25, 18, 1, 5, 45, 77, 6, 45, 64, 83, 8, 7}
	fmt.Println(SelectionSort(arr))
	elapsedTotal := time.Since(totalStart)
	fmt.Println("total time elapsed: ", elapsedTotal) //	45.3254µs
}
