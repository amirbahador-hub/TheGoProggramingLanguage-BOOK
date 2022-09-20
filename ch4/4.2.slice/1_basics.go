package main

import "fmt"

func main() {
	months := [...]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f", 7: "g"}

	fmt.Printf("type: %T, value: %s \n", months, months) // type: [8]string, value: [ a b c d e f g] -> type is array

	Q1 := months[4:7]
	Q2 := months[6:8] // pay attention in 8!

	fmt.Printf("type: %T, value: %s \n", Q1, Q1) // type: []string, value: [d e f] -> type is slice
	fmt.Printf("type: %T, value: %s \n", Q2, Q2) // type: []string, value: [f g] -> type is slice

	for _, q1 := range Q1 {
		for _, q2 := range Q2 {
			if q1 == q2 {
				fmt.Println(q1) // print f
			}
		}
	}

	fmt.Println(months[5:]) // [e f g]
	fmt.Println(months[:3]) // [ a b]
	fmt.Println(months[:8]) // [ a b c d r f g]
	//fmt.Println(months[:9]) // invalid argument: array index 9 out of bounds [0:9]

}
