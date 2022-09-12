package main

import (
	"binarysearch"
	"fmt"
)

func main() {
	values := [...]int{2, 5, 7, 8, 11, 15, 25, 82, 102}
	index := binaryseasrch.BinarySearch(0, len(values)-1, 11, values)
	fmt.Printf("index is %d and value is %d", index, values[index])
}
