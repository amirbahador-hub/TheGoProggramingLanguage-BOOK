package main

import "fmt"

func binarySearch(left int, right int, x int, array []int) int {
	if left > right {
		return 0
	} else {
		mid := (left + right) / 2
		if array[mid] == x {
			return mid
		} else if array[mid] > x {
			return binarySearch(mid+1, right, x, array)
		} else {
			return binarySearch(left, mid-1, x, array)
		}
	}
}

func appendInt(x []int, y ...int) []int {
	var z []int
	//zlen := len(x) + len(y)
	// ...expand z to at least zlen...\
	fmt.Printf("%T \n", y)
	fmt.Println(len(y))
	fmt.Println(cap(y))
	copy(z[len(x):], y)
	return z
}

var m = make(map[string]int)

func k(list []string) string  { return fmt.Sprintf("%q", list) }
func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

func main() {
	var slicec []string
	slicec = append(slicec, "ali")
	slicec = append(slicec, "amir")
	var slice2 []string
	slice2 = append(s)
	fmt.Println(k(slicec))
	Add(slicec)
	fmt.Println(Count(slicec))
	fmt.Println(slicec)
	fmt.Println(m)
}
