package main

import (
	"binarysearch"
	"fmt"
)

func main() {
	values1 := []int{2, 5, 7, 8, 11, 15, 25, 82, 102}
	index1 := binaryseasrch.BinarySearch(0, len(values1)-1, 11, values1)
	fmt.Printf("index is %d and value is %d \n", index1, values1[index1])

	values2 := [...]int{2, 5, 7, 8, 11, 15, 25, 82, 102}
	index2 := binaryseasrch.BinarySearch(0, len(values2)-1, 11, values2[:])
	fmt.Printf("index is %d and value is %d \n", index2, values2[index2])
}
