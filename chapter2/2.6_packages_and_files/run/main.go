package main

import (
	"fmt"
	"tempconv"
)

func main() {
	fmt.Println(tempconv.KelvinTemp.CToK().String())
	fmt.Println(tempconv.CelsiusTemp.KToC().String())
}
