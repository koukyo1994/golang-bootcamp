package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	MAXITER  = 1000
	TOL      = 1e-10
	TOLANS   = 1e-5
	CONTRAST = 15
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			ans, cnt := newton(z)
			img.Set(px, py, decideColor(ans, cnt))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func f(z complex128) complex128 {
	return z*z*z*z - 1
}

func fprime(z complex128) complex128 {
	return 4 * z * z * z
}

func newton(z complex128) (complex128, int) {
	var cnt int
	for {
		z = z - f(z)/fprime(z)
		cnt++
		if cmplx.Abs(f(z)) < TOL {
			break
		} else if cnt > MAXITER {
			break
		}
	}
	return z, cnt
}

func decideColor(z complex128, cnt int) color.RGBA {
	cnt = cnt * CONTRAST
	if cnt > 255 {
		cnt = 255
	}
	max := float64(cnt)
	if cmplx.Abs(1-z) < TOLANS {
		return color.RGBA{uint8(max), 0, 0, 255}
	} else if cmplx.Abs(-1-z) < TOLANS {
		return color.RGBA{uint8(0.5 * max), uint8(max), 0, 255}
	} else if cmplx.Abs(-1i-z) < TOLANS {
		return color.RGBA{0, uint8(max), uint8(max), 255}
	} else if cmplx.Abs(1i-z) < TOLANS {
		return color.RGBA{uint8(0.5 * max), 0, uint8(max), 255}
	} else {
		return color.RGBA{0, 0, 0, 255}
	}
}
