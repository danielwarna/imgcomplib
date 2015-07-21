package main

import "github.com/danielwarna/comparelib/bucketcompare"

func main() {
	bucketcompare.CalculateVector("testimages/test1.jpg")

	bucketcompare.CalculateVector("testimages/test1_large.jpg")

}
