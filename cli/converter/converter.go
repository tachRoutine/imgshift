package converter

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
)

const OUTPUT_DIR = "./imgshift_output/"

func init() {
	_, err := os.Stat(OUTPUT_DIR)
	if os.IsNotExist(err) {
		err = os.MkdirAll(OUTPUT_DIR, 0o755)
		if err != nil {
			fmt.Printf("Failed to create output directory: %v\n", err)
		}
	}
}

func Png2jpeg(input string) error {
	fmt.Println("Converting PNG to JPEG...")
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

	encodeErr := jpeg.Encode(jpegFile, pngImage, &jpeg.Options{Quality: 90})
	if encodeErr != nil {
		return fmt.Errorf("failed to encode JPEG image: %v", encodeErr)
	}
	fmt.Printf("Converting %s to %s format...\n", input, output)
	return nil
}

func Jpeg2png(input string) error {
	fmt.Println("Converting JPEG to PNG...")
	output := fmt.Sprintf("%s%s.png", OUTPUT_DIR, input)
	fmt.Printf("Converting %s to %s format...\n", input, output)
	return nil
}
