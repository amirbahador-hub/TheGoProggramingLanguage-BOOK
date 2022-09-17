package main

import (
	"fmt"
	"math"
	"time"
)

func initFibonacci(len int) (int, int, int) {
	fk0, fk1, fk2 := 0, 1, 1
	for fk2 <= len {
		fk0 = fk1
		fk1 = fk2
		fk2 = fk0 + fk1
	}
	return fk0, fk1, fk2
}

func shiftFibonacci(fk0 *int, fk1 *int, fk2 *int, step int) {
	for i := step; i != 0; i-- {
		*fk2 = *fk1
		*fk1 = *fk0
		*fk0 = *fk2 - *fk1
	}
}

func FibonacciSearch(arr []int, find int) int {
	var index, offset = 0, -1                // init index and offset variables
	fk0, fk1, fk2 := initFibonacci(len(arr)) // init fibonacci variables

	for {
		index = int(math.Min(float64(offset+fk0), float64(len(arr)-1))) // find min(offset + fk0, len(array) -1)
		if arr[index] < find {
			shiftFibonacci(&fk0, &fk1, &fk2, 1) //	if value smaller than find, one step fibonacci index turn left and offset shift right to index
			offset = index
		} else if arr[index] > find {
			shiftFibonacci(&fk0, &fk1, &fk2, 2) //	if value bigger than find, two step fibonacci index turn left
		} else {
			return index // if arr[index] == find, return index
		}
		if fk0 == 0 {
			break // this loop repeat until f0 set to 0
		}
	}

	return -1 // if can't find fibonacci, return -1
}

func main() {
	arr := []int{1, 3, 5, 8, 11, 13, 14, 19, 27}

	totalStart := time.Now()

	start := time.Now()
	fmt.Println(FibonacciSearch(arr, 1))
	elapsed := time.Since(start)
	fmt.Println("elapsed time for 1", elapsed)

	start = time.Now()
	fmt.Println(FibonacciSearch(arr, 5))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 5", elapsed)

	start = time.Now()
	fmt.Println(FibonacciSearch(arr, 8))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 8", elapsed)

	start = time.Now()
	fmt.Println(FibonacciSearch(arr, 11))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 11", elapsed)

	start = time.Now()
	fmt.Println(FibonacciSearch(arr, 13))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 13", elapsed)

	start = time.Now()
	fmt.Println(FibonacciSearch(arr, 19))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 19", elapsed)

	start = time.Now()
	fmt.Println(FibonacciSearch(arr, 27))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 27", elapsed)

	start = time.Now()
	fmt.Println(FibonacciSearch(arr, 12))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 12", elapsed)

	start = time.Now()
	fmt.Println(FibonacciSearch(arr, 200))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 200", elapsed)

	elapsedTotal := time.Since(totalStart)

	fmt.Println("total time elapsed: ", elapsedTotal) //	130.048µs
}
