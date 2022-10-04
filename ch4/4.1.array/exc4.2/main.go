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
		fmt.Printf("%x \n", sha512.Sum384([]byte(*value)))
	case "sha512":
		fmt.Printf("%x \n", sha512.Sum512([]byte(*value)))
	default:
		fmt.Printf("%x \n", sha256.Sum256([]byte(*value)))
	}
}

// run : go run main.go -t sha384 -v abc
