package main

import "fmt"

func main() {

	const (
		AA int = 126
		BB
		CC
		DD int = 135
		EE
	)

	fmt.Printf(" %d \n %d \n %d \n %d \n %d \n", AA, BB, CC, DD, EE) // print all 126 except DD and EE print 135

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "a", GBP: "b", RMB: "c"}
	fmt.Println(symbol, symbol[USD]) // return [$ a b c] $

}
