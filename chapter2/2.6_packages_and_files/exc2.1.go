package tempconv

func (c Celsius) CToK() Kelvin {
	return Kelvin(c + 273.15)
}

func (k Kelvin) KToC() Celsius {
	return Celsius(k - 273.15)
}
