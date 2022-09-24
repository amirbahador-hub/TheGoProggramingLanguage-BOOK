package main

import "fmt"

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	} else {
		for key, value := range x {
			if _, ok := y[key]; !ok || value != y[key] {
				return false
			}
		}
		return true
	}
}

func main() {
	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))                  //	false
	fmt.Println(equal(map[string]int{"A": 0, "B": 42}, map[string]int{"A": 0, "B": 42})) //	true
}
