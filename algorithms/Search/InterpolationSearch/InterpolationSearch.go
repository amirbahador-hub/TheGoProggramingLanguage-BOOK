package main

import (
	"fmt"
	"time"
)

/*
explain about formula of Interpolation search
this method use specific distribute value in sorted array like linear distribute

(1): linear distribute is y = m * x + n,
(2): and slope is m = (y2 - y1) / (x2 - x1)
	 so we have low and high index like binary search, but we have optimized change rage of low and high by distributed values
(3): arr[low] = m * low + n
(4): arr[high] = m * high + n
(5): x = m * pos + n -> final values
(6): m = (arr[high] - arr[low]) / (high - low)
	 we subtract 3 from 5 for n
(7): x = m * pos + arr[low] - m * low
(8): x - arr[low] = m * (pos - low)
(9): pos = low + (x - arr[low]) / m
	 finally, we subtract 6 from 9
	 pos = low + (x - arr[low]) * (high - low) / (arr[high] - arr[low])
*/

func interpolationSearch(arr []int, find int, defaultIndex ...int) int {
	var defaultIndexValue int   //	this variable store defaultIndex from argument
	if len(defaultIndex) == 0 { //	check if defaultIndex parameter does not exist set defaultIndexValue is 0
		defaultIndexValue = 0
	} else {
		defaultIndexValue = defaultIndex[0] //	check if defaultIndex is existed, set defaultIndexValue is first index of parameter
	}

	if len(arr) == 0 { // if value does not exist in array and smaller that highest array value, we should check this
		return -1
	}

	pos := (find - arr[0]) * (len(arr) - 1) / (arr[len(arr)-1] - arr[0])

	if pos >= len(arr) {
		return -1 // if value does not exist in array and bigger that highest array value, we should check this
	}

	if arr[pos] == find {
		/*	if find value, return index of value (pos) plus defaultIndexValue,
			because defaultIndexValue illustrate last min var index from last recursive func	*/
		return pos + defaultIndexValue
	} else if find < arr[pos] {
		return interpolationSearch(arr[:pos], find, defaultIndexValue) // search on right side
	} else {
		defaultIndexValue += pos
		return interpolationSearch(arr[pos+1:], find, defaultIndexValue+1) //	search on left side, consider defaultIndexValue for preserve index min from last recursive func
	}
}

func main() {
	arr := []int{1, 3, 5, 8, 11, 13, 14, 19, 27}

	totalStart := time.Now()

	start := time.Now()
	fmt.Println(interpolationSearch(arr, 1))
	elapsed := time.Since(start)
	fmt.Println("elapsed time for 1", elapsed)

	start = time.Now()
	fmt.Println(interpolationSearch(arr, 5))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 5", elapsed)

	start = time.Now()
	fmt.Println(interpolationSearch(arr, 8))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 8", elapsed)

	start = time.Now()
	fmt.Println(interpolationSearch(arr, 11))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 11", elapsed)

	start = time.Now()
	fmt.Println(interpolationSearch(arr, 13))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 13", elapsed)

	start = time.Now()
	fmt.Println(interpolationSearch(arr, 19))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 19", elapsed)

	start = time.Now()
	fmt.Println(interpolationSearch(arr, 27))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 27", elapsed)

	start = time.Now()
	fmt.Println(interpolationSearch(arr, 12))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 12", elapsed)

	start = time.Now()
	fmt.Println(interpolationSearch(arr, 200))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 200", elapsed)

	elapsedTotal := time.Since(totalStart)

	fmt.Println("total time elapsed: ", elapsedTotal) //	90.3374µs
}
