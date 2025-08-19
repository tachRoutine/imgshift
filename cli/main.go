package main

import (
	"flag"
	"fmt"
	"strings"
)


func main(){
	args := flag.Args()
	toFormat := flag.String("format", "jpeg", "Output format (jpeg|png|gif)")
	fmt.Println("Arguments passed:", args)
	path := args[0]
	flag.Parse()

	if err := validateImgFile(path); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("converting to", *toFormat)

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