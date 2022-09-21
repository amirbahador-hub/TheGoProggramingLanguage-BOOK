package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

func main() {
	var runes []rune
	fmt.Printf("type is %T , len is : %d, cap is : %d \n", runes, len(runes), cap(runes)) //	type is []int32 , len is : 0, cap is : 0

	for _, value := range "hello سلام" {
		runes = append(runes, value)
	}

	fmt.Println(runes) //	[104 101 108 108 111 32 1587 1604 1575 1605]

	fmt.Printf("value is %q , len is : %d, cap is : %d \n", runes, len(runes), cap(runes)) //	value is ['h' 'e' 'l' 'l' 'o' ' ' 'س' 'ل' 'ا' 'م'] , len is : 10, cap is : 16

	convenientlyRunes := []rune("hello سلام")

	fmt.Println(convenientlyRunes) //	[104 101 108 108 111 32 1587 1604 1575 1605]

	fmt.Printf("value is %q , len is : %d, cap is : %d \n", runes, len(runes), cap(runes)) //	value is ['h' 'e' 'l' 'l' 'o' ' ' 'س' 'ل' 'ا' 'م'] , len is : 10, cap is : 16

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	//0 cap=1	[0]
	//1 cap=2	[0 1]
	//2 cap=4	[0 1 2]
	//3 cap=4	[0 1 2 3]
	//4 cap=8	[0 1 2 3 4]
	//5 cap=8	[0 1 2 3 4 5]
	//6 cap=8	[0 1 2 3 4 5 6]
	//7 cap=8	[0 1 2 3 4 5 6 7]
	//8 cap=16	[0 1 2 3 4 5 6 7 8]
	//9 cap=16	[0 1 2 3 4 5 6 7 8 9]

}
