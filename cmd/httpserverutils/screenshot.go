package httpserverutils

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

// TODO
func GetScreenshotAroundCoords(x int, y int) {

	// Capture the entire screen
	bounds := screenshot.GetDisplayBounds(0) // Assuming only one monitor
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		fmt.Println("Failed to capture screen:", err)
		return
	}

	// Define the dimensions for the cropped image
	width := 200
	height := 200

	// Calculate the bounds for the cropped area
	cropBounds := image.Rect(x-width/2, y-height/2, x+width/2, y+height/2)

	// Crop the image
	croppedImg := img.SubImage(cropBounds)

	// Save the cropped image to a file
	outputFile, err := os.Create("screenshot.png")
	if err != nil {
		fmt.Println("Failed to create output file:", err)
		return
	}
	defer outputFile.Close()

	// Encode the cropped image as PNG and save it to the file
	err = png.Encode(outputFile, croppedImg)
	if err != nil {
		fmt.Println("Failed to encode image:", err)
		return
	}

	fmt.Println("Screenshot captured and saved successfully!")

}
