package main

import (
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		temp, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parse : %d \n")
			os.Exit(1)
		}
		fmt.Println("celicius temp is : ", tempconv.Kelvin(temp).KToC())
		//repeat this for another purpose
	}
}
