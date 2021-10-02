package main

import (
	"flag"
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // キャンバスの大きさ (画素数)
	cells         = 100                 // 格子のマス目の数
	xyrange       = 30.0                // 軸の範囲 (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x単位およびy単位あたりの画素数
	zscale        = height * 0.4        // z単位あたりの画素数
	angle         = math.Pi / 8         // x, y軸の角度 (30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var shape = flag.String("shape", "surface after water drops", "shape to draw")

func corner(fn func(float64, float64) float64, i, j int) (float64, float64) {
	// マス目 (i, j)の角の点(x, y)を見つける。
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する。
	z := fn(x, y)

	// (x, y, z)を2-D SVGキャンバス(sx, sy)へ等角的に投影。
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func surfaceAfterWaterDrops(x, y float64) float64 {
	r := math.Hypot(x, y) // (x, y)の距離
	return math.Sin(r) / r
}

func saddle(x, y float64) float64 {
	return 0.001 * (2*x*x - 3*y*y)
}

func bump(x, y float64) float64 {
	return 0.05 * (math.Pow(math.Sin(x/2), 2) + math.Pow(math.Sin(y), 2))
}

func eggPack(x, y float64) float64 {
	return 0.3 * (math.Pow(math.Sin(x/2), 2) + math.Pow(math.Sin(y/2), 2))
}

func main() {
	flag.Parse()
	var fn func(float64, float64) float64
	if *shape == "surface after water drops" {
		fn = surfaceAfterWaterDrops
	} else if *shape == "saddle" {
		fn = saddle
	} else if *shape == "bump" {
		fn = bump
	} else if *shape == "egg pack" {
		fn = eggPack
	} else {
		panic("unknown shape")
	}
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(fn, i+1, j)
			bx, by := corner(fn, i, j)
			cx, cy := corner(fn, i, j+1)
			dx, dy := corner(fn, i+1, j+1)

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

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill-opacity='0.2' fill='grey'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}
