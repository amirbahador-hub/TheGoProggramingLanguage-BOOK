package main

import "fmt"

type tree struct {
	value int
	left  *tree
	right *tree
}

func add(t *tree, v int) *tree {
	if t == nil {
		t = new(tree)
		t.value = v
		return t
	}
	if v < t.value {
		t.left = add(t.left, v)
	} else {
		t.right = add(t.right, v)
	}
	return t
}

func appendValues(v []int, t *tree) []int {
	if t != nil {
		v = appendValues(v, t.left)
		v = append(v, t.value)
		v = appendValues(v, t.right)
	}
	return v
}

func sort(values []int) []int {
	var t *tree
	for _, v := range values {
		t = add(t, v)
	}
	values = appendValues(values[:0], t)
	return values
}

func main() {
	values := []int{2, 1, 3, 5, 4, 9, 7, 8}
	fmt.Println(sort(values))
}
