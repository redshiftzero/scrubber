package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xiam/exif"
)

func cleanImage(inputImagePtr *string) {
	data, _ := exif.Read(*inputImagePtr)
	for key, val := range data.Tags {
		fmt.Printf("%s = %s\n", key, val)
	}
	return
}

func main() {
	inputImagePtr := flag.String("image", "", "Image file to scrub metadata from")
	outputImagePtr := flag.String("output", "", "Output file of cleaned image")
	jsonOutputPtr := flag.Bool("json", false, "Print JSON metadata to stdout")
	flag.Parse()

	// Bail out if user did not provide an image
	if *inputImagePtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("image: ", *inputImagePtr)
	fmt.Println("output: ", *outputImagePtr)
	fmt.Println("jsonOutput: ", *jsonOutputPtr)

	cleanImage(inputImagePtr)
}
