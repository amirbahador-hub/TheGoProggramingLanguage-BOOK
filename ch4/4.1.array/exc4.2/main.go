package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	shaType := flag.String("t", "sha256", "sha type")
	value := flag.String("v", "value", "value")
	flag.Parse()
	switch *shaType {
	case "sha384":
		fmt.Println(sha512.Sum384([]byte(*value)))
	case "sha512":
		fmt.Println(sha512.Sum512([]byte(*value)))
	default:
		fmt.Println(sha256.Sum256([]byte(*value)))
	}
}

// run : go run main.go -t sha384 -v abc
