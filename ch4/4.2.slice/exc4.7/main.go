package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(bytes []byte) []byte {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return bytes
}

func reverseRune(bytes []byte) []byte {
	for i := 0; i < len(bytes); {
		_, size := utf8.DecodeRune(bytes[i:])
		reverse(bytes[i : i+size])
		i += size
	}
	reverse(bytes)
	return bytes
}

func main() {
	fmt.Println(reverseRune([]byte("ab cسسس")))
}
