package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	blackIndex = 0 // 背景
	greenIndex = 1 // 線の色
)

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	freq := rand.Float64() * 3.0 // 発振器yの相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意: エラーを無視
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		cycles  = 5     // 発振器xが完了する周回の回数
		res     = 0.001 // 回転の分解能
		size    = 100   // 画像キャンバスは[-size..+size]の範囲を扱う
		nframes = 64    // アニメーションフレーム数
		delay   = 8     // 10ms単位でのフレーム間の遅延
	)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if k == "cycles" {
			if cycles_, err := strconv.Atoi(v[0]); err == nil {
				cycles = cycles_
			} else {
				log.Print(err)
			}
		} else if k == "res" {
			if res_, err := strconv.ParseFloat(v[0], 64); err == nil {
				res = res_
			} else {
				log.Print(err)
			}
		} else if k == "size" {
			if size_, err := strconv.Atoi(v[0]); err == nil {
				size = size_
			} else {
				log.Print(err)
			}
		} else if k == "nframes" {
			if nframes_, err := strconv.Atoi(v[0]); err == nil {
				nframes = nframes_
			} else {
				log.Print(err)
			}
		} else if k == "delay" {
			if delay_, err := strconv.Atoi(v[0]); err == nil {
				delay = delay_
			} else {
				log.Print(err)
			}
		}
	}
	lissajous(w, cycles, res, size, nframes, delay)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}
