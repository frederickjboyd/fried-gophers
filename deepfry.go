package main

import (
	"image"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/noise"
)

func adjustSaturation(img image.Image, val float64) (result image.Image) {
	result = adjust.Saturation(img, val)
	return
}

func genNoise() (result image.Image) {
	result = noise.Generate(280, 280, &noise.Options{Monochrome: true, NoiseFn: noise.Gaussian})
	return
}
