package tempconv

import (
	tc "bootcamp/ch02/tempconv"
	"flag"
	"fmt"
)

// *celsiusFlagはflag.Valueインターフェースを満足する
// String()は既に定義されている
type celsiusFlag struct{ tc.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // エラーは検査しない
	switch unit {
	case "C", "°C":
		f.Celsius = tc.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tc.FToC(tc.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value tc.Celsius, usage string) *tc.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
