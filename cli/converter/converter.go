package converter

import (
	"fmt"
	"image/png"
	"os"
)

const OUTPUT_DIR = "./output/"

func Png2jpeg(input string) error {
	output := fmt.Sprintf("%s%s.jpeg", OUTPUT_DIR, input)
	image, err := os.Open(input)
	if err != nil {
		return err
	}
	defer image.Close()

	pngImage, err := png.Decode(image)
	if err != nil {
		return fmt.Errorf("failed to decode PNG image: %v", err)
	}

	jpegFile, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("failed to create JPEG file: %v", err)
	}
	defer jpegFile.Close()

	if err := jpeg.Encode(jpegFile, pngImage, nil); err != nil {
		return fmt.Errorf("failed to encode JPEG image: %v", err)
	}

	fmt.Printf("Converting %s to %s format...\n", input, output)
	
	return nil
}

func Jpeg2png(input string) error {
	output := fmt.Sprintf("%s%s.png", OUTPUT_DIR, input)
	fmt.Printf("Converting %s to %s format...\n", input, output)
	return nil
}

