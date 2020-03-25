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

func main() {
	x := 0
	y := 0
	width := 100
	height := 50

	img := image.NewRGBA(image.Rect(x, y, width, height))
	fillRect(img, color.RGBA{255, 0, 0, 0})

	file, _ := os.Create("sample.jpg")
	defer file.Close()

	if err := jpeg.Encode(file, img, &jpeg.Options{100}); err != nil {
		panic(err)
	}
}
