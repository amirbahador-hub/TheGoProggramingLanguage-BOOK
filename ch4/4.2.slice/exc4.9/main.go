package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordNumber := make(map[string]int)
	if len(os.Args) != 2 {
		fmt.Println("please enter file name after go run main!")
	}
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Errorf("error : %f", err)
		os.Exit(1)
	}
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		wordNumber[input.Text()]++
	}
	
	fmt.Println(wordNumber)
}
