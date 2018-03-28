package main

import (
	"encoding/hex"
	"flag"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"
)

func main() {
	imagePath := flag.String("image", "", "Absolute path to image file.")
	flag.Parse()

	file, err := os.Open(*imagePath)
	if err != nil {
		log.Fatal(err)
	}

	var decoded image.Image
	decoded, _, err = image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}

	var (
		rec                    = decoded.Bounds()
		min, max               = rec.Min, rec.Max
		minX, maxX, minY, maxY = min.X, max.X, min.Y, max.Y
		currentColour          color.Color
		count                  int
		temp, output           string
	)

	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			colour := decoded.At(x, y)
			if currentColour != colour {
				if count > 0 {
					if count == '-' {
						hexConv, err := hex.DecodeString(temp)
						if err != nil {
							log.Fatal(err)
						}
						output += string(hexConv[0])
						temp = ""
					} else {
						temp += string(count)
					}
				}
				currentColour = colour
				count = 0
			}
			count++
		}
	}

	log.Print(output)
}
