package main

import (
	"flag"
	"fmt"
)


func main(){
	args := flag.Args()
	fmt.Println("Arguments passed to the program:")
	flag.Parse()

	if len(args) == 0{
		fmt.Println("No argument provided")
		return
	}
	
}