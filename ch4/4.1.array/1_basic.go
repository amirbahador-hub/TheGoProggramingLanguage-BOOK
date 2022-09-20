package main

import "fmt"

func main() {
	var a [3]int

	fmt.Println(a[0])        // return zero value
	fmt.Println(a[len(a)-1]) //	return zero value

	for key, value := range a {
		fmt.Printf("key: %d, value: %d \n", key, value) // return zero value and key of array like 1, 2, 3
	}

	for _, value := range a {
		fmt.Printf("value: %d \n", value) // return only zero value of int array
	}

	var b [3]int = [3]int{1, 2}
	//b = [3]int{1, 2} // this is true
	//b = [4]int{1, 2} // cannot use [4]int{…} (value of type [4]int) as type [3]int in assignment
	//b = [4]int{1, 2, 3, 4} // cannot use [4]int{…} (value of type [4]int) as type [3]int in assignment

	fmt.Println(b[2]) // return zero value of int `0`

	var c = [...]int{1, 2, 3} // with ellipsis, generate array size with length of array associate

	fmt.Printf("type of this variable is : %T \n", c) // return [3]int

	var d = [...]int{5: -1, 8: 2}

	for key, value := range d {
		fmt.Printf("key: %d, value: %d \n", key, value) // generate all element with zero value expect 5th and 8th element and print -1
	}
}
