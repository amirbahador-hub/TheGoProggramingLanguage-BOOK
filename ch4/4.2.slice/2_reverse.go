package main

import "fmt"

func reverseSlice(sli []int) {
	for i, j := 0, len(sli)-1; i < j; i, j = i+1, j-1 {
		sli[i], sli[j] = sli[j], sli[i]
	}
}

func reverseArray(arr *[2]int) {
	for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	b := [...]int{0, 1, 2, 3, 4, 5}

	//reverseSlice(a) // cannot use a (variable of type [6]int) as type []int in argument to reverseSlice

	//rotate two-step in left
	reverseSlice(a[:2])
	reverseSlice(a[2:])
	reverseSlice(a[:])
	fmt.Println(a) //	[2 3 4 5 0 1]

	//reverseArray(a[:2]) // cannot use a[:2] (value of type []int) as type [6]int in argument to reverseArray

	//rotate two-step in left

	var bArray [2]int
	copy(bArray[:], b[:2])
	fmt.Println(b)                           //	[0 1 2 3 4 5]
	fmt.Printf("%T , %d \n", bArray, bArray) //	[2]int , [0 1] -> bArray type is array
	fmt.Println(bArray)                      //	[0 1]
	reverseArray(&bArray)
	fmt.Println(bArray) //	[1 0]
}
