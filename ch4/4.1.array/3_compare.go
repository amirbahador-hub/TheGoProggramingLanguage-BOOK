package main

import (
	"fmt"
)
import "crypto/sha256"

func main() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	d := [...]int{1, 3}

	fmt.Println(a == b, a == c, a == d, b == c, b == d, c == d) //	true false false false false true

	//e := [3]int{1, 2}
	//fmt.Println(a == e) // compile error: mismatched types [2]int and [3]int

	bytes := []byte("abcd")
	fmt.Println(bytes)         //	[97 98 99 100]
	fmt.Printf("%x \n", bytes) //	61626364 -> 61 62 63 64
	fmt.Printf("%T \n", bytes) // uint8

	cr1 := sha256.Sum256([]byte("x")) // print long array [12, 22 , ... ]
	cr2 := sha256.Sum256([]byte("X")) // print long array [12, 22 , ... ]
	fmt.Println(cr1)
	fmt.Println(cr2)

	fmt.Printf("%x\n%x\n%t\n%T\n", cr1, cr2, cr1 == cr2, cr1)
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	//false
	//[32]uint8
}
