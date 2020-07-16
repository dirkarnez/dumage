package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math/rand"
	"os"
)

func main() {
	file, err := os.Create("test.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// width := 1920 * 8
	// height := 1080 * 8
	width := 3
	height := 3

	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: width, Y: height}})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{R: uint8(rand.Intn(256)), G: uint8(rand.Intn(256)), B: uint8(rand.Intn(256)), A: 0xff})
		}
	}
	err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
	if err != nil {
		fmt.Println(err)
	}
}
