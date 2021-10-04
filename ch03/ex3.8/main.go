package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
	"math/cmplx"
	"os"
	"runtime"
	"strconv"
)

func mandelbrot128(z complex128) color.RGBA {
	const iterations = 360

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return hsv2rgb(hsv(z, n))
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func mandelbrot64(z complex64) color.RGBA {
	const iterations = 360

	var v complex64
	for n := 0; n < iterations; n++ {
		v = v*v + z
		// complex128でないと動作しないためキャストする
		// 解像度には影響しない
		if cmplx.Abs(complex128(v)) > 2 {
			return hsv2rgb(hsv(complex128(z), n))
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func mandelbrotBigFloat(real *big.Float, imag *big.Float) color.RGBA {
	const iterations = 360

	var vreal = big.NewFloat(0)
	var vimag = big.NewFloat(0)
	for n := 0; n < iterations; n++ {
		vreal, vimag = cmplxSquareBigFloat(vreal, vimag)
		vreal.Add(vreal, real)
		vimag.Add(vimag, imag)
		if cmplxAbsBigFloat(vreal, vimag).Cmp(big.NewFloat(2)) == 1 {
			realFloat64, _ := real.Float64()
			imagFloat64, _ := imag.Float64()
			return hsv2rgb(hsv(complex(realFloat64, imagFloat64), n))
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func cmplxSquareBigFloat(real *big.Float, imag *big.Float) (*big.Float, *big.Float) {
	newReal := new(big.Float).Sub(new(big.Float).Mul(real, real), new(big.Float).Mul(imag, imag))
	newImag := new(big.Float).Add(new(big.Float).Mul(real, imag), new(big.Float).Mul(real, imag))
	return newReal, newImag
}

func cmplxAbsBigFloat(real *big.Float, imag *big.Float) *big.Float {
	return new(big.Float).Sqrt(new(big.Float).Add(new(big.Float).Mul(real, real), new(big.Float).Mul(imag, imag)))
}

func mandelbrotBigRat(real *big.Rat, imag *big.Rat) color.RGBA {
	const iterations = 60

	var vreal = big.NewRat(0, 1)
	var vimag = big.NewRat(0, 1)
	for n := 0; n < iterations; n++ {
		vreal, vimag = cmplxSquareBigRat(vreal, vimag)
		vreal.Add(vreal, real)
		vimag.Add(vimag, imag)
		if cmplxSumSquareBigRat(vreal, vimag).Cmp(big.NewRat(4, 1)) == 1 {
			realFloat64, _ := real.Float64()
			imagFloat64, _ := imag.Float64()
			return hsv2rgb(hsv(complex(realFloat64, imagFloat64), n))
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func cmplxSquareBigRat(real *big.Rat, imag *big.Rat) (*big.Rat, *big.Rat) {
	newReal := new(big.Rat).Sub(new(big.Rat).Mul(real, real), new(big.Rat).Mul(imag, imag))
	newImag := new(big.Rat).Add(new(big.Rat).Mul(real, imag), new(big.Rat).Mul(real, imag))
	return newReal, newImag
}

func cmplxSumSquareBigRat(real *big.Rat, imag *big.Rat) *big.Rat {
	return new(big.Rat).Add(new(big.Rat).Mul(real, real), new(big.Rat).Mul(imag, imag))
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)

	var width, height int

	var resolutions = []int{256, 512, 1024}
	var precisions = []string{"complex64", "complex128", "big.Float", "big.Rat"}

	var mem runtime.MemStats

	for _, precision := range precisions {
		for _, resolution := range resolutions {
			fmt.Printf("Calculating %s, %d...\n", precision, resolution)

			runtime.ReadMemStats(&mem)
			fmt.Printf("Memory: %d\n", mem.Alloc)

			f, err := os.Create("assets/" + precision + "_" + strconv.Itoa(resolution) + ".png")
			if err != nil {
				panic(err)
			}
			defer f.Close()

			width = resolution
			height = resolution

			img := image.NewRGBA(image.Rect(0, 0, width, height))
			for py := 0; py < height; py++ {
				y := float64(py)/float64(height)*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/float64(width)*(xmax-xmin) + xmin
					switch precision {
					case "complex64":
						z := complex(float32(x), float32(y))
						img.Set(px, py, mandelbrot64(complex64(z)))
					case "complex128":
						z := complex(x, y)
						img.Set(px, py, mandelbrot128(z))
					case "big.Float":
						real := big.NewFloat(x)
						imag := big.NewFloat(y)
						img.Set(px, py, mandelbrotBigFloat(real, imag))
					case "big.Rat":
						real := big.NewRat(int64(x*1e6), 1e6)
						imag := big.NewRat(int64(y*1e6), 1e6)
						img.Set(px, py, mandelbrotBigRat(real, imag))
					}
				}
			}
			runtime.ReadMemStats(&mem)
			fmt.Printf("Memory: %d\n", mem.Alloc)
			png.Encode(f, img)

			runtime.GC()
		}
	}
}

func hsv(z complex128, n int) (h, s, v int) {
	angle := math.Atan2(imag(z), real(z))
	abs := cmplx.Abs(z)

	h = n
	s = int((abs - 2) * 100)
	if s > 255 {
		s = 255
	}
	v = int(math.Abs(angle*180/math.Pi) / 360 * 255)
	return h, s, v
}

func hsv2rgb(h, s, v int) color.RGBA {
	max := float64(v)
	min := max - (float64(s) / 255 * max)

	var r, g, b float64
	switch {
	case h < 60:
		r = max
		g = (float64(h)/60)*(max-min) + min
		b = min
	case h < 120:
		r = (120-float64(h))/60*(max-min) + min
		g = max
		b = min
	case h < 180:
		r = min
		g = max
		b = ((float64(h)-120)/60)*(max-min) + min
	case h < 240:
		r = min
		g = (240-float64(h))/60*(max-min) + min
		b = max
	case h < 300:
		r = ((float64(h)-240)/60)*(max-min) + min
		g = min
		b = max
	case h < 360:
		r = max
		g = min
		b = (360-float64(h))/60*(max-min) + min
	default:
		r = 0
		g = 0
		b = 0
	}

	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}
