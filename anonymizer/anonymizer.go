package anonymizer

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

var outPath = "./output/"

func RunAnonymize(fileMap map[string][]string, topCrop int, rightCrop int, bottomCrop int, leftCrop int) {
	for path, files := range fileMap {
		for _, file := range files {
			anonymizeImage(path, file, topCrop, rightCrop, bottomCrop, leftCrop)
		}
	}
}

func anonymizeImage(path string, file string, topCrop int, rightCrop int, bottomCrop int, leftCrop int) {
	currentImageData, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}
	defer currentImageData.Close()

	// Determine image format
	format := strings.ToLower(filepath.Ext(file))

	// Decode image based on format
	var img image.Image
	switch format {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(currentImageData)
	case ".png":
		img, err = png.Decode(currentImageData)
	default:
		fmt.Printf("Unsupported image format: %s\n", format)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error decoding image: %s\n", err)
		os.Exit(1)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	fmt.Printf("Image dimensions: Width: %d, Height: %d\n", width, height)

	// Create a new RGBA image with the same dimensions as the original
	rgbaImg := image.NewRGBA(bounds)

	// Define purple color
	purple := color.RGBA{R: 128, G: 0, B: 128, A: 255}

	// Copy the image and add purple pixels to cropped areas
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x < leftCrop || x >= width-rightCrop || y < topCrop || y >= height-bottomCrop {
				rgbaImg.Set(x, y, purple)
			} else {
				rgbaImg.Set(x, y, img.At(x, y))
			}
		}
	}

	// Create a new buffer to store the encoded image
	imgBytes := bytes.NewBuffer([]byte{})

	// Encode image based on format
	switch format {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(imgBytes, rgbaImg, nil)
	case ".png":
		err = png.Encode(imgBytes, rgbaImg)
	}

	if err != nil {
		fmt.Printf("Error encoding image: %s\n", err)
		os.Exit(1)
	}

	// Ensure the output directory exists
	err = os.MkdirAll(filepath.Dir(outPath+file), 0755)
	if err != nil {
		fmt.Printf("Error creating output directory: %s\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(outPath+"out_"+file, imgBytes.Bytes(), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %s\n", err)
		os.Exit(1)
	}
}
