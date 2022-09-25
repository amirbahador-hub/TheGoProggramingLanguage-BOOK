package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	point1 := point{1, 2}
	point2 := point{2, 1}
	point3 := point{1, 2}

	fmt.Println(point1 == point2) //	false
	fmt.Println(point1 == point3) //	true
	fmt.Println(point2 == point3) //	false

	navigateMap := make(map[point]int)

	navigateMap[point1]++
	navigateMap[point2]++
	navigateMap[point3]++
	navigateMap[point3]++

	fmt.Println(navigateMap) //	map[{1 2}:3 {2 1}:1]

	navigateMap2 := make(map[point]int)

	navigateMap2[point1]++
	navigateMap2[point2]++
	navigateMap2[point3]++
	navigateMap2[point3]++

	fmt.Println(navigateMap2) //	map[{1 2}:3 {2 1}:1]

	navigateMap3 := make(map[point]int)

	navigateMap3[point1]++
	navigateMap3[point2]++
	navigateMap3[point2]++
	navigateMap3[point3]++

	fmt.Println(navigateMap3) //	map[{1 2}:2 {2 1}:2]

	//fmt.Println(navigateMap == navigateMap2)  //	invalid operation: navigateMap == navigateMap2 (map can only be compared to nil)
	//fmt.Println(navigateMap == navigateMap3)  //	invalid operation: navigateMap == navigateMap2 (map can only be compared to nil)
	//fmt.Println(navigateMap2 == navigateMap3) //	invalid operation: navigateMap == navigateMap2 (map can only be compared to nil)

	fmt.Println(equal(navigateMap, navigateMap2))  //	true
	fmt.Println(equal(navigateMap, navigateMap3))  //	false
	fmt.Println(equal(navigateMap2, navigateMap3)) // 	false
}

func equal(x, y map[point]int) bool {
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
