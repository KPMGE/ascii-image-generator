package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	asciigen "github.com/ascii-image-generator/ascii-gen"
)

func main() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	if len(os.Args) < 2 {
		fmt.Println("usage: executable <image name>")
		os.Exit(1)
	}

	imageName := os.Args[1]
	file, err := os.Open(imageName)

	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}

	defer file.Close()

	pixels, err := asciigen.GetPixels(file)

	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	asciigen.DrawAscii(pixels)
}
