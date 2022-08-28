package tempconv

import "fmt"

func (c Celsius) KToC() Kelvin {
	return Kelvin(c + 273.15)
}

func (k Kelvin) CToK() Celsius {
	return Celsius(k - 273.15)
}

func main() {
	fmt.Println(KelvinTemp.CToK().String())
	fmt.Println(CelsiusTemp.KToC().String())
}
