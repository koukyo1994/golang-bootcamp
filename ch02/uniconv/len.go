package uniconv

import "fmt"

type Meter float64
type Feet float64

const (
	MeterPerFoot = 3.28084
	FootPerMeter = 1 / MeterPerFoot
)

func (m Meter) String() string {
	return fmt.Sprintf("%g meter", m)
}

func (f Feet) String() string {
	return fmt.Sprintf("%g feet", f)
}

func (m Meter) ToFeet() Feet {
	return Feet(m * FootPerMeter)
}

func (f Feet) ToMeter() Meter {
	return Meter(f * MeterPerFoot)
}
