package main

import (
	"flag"
	"fmt"
	"imgshift-cli/converter"
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

func getImageType(image string) (string, error) {
	if strings.HasSuffix(image, "jpeg") {
		return "jpeg", nil
	} else if strings.HasSuffix(image, "png") {
		return "png", nil
	} else if strings.HasSuffix(image, "gif") {
		return "gif", nil
	}
	return "", fmt.Errorf("unknown image format: %s", image)
}

func convertImageFormat(image string, toFormat string) error {
	switch toFormat {
	case "jpeg":
		return converter.Png2jpeg(image)
	case "png":
		return converter.Jpeg2png(image)
	default:
		return fmt.Errorf("unsupported format: %s", toFormat)
	}
}