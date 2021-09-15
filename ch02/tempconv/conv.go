package tempconv

// CToFは摂氏の温度を華氏へ変換します。
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToCは華氏の温度を摂氏へ変換します。
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToCは絶対温度を摂氏へ変換します。
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// CToKは摂氏を絶対温度へ変換します。
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KToFは絶対温度を華氏へ変換します。
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15)*9/5) + 32 }

// FToKは華氏を絶対温度へ変換します。
func FToK(f Fahrenheit) Kelvin { return Kelvin((f-32)*5/9 + 273.15) }
