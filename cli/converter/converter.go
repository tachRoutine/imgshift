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

	pngImage, err := png.Decode(image)
	if err != nil {
		return err
	}

	jpegFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer jpegFile.Close()
	defer image.Close()
	fmt.Printf("Converting %s to %s format...\n", input, output)
	return nil
}

func Jpeg2png(input string) error {
	output := fmt.Sprintf("%s%s.png", OUTPUT_DIR, input)
	fmt.Printf("Converting %s to %s format...\n", input, output)
	return nil
}

