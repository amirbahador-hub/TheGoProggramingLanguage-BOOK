package tempconv

import "fmt"

type Kelvin float64
type Celsius float64

const (
	KelvinTemp  Kelvin  = 120
	CelsiusTemp Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%0.2f °C", c)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%0.2f °k", k)
}
