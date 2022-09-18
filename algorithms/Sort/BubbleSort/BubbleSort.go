package main

import (
	"fmt"
	"time"
)

//bubble sort is check repeatedly swapping adjustment element
//this is not suitable for large array
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ { //	n-i-1 because after every value check, don't need check again of last elements
			if arr[j] > arr[j+1] { // check next value is bigger or not
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	totalStart := time.Now()
	arr := []int{33, 98, 74, 3, 13, 55, 20, 10, 2, 0, 14, 43, 25, 18, 1, 5, 45, 77, 6, 45, 64, 83, 8, 7}
	fmt.Println(bubbleSort(arr))
	elapsedTotal := time.Since(totalStart)
	fmt.Println("total time elapsed: ", elapsedTotal) //	42.5514µs
}
