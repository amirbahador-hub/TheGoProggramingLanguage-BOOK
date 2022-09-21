package main

import "fmt"

func equalSlice(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func equalArray(x, y [6]int) bool {
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := []int{0, 1, 2, 3, 4, 5}
	c := []int{-1, 1, 2, 3, 4, 5}
	fmt.Println(equalSlice(a, b)) //	true
	fmt.Println(equalSlice(a, c)) //	false

	e := [6]int{0, 1, 2, 3, 4, 5}
	f := [6]int{0, 1, 2, 3, 4, 5}
	g := [6]int{-1, 1, 2, 3, 4, 5}

	fmt.Println(equalArray(e, f)) //	true
	fmt.Println(equalArray(f, g)) //	false

	var s []int
	fmt.Printf("value is : %d, length is : %d, equal to nil : %t \n", s, len(s), s == nil) //	value is : [], length is : 0, equal to nil : true

	s = nil
	fmt.Printf("value is : %d, length is : %d, equal to nil : %t \n", s, len(s), s == nil) //	value is : [], length is : 0, equal to nil : true

	s = []int(nil)
	fmt.Printf("value is : %d, length is : %d, equal to nil : %t \n", s, len(s), s == nil) //	value is : [], length is : 0, equal to nil : true

	s = []int{}
	fmt.Printf("value is : %d, length is : %d, equal to nil : %t \n", s, len(s), s == nil) //	value is : [], length is : 0, equal to nil : false
}
