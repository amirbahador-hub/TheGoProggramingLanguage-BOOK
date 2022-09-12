package binaryseasrch

func BinarySearch(left int, right int, x int, array []int) int {
	if left > right {
		return 0
	} else {
		mid := (left + right) / 2
		if array[mid] == x {
			return mid
		} else if array[mid] > x {
			return BinarySearch(mid+1, right, x, array)
		} else {
			return BinarySearch(left, mid-1, x, array)
		}
	}
}
