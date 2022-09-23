package main

import "fmt"

func nonempty(arr []string) []string {
	i := 0
	for _, value := range arr {
		if value != "" {
			arr[i] = value
			i++
		}
	}
	return arr[:i]
}

func nonempty2(arr []string) []string {
	out := arr[:0] // zero-length slice of original
	for _, s := range arr {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func nonempty3(arr []string) []string {
	var out []string
	copy(out, arr)
	i := 0
	for _, value := range out {
		if value != "" {
			out[i] = value
			i++
		}
	}
	return out[:i]
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	data := []string{"a", "", "b"}
	fmt.Println(data)           //	[a   b]
	fmt.Println(nonempty(data)) // [a b]
	fmt.Println(data)           //	[a b b]

	data2 := []string{"a", "", "b"}
	fmt.Println(data2)            //	[a   b]
	fmt.Println(nonempty2(data2)) // [a b]
	fmt.Println(data2)            //	[a b b]

	data3 := []string{"a", "", "b"}
	fmt.Println(data3)            //	[a  b]
	fmt.Println(nonempty3(data3)) // [a b]
	fmt.Println(data3)            //	[a  b]

	//stack
	stack := []int{1, 2}
	stack = append(stack, 3) // push v
	fmt.Println(stack)       //	[1 2 3]

	top := stack[len(stack)-1] // top of stack
	fmt.Println(top)           //	3

	stack = stack[:len(stack)-1] // pop
	fmt.Println(stack)           //	[1 2]

	s := []int{1, 2, 3, 5}
	fmt.Println(remove(s, 3))  // [1 2 3]
	fmt.Println(remove2(s, 1)) // [1 5 3]
}
