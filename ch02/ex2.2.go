package main

import (
	"bufio"
	"ch02/uniconv"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var value float64
	if len(os.Args) != 2 {
		sc := bufio.NewScanner(os.Stdin)
		if sc.Scan() {
			value, _ = strconv.ParseFloat(sc.Text(), 64)
		}
	} else {
		value, _ = strconv.ParseFloat(os.Args[1], 64)
	}

	c := uniconv.Celsius(value)
	f := uniconv.Fahrenheit(value)

	meter := uniconv.Meter(value)
	feet := uniconv.Feet(value)

	p := uniconv.Pound(value)
	k := uniconv.Kilogram(value)

	fmt.Printf("%s = %s, %s = %s\n", c, c.ToFahrenheit(), f, f.ToCelsius())
	fmt.Printf("%s = %s, %s = %s\n", meter, meter.ToFeet(), feet, feet.ToMeter())
	fmt.Printf("%s = %s, %s = %s\n", p, p.ToKilogram(), k, k.ToPound())
}
