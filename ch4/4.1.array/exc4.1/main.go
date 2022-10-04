package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x byte) int {
	return int(pc[x>>(0*8)] +
		pc[x>>(1*8)] +
		pc[x>>(2*8)] +
		pc[x>>(3*8)] +
		pc[x>>(4*8)] +
		pc[x>>(5*8)] +
		pc[x>>(6*8)] +
		pc[x>>(7*8)])
}

func CountDiff(sha1, sha2 [32]uint8) int {
	var diff int
	for i := 0; i < 32; i++ {
		diff += PopCount(sha1[i] ^ sha2[i])
	}
	return diff
}

func main() {
	c1 := sha256.Sum256([]byte("abc"))
	c2 := sha256.Sum256([]byte("ABC"))
	fmt.Println(CountDiff(c1, c2))
}
