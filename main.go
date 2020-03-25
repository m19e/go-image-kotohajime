package main // import "gen"

import (
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
)

func fillRect(img *image.RGBA, col color.Color) {
	rect := img.Rect

	for h := rect.Min.Y; h < rect.Max.Y; h++ {
		for v := rect.Min.X; v < rect.Max.X; v++ {
			img.Set(v, h, col)
		}
	}
}

func drawBounds(img *image.RGBA, col color.Color) {
	rect := img.Rect

	// 上下の枠
	for h := 0; h < rect.Max.X; h++ {
		img.Set(h, 0, col)
		img.Set(h, rect.Max.Y-1, col)
	}

	// 左右の枠
	for v := 0; v < rect.Max.Y; v++ {
		img.Set(0, v, col)
		img.Set(rect.Max.X-1, v, col)
	}
}

type Circle struct {
	p image.Point
	r int
}

func (c *Circle) drawCircle(img *image.RGBA, col color.Color) {
	for rad := 0.0; rad < 2.0*float64(c.r); rad += 0.1 {
		x := int(float64(c.p.X) + float64(c.r)*math.Cos(rad))
		y := int(float64(c.p.Y) + float64(c.r)*math.Sin(rad))
		img.Set(x, y, col)
	}
}

func main() {
	x := 0
	y := 0
	width := 500
	height := 500

	img := image.NewRGBA(image.Rect(x, y, width, height))
	// dye White
	fillRect(img, color.RGBA{255, 255, 255, 0})

	// draw Bounds
	// drawBounds(img, color.RGBA{255, 0, 0, 0})

	// draw Circle
	// center point
	center := image.Point{250, 250}
	// create circle
	circle := Circle{center, 50}
	// Draw
	circle.drawCircle(img, color.RGBA{255, 0, 0, 0})

	file, _ := os.Create("sample.jpg")
	defer file.Close()

	if err := jpeg.Encode(file, img, &jpeg.Options{100}); err != nil {
		panic(err)
	}
}
