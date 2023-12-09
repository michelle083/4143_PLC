
// Package Declaration
package main

// Import required packages
import (
	"fmt"

	"github.com/michelle083/img_mod/Colors"
	"github.com/michelle083/img_mod/GetPic"
	"github.com/michelle083/img_mod/Grayscale"
	"github.com/michelle083/img_mod/Text"
)

func main() {
	// Prints Heading of the Program
	fmt.Println("Michelle Orru")
	fmt.Println("P03 - Image Ascii Art")

	// Prints black space
	fmt.Print("\n")

	// Call function to get picture from URL
	GetPic.DownloadPicture()

	// Prints blank space
	fmt.Print("\n")

	// Call function to process Pixel Colors
	Colors.PrintPixels()

	// Prints blank space
	fmt.Print("\n")

	// Call function to grayscale image
	Grayscale.GrayScale()

	// Prints blank space
	fmt.Print("\n")

	// Call function to print colored text to image
	Text.PrintText()
}
