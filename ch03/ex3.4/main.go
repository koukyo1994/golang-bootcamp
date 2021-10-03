package main

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

const (
	MAX = 255
	MIN = 0
)

func corner(i, j int, cells float64, xyrange float64, width float64, height float64, xyscale float64, angle float64, zscale float64) (float64, float64, float64, float64) {
	// マス目 (i, j)の角の点(x, y)を見つける。
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する。
	z := f(x, y)
	r := math.Hypot(x, y)

	cos := math.Cos(angle)
	sin := math.Sin(angle)

	// (x, y, z)を2-D SVGキャンバス(sx, sy)へ等角的に投影。
	sx := width/2 + (x-y)*cos*xyscale
	sy := height/2 + (x+y)*sin*xyscale - z*zscale
	return sx, sy, z, r
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // (x, y)の距離
	return math.Sin(r) / r
}

func decideColor(z, r float64) color.RGBA {
	// zを0 - 255の範囲に変換する。
	ztransformed := -(z*r - 1) / 2 * MAX
	switch {
	case ztransformed <= 60:
		return color.RGBA{MAX, uint8(ztransformed/60*(MAX-MIN) + MIN), MIN, 255}
	case 60 < ztransformed && ztransformed <= 120:
		return color.RGBA{uint8((120-ztransformed)/60*(MAX-MIN) + MIN), MAX, MIN, 255}
	case 120 < ztransformed && ztransformed <= 180:
		return color.RGBA{MIN, MAX, uint8((ztransformed-120)/60*(MAX-MIN) + MIN), 255}
	case 180 < ztransformed && ztransformed <= 240:
		return color.RGBA{MIN, uint8((240-ztransformed)/60*(MAX-MIN) + MIN), MAX, 255}
	case 240 < ztransformed:
		return color.RGBA{uint8((ztransformed-240)/60*(MAX-MIN) + MIN), MIN, MAX, 255}
	}
	// ここには来ないはず。
	return color.RGBA{MIN, MIN, MIN, 255}
}

func colorToString(c color.RGBA) string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c.R, c.G, c.B)
}

func isin(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

func drawSurface(out io.Writer, cells int, xyrange float64, width float64, height float64, xyscale float64, angle float64, zscale float64, colorStr string) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke-width:0.7' "+
		"width='%d' height='%d'>", int(width), int(height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, ar := corner(i+1, j, float64(cells), xyrange, width, height, xyscale, angle, zscale)
			bx, by, bz, br := corner(i, j, float64(cells), xyrange, width, height, xyscale, angle, zscale)
			cx, cy, cz, cr := corner(i, j+1, float64(cells), xyrange, width, height, xyscale, angle, zscale)
			dx, dy, dz, dr := corner(i+1, j+1, float64(cells), xyrange, width, height, xyscale, angle, zscale)

			// NaNを除外するために、それぞれの点をチェックし、見つけた場合はcontinueする。
			hasNaN := false
			for _, val := range []float64{ax, ay, bx, by, cx, cy, dx, dy} {
				if math.IsNaN(val) {
					hasNaN = true
					break
				}
			}
			if hasNaN {
				continue
			}

			// 色を決定する。
			var fillColor string
			if colorStr == "rainbow" {
				fillColor = colorToString(decideColor((az+bz+cz+dz)/4, (ar+br+cr+dr)/4))
			} else if strings.HasPrefix(colorStr, "#") {
				c := color.RGBA{255, 255, 255, 255}
				switch len(colorStr) {
				case 7:
					_, err := fmt.Sscanf(colorStr, "#%02x%02x%02x", &c.R, &c.G, &c.B)
					if err != nil {
						c = color.RGBA{255, 255, 255, 255}
					}
				case 4:
					_, err := fmt.Sscanf(colorStr, "#%01x%01x%01x", &c.R, &c.G, &c.B)
					if err != nil {
						c = color.RGBA{255, 255, 255, 255}
					}
					c.R *= 17
					c.G *= 17
					c.B *= 17
				}
				fillColor = colorToString(c)
			} else if isin(colorStr, []string{"black", "white", "red", "green", "blue", "yellow", "orange"}) {
				fillColor = colorStr
			} else {
				fillColor = "white"
			}

			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, fillColor)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		width, height int     = 600, 320
		cells         int     = 100
		xyrange       float64 = 30
		angle         float64 = math.Pi / 6
		colorStr      string  = "#000000"
	)
	xyscale := float64(width) / 2 / xyrange
	zscale := float64(height) * 0.4

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		switch k {
		case "width": // 横幅
			if len(v) > 0 {
				width_, _ := strconv.Atoi(v[0])
				if width_ > 0 {
					width = width_
				}
			}
		case "height": // 縦幅
			if len(v) > 0 {
				height_, _ := strconv.Atoi(v[0])
				if height_ > 0 {
					height = height_
				}
			}
		case "cells": // マス目
			if len(v) > 0 {
				cells_, _ := strconv.Atoi(v[0])
				if cells_ > 0 {
					cells = cells_
				}
			}
		case "xyrange": // x, yの範囲
			if len(v) > 0 {
				xyrange_, _ := strconv.ParseFloat(v[0], 64)
				if xyrange_ > 0 {
					xyrange = xyrange_
				}
			}
		case "angle": // 角度
			if len(v) > 0 {
				angle_, _ := strconv.ParseFloat(v[0], 64)
				angle = math.Pi * angle_ / 180
			}
		case "color": // 色
			if len(v) > 0 {
				colorStr = v[0]
			}
		}
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	drawSurface(w, cells, xyrange, float64(width), float64(height), xyscale, angle, zscale, colorStr)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}
