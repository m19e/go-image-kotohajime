package main // import "gen"

import (
	"image"
	"image/color"
	"image/jpeg"
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

func main() {
	x := 0
	y := 0
	width := 100
	height := 50

	img := image.NewRGBA(image.Rect(x, y, width, height))
	fillRect(img, color.RGBA{255, 255, 255, 0})
	drawBounds(img, color.RGBA{255, 0, 0, 0})

	file, _ := os.Create("sample.jpg")
	defer file.Close()

	if err := jpeg.Encode(file, img, &jpeg.Options{100}); err != nil {
		panic(err)
	}
}
