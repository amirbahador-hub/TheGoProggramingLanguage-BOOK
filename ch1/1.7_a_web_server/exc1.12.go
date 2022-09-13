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

var paletteExc = []color.Color{color.RGBA{R: 0, G: 255, B: 0, A: 1}, color.RGBA{R: 0, G: 0, B: 255, A: 1}}

func main() {
	http.HandleFunc("/", handlerExc)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handlerExc(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	var cycl int
	var err error
	if cycles, ok := r.Form["cycles"]; ok {
		cycl, err = strconv.Atoi(cycles[0])
		if err != nil {
			cycl = 5
			log.Print(err)
		}
	} else {
		cycl = 5
	}
	lissajousExc(w, cycl)
}

const (
	greenIndex = 0
	blueIndex  = 1
)

func lissajousExc(out io.Writer, cycles int) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, paletteExc)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blueIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
