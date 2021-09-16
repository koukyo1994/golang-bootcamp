package uniconv

import "fmt"

type Pound float64
type Kilogram float64

const (
	PoundToKilogram = 0.45359237
	KilogramToPound = 2.20462262
)

func (p Pound) ToKilogram() Kilogram {
	return Kilogram(p * PoundToKilogram)
}

func (k Kilogram) ToPound() Pound {
	return Pound(k * KilogramToPound)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g pound", p)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g kg", k)
}
