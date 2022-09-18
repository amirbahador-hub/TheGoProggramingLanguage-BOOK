package main

import "fmt"

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
	arr := []int{8, 5, 3, 7, 6}
	fmt.Println(bubbleSort(arr))
}
