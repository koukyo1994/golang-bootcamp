package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320            // キャンバスの大きさ (画素数)
	cells         = 100                 // 格子のマス目の数
	xyrange       = 30.0                // 軸の範囲 (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x単位およびy単位あたりの画素数
	zscale        = height * 0.4        // z単位あたりの画素数
	angle         = math.Pi / 6         // x, y軸の角度 (30°)
)

const (
	MAX = 255 // HSV -> RGBに使われるMAX値
	MIN = 0   // HSV -> RGBに使われるMIN値
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func corner(i, j int) (float64, float64, float64, float64) {
	// マス目 (i, j)の角の点(x, y)を見つける。
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する。
	z := f(x, y)
	r := math.Hypot(x, y)

	// (x, y, z)を2-D SVGキャンバス(sx, sy)へ等角的に投影。
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
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

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, ar := corner(i+1, j)
			bx, by, bz, br := corner(i, j)
			cx, cy, cz, cr := corner(i, j+1)
			dx, dy, dz, dr := corner(i+1, j+1)

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
			fillColor := colorToString(decideColor((az+bz+cz+dz)/4, (ar+br+cr+dr)/4))

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s' stroke='black'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, fillColor)
		}
	}
	fmt.Println("</svg>")
}
