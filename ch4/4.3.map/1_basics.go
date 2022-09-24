package main

import (
	"fmt"
	"sort"
)

func main() {
	//first way to define
	ages := make(map[string]int)

	ages["ali"] = 24
	ages["mohammad"] = 20

	fmt.Println(ages)        //	map[ali:24 mohammad:20]
	fmt.Println(ages["ali"]) //	24

	ages2 := map[string]int{
		"ali":      24,
		"mohammad": 20,
	}

	fmt.Println(ages2) //	map[ali:24 mohammad:20]

	delete(ages, "ali")
	fmt.Println(ages)           //	map[mohammad:20]
	fmt.Println(ages["ali"])    //	0
	fmt.Println(ages["abcdef"]) //	0

	ages["mohammad"] = ages["mohammad"] + 1
	fmt.Println(ages["mohammad"]) //	21
	ages["mohammad"] += 1
	fmt.Println(ages["mohammad"]) //	22
	ages["mohammad"]++
	fmt.Println(ages["mohammad"]) //	23
	ages["abcdef"]++              // define new key var!!!
	fmt.Println(ages)             //	map[abcdef:1 mohammad:23] -> define new key var!
	//address := &ages["mohammad"]  //	cannot take address of ages["mohammad"] (map index expression of type int)

	// not same in slice
	var arr []int
	arr = append(arr, 1)
	fmt.Println(arr)     // [1]
	address := &arr[0]   //
	fmt.Println(address) //	0xc000014158

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
		//mohammad	23
		//abcdef	1
	}

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
		//	abcdef	1
		//	mohammad	23
	}

	var m map[string]int
	fmt.Println(m == nil)    // "true"
	fmt.Println(len(m) == 0) // "true"

	//m["ali"] = 20	//	panic: assignment to entry in nil map

	_, ok := ages["bob"]
	if !ok {
		fmt.Println("key is not exist")
	}
	if _, ok := ages["bob"]; !ok {
		fmt.Println("key is not exist")
	}

}
