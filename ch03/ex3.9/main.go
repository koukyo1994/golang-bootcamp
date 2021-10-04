package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		x, y       float64 = 0.0, 0.0
		resolution int     = 1024
	)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		switch k {
		case "x":
			x, _ = strconv.ParseFloat(v[0], 64)
		case "y":
			y, _ = strconv.ParseFloat(v[0], 64)
		case "resolution":
			resolution, _ = strconv.Atoi(v[0])
		}
	}
	w.Header().Set("Content-Type", "image/png")
	drawImage(w, x, y, resolution)
}

func drawImage(out io.Writer, x, y float64, resolution int) {
	var width, height = resolution, resolution
	var xmin, ymin, xmax, ymax = x - 2, y - 2, x + 2, y + 2

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
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
