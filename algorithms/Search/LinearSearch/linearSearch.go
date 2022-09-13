package main

import (
	"fmt"
	"time"
)

func linearSearch(arr []int, find int) int {
	for key, value := range arr {
		if value == find {
			return key
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 8, 11, 13, 14, 19, 27}

	totlalStart := time.Now()

	start := time.Now()
	fmt.Println(linearSearch(arr, 1))
	elapsed := time.Since(start)
	fmt.Println("elapsed time for 1", elapsed)

	start = time.Now()
	fmt.Println(linearSearch(arr, 5))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 5", elapsed)

	start = time.Now()
	fmt.Println(linearSearch(arr, 8))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 8", elapsed)

	start = time.Now()
	fmt.Println(linearSearch(arr, 11))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 11", elapsed)

	start = time.Now()
	fmt.Println(linearSearch(arr, 13))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 13", elapsed)

	start = time.Now()
	fmt.Println(linearSearch(arr, 19))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 19", elapsed)

	start = time.Now()
	fmt.Println(linearSearch(arr, 27))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 27", elapsed)

	start = time.Now()
	fmt.Println(linearSearch(arr, 12))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 12", elapsed)

	elapsedTotal := time.Since(totlalStart)

	fmt.Println("total time elapsed: ", elapsedTotal) //	92.537µs

}
