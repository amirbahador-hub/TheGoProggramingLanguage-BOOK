package main

import (
	"fmt"
	"time"
)

//	this is not suitable for large array
//	it's repeatably like sort card
func InsertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	totalStart := time.Now()
	arr := []int{33, 98, 74, 3, 13, 55, 20, 10, 2, 0, 14, 43, 25, 18, 1, 5, 45, 77, 6, 45, 64, 83, 8, 7}
	fmt.Println(InsertionSort(arr))
	elapsedTotal := time.Since(totalStart)
	fmt.Println("total time elapsed: ", elapsedTotal) //	32.0584µs
}
