package bucketcompare

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"math"
	"os"
)

func CalculateVector(imagePath string) {

	bucketCount := 9

	//Size in percentage of imagesize
	bucketSize := 0.02

	file, err := os.Open(imagePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		panic(err)
	}

	fmt.Print("Filename", file.Name())
	fmt.Println("Imageheight ", img.Bounds().Max.X, " width ", img.Bounds().Max.Y)

	imgHeight := img.Bounds().Max.Y
	imgWidth := img.Bounds().Max.X

	bucketHeight := int(math.Floor(float64(imgHeight) * bucketSize))
	bucketWidth := int(math.Floor(float64(imgWidth) * bucketSize))

	heightSpacing := int(math.Floor(float64(imgHeight / bucketCount)))
	widthSpacing := int(math.Floor(float64(imgWidth / bucketCount)))

	heightOffset := int(heightSpacing/2) - int(bucketHeight/2)
	for i := 0; i < bucketCount; i++ {

		widthOffset := int(widthSpacing / 2)
		for j := 0; j < bucketCount; j++ {
			widthOffset += widthSpacing - int(bucketWidth/2)
			bucketGray := averageGray(img, widthOffset, heightOffset, bucketHeight, bucketWidth)
			fmt.Println("Bucket woff", widthOffset, " and hoff ", heightOffset, " and gray ", bucketGray)
		}

		heightOffset += heightSpacing
	}

	fmt.Println("BH ", bucketHeight, " BW ", bucketWidth)

}

func averageGray(img image.Image, widthOffset int, heightOffset int, height int, width int) int {
	var grayVal = make([]int, height*width)

	arrIndex := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {

			//r, g, b, a := img.At(widthOffset+j, heightOffset+i).RGBA()
			pixel := img.At(widthOffset+j, heightOffset+i)
			grayVal[arrIndex] = RGBAToGray(pixel.RGBA())
			arrIndex += 1
		}
	}

	av := 0
	for i := 0; i < len(grayVal); i++ {
		av += grayVal[i]
	}

	return int(av / len(grayVal))
}

func RGBAToGray(r, g, b, a uint32) int {
	gray := uint8((0.299 * float64(r)) + (0.587 * float64(g)) + (0.114 * float64(b)))
	return int(gray)
}

func compareVectors() {
	return
}

func compareImages(img1 string, img2 string) {

}
