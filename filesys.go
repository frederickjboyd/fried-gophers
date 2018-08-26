package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func openImage(path string) (img image.Image, format string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}
	img, format, err = image.Decode(file)
	if err != nil {
		log.Fatalf("Failed to decode file: %v", err)
	}
	return
}

func writeImage(img image.Image, name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatalf("Unable to create new file: %v", err)
	}
	jpeg.Encode(file, img, nil)
}

// Creates a test image with name "name" and dimensions x by y
func createTestImage(name string, x, y int) {
	newFile, err := os.Create(name)
	newImage := image.NewRGBA(image.Rect(0, 0, x, y))
	defer newFile.Close()
	if err != nil {
		log.Fatalf("Error creating new file")
	}
	jpeg.Encode(newFile, newImage, nil)
}
