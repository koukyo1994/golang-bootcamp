package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
	"time"
)

type PixelValue struct {
	X, Y  int
	Color color.RGBA
}

func main() {
	start := time.Now()

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	pixelChans := make(chan PixelValue, width*height)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			go func(px, py int, z complex128) {
				const iterations = 360
				var v complex128
				c := color.RGBA{0, 0, 0, 255}

				for n := 0; n < iterations; n++ {
					v = v*v + z
					if cmplx.Abs(v) > 2 {
						c = hsv2rgb(hsv(z, n))
						break
					}
				}
				pixelChans <- PixelValue{px, py, c}
			}(px, py, z)
		}
	}
	for i := 0; i < width*height; i++ {
		pv := <-pixelChans
		img.Set(pv.X, pv.Y, pv.Color)
	}
	f, err := os.OpenFile("/dev/null", os.O_RDWR, os.ModeAppend)
	if err != nil {
		log.Fatal("Could not open /dev/null")
	}
	png.Encode(f, img) // NOTE: ignoring errors
	elapsed := time.Since(start).Seconds()
	fmt.Printf("%.5f seconds elapsed\n", elapsed)
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
