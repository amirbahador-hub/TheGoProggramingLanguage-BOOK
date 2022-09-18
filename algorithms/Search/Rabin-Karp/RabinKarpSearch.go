package main

import "fmt"

type Pattern string
type Text string

type RabinCarp struct {
	d     int //	possible character number
	prime int //	a prime number
	n     int //	number text length
	m     int //	number pattern length
	h     int //	d ^ (m - 1) but this never bigger than prime value
	tHash int //	hash value for text, but this never bigger than prime value
	pHash int //	hash value for pattern, but this never bigger than prime value
}

func buildRabinCarpParameter(text Text, pattern Pattern) *RabinCarp {
	d, prime, n, m, h := 256, 101, len(text), len(pattern), 1
	var tHash, pHash int
	for i := 0; i < m-1; i++ {
		h = (h * d) % prime
	}

	for i := 0; i < m; i++ {
		tHash = (h*tHash + int(text[i])) % prime
		pHash = (h*pHash + int(pattern[i])) % prime
	}

	return &RabinCarp{d: d, prime: prime, n: n, m: m, h: h, tHash: tHash, pHash: pHash}
}

func (rabinCarpParameter *RabinCarp) updateHashValue(index int, text Text) {
	//	hash( txt[s+1 .. s+m] ) = ( d ( hash( txt[s .. s+m-1]) – txt[s]*h ) + txt[s + m] ) mod q
	rabinCarpParameter.tHash = (rabinCarpParameter.d*(rabinCarpParameter.tHash-int(text[index])*rabinCarpParameter.h) + int(text[index+rabinCarpParameter.m])) % rabinCarpParameter.prime
	if rabinCarpParameter.tHash < 0 {
		rabinCarpParameter.tHash = rabinCarpParameter.tHash + rabinCarpParameter.prime
	}
}

func (pattern Pattern) rabinKarpMatchPattern(text Text) <-chan int {
	rabinCarpParameter := *buildRabinCarpParameter(text, pattern)

	channel := make(chan int)

	var j int
	go func() {
		for {
			for i := 0; i <= rabinCarpParameter.n-rabinCarpParameter.m; i++ {
				if rabinCarpParameter.tHash == rabinCarpParameter.pHash {
					for j = 0; j < rabinCarpParameter.m; j++ {
						if pattern[j] != text[i+j] {
							break
						}
					}
					if j == rabinCarpParameter.m {
						channel <- i
					}
				}

				if i < rabinCarpParameter.n-rabinCarpParameter.m {
					rabinCarpParameter.updateHashValue(i, text)
				}
			}
			close(channel)
			return
		}
	}()
	return channel
}

func main() {
	var text Text = "hellohehe"
	var pattern Pattern = "he"
	for r := range pattern.rabinKarpMatchPattern(text) {
		fmt.Println("index if pattern is :", r)
	}
}
