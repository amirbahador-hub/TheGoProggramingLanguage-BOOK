package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type categories map[string]int

func (c categories) counterCategories(r rune) {
	if unicode.IsControl(r) {
		c["Control"]++
	}
	if unicode.IsDigit(r) {
		c["Digit"]++
	}
	if unicode.IsGraphic(r) {
		c["Graphic"]++
	}
	if unicode.IsLetter(r) {
		c["Letter"]++
	}
	if unicode.IsLower(r) {
		c["Lower"]++
	}
	if unicode.IsMark(r) {
		c["Mark"]++
	}
	if unicode.IsNumber(r) {
		c["Number"]++
	}
	if unicode.IsPrint(r) {
		c["Printable"]++
	} else {
		c["NonPrintable"]++
	}

	if unicode.IsPunct(r) {
		c["Punct"]++
	}
	if unicode.IsSpace(r) {
		c["Space"]++
	}
	if unicode.IsSymbol(r) {
		c["Symbol"]++
	}
	if unicode.IsTitle(r) {
		c["Title"]++
	}
	if unicode.IsUpper(r) {
		c["Upper"]++
	}
}

func main() {
	count := make(map[rune]int)        //	count of unicode character
	var utflength [utf8.UTFMax + 1]int // count of length of utf8 encoding
	invalid := 0
	c := make(categories)

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount %v \n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		count[r]++
		utflength[n]++
		c.counterCategories(r)
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range count {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflength {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ncategory\tcount\n")
	for category, count := range c {
		fmt.Printf("%s\t\t%d\n", category, count)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
