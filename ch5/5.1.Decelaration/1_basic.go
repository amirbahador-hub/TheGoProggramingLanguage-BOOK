package main

import "fmt"

func f1(i, j, k int, s, t string) {
	fmt.Println(i, j, k)
	fmt.Println(s, t)
}
func f2(i int, j int, k int, s string, t string) {
	fmt.Println(i, j, k)
	fmt.Println(s, t)
}

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return
}

func first(x int, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}

func main() {
	f1(1, 2, 3, "hello", "how")
	//1 2 3
	//hello how
	f2(1, 2, 3, "hello", "how")
	//1 2 3
	//hello how
	//both of them same

}
