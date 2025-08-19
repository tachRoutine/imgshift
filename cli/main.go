package main

import (
	"flag"
	"fmt"
	"strings"
)


func main(){
	args := flag.Args()
	toFormat := flag.String("format", "jpeg", "Output format (jpeg|png|gif)")
	flag.Parse()


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