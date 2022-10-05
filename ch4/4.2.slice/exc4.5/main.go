package main

import "fmt"

func removeDuplicateAdjacent(arr []string) []string {
	j := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[j] {
			j++
			arr[j] = arr[i]
		}
	}
	return arr[:j+1]
}

func main() {
	arr := []string{"a", "b", "c", "c", "a", "c", "b", "b"}
	fmt.Println(removeDuplicateAdjacent(arr))
}
