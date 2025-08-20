package main

import (
	"flag"
	"fmt"
	"strings"
)


func main(){
	toFormat := flag.String("format", "jpeg", "Output format (jpeg|png|gif)")
	flag.Parse()
	args := flag.Args()
	path := args[0]
	
	if err := validateImgFile(path); err != nil {
		fmt.Println(err)
		return
	}

	convertImageFormat(path, *toFormat)

}

func validateImgFile(path string) error {
	allowedFormats := []string{"jpeg", "png", "gif"}
	for _, format := range allowedFormats{
		if strings.HasSuffix(path, format) {
			return nil
		}
	}
	return fmt.Errorf("invalid image format: %s", path)
}

func convertImageFormat(image string, toFormat string) error {
	fmt.Printf("Converting %s to %s format...\n", image, toFormat)
	return nil
}