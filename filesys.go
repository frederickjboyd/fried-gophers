package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
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

func writeImage(img image.Image, name string, dir string, quality int) {
	if len(dir) == 0 {
		dir = filepath.Dir(args[0])
		log.Println(dir)
	}
	if opts.Verbose {
		log.Printf("writing to %s", dir)
	}
	file, err := os.Create(filepath.Join(dir, filepath.Base(name)))
	if err != nil {
		log.Fatalf("Unable to create new file: %v", err)
	}
	jpeg.Encode(file, img, &jpeg.Options{Quality: quality})
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
