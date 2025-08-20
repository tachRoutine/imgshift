package main

import (
	"flag"
	"fmt"
	"imgshift-cli/converter"
	"path/filepath"
	"strings"
)

func main() {
	toFormat := flag.String("to", "jpeg", "Output format (jpeg|png|gif)")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Error: Please provide an input image file")
		return
	}

	inputPath := flag.Args()[0]

	inputFormat, err := validateInputFile(inputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	outputFormat := strings.ToLower(*toFormat)
	if err := validateOutputFormat(outputFormat); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if inputFormat == outputFormat {
		fmt.Printf("Input file is already in %s format\n", outputFormat)
		return
	}

	if err := convertImageFormat(inputPath, inputFormat, outputFormat); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Conversion completed successfully!")
}

func validateInputFile(path string) (string, error) {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(path), "."))

	validFormats := map[string]bool{
		"jpeg": true,
		"jpg":  true,
		"png":  true,
		"gif":  true,
	}

	if !validFormats[ext] {
		return "", fmt.Errorf("unsupported input format: %s. Supported formats are: jpeg, jpg, png, gif", ext)
	}

	if ext == "jpg" {
		ext = "jpeg"
	}

	return ext, nil
}

func validateOutputFormat(format string) error {
	validFormats := map[string]bool{
		"jpeg": true,
		"png":  true,
		"gif":  true,
	}

	if !validFormats[format] {
		return fmt.Errorf("unsupported output format: %s. Supported formats are: jpeg, png, gif", format)
	}

	return nil
}

func convertImageFormat(inputPath, inputFormat, outputFormat string) error {
	fmt.Printf("Converting %s to %s format...\n", inputFormat, outputFormat)

	switch {
	case inputFormat == "png" && outputFormat == "jpeg":
		return converter.Png2jpeg(inputPath)
	case inputFormat == "jpeg" && outputFormat == "png":
		return converter.Jpeg2png(inputPath)
	case outputFormat == "gif":
		return fmt.Errorf("gif conversion not yet implemented")
	default:
		return fmt.Errorf("conversion from %s to %s is not supported", inputFormat, outputFormat)
	}
}
