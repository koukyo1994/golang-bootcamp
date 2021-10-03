package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	samplingHeight := (ymax - ymin) / float64(height) * 0.5
	samplingWidth := (xmax - xmin) / float64(width) * 0.5
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin

		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			z0 := complex(x, y)
			z1 := complex(x+samplingWidth, y)
			z2 := complex(x, y+samplingHeight)
			z3 := complex(x+samplingWidth, y+samplingHeight)

			c0 := mandelbrot(z0)
			c1 := mandelbrot(z1)
			c2 := mandelbrot(z2)
			c3 := mandelbrot(z3)

			r := (float64(c0.R) + float64(c1.R) + float64(c2.R) + float64(c3.R)) / 4
			g := (float64(c0.G) + float64(c1.G) + float64(c2.G) + float64(c3.G)) / 4
			b := (float64(c0.B) + float64(c1.B) + float64(c2.B) + float64(c3.B)) / 4

			c := color.RGBA{uint8(r), uint8(g), uint8(b), 255}
			img.Set(px, py, c)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
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

func mandelbrot(z complex128) color.RGBA {
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
