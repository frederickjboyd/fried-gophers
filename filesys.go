package main

import (
	"image"
	"log"
	"os"
)

func openImage(path string) (img image.Image, format string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}
	log.Print(file)
	return
}
