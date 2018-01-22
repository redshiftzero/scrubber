package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xiam/exif"
)

func getDefaultOutputFilename(inputImage string) string {
	inputFileName := filepath.Base(inputImage)
	outputStr := strings.Split(inputFileName, ".")
	outputFileName := outputStr[0] + "_clean." + outputStr[1]
	return outputFileName
}

func doCleaning(inputImagePtr *string, outputImagePtr *string, jsonOutputPtr *bool) error {
	// Check user provided an input file
	if *inputImagePtr == "" {
		err := errors.New("user did not provide an image")
		fmt.Println("[!]", err)
		return err
	}

	// Check user provided an actual file
	if _, err := os.Stat(*inputImagePtr); os.IsNotExist(err) {
		err := errors.New("file does not exist")
		fmt.Println("[!]", err)
		return err
	}

	// Name output file if not named
	if *outputImagePtr == "" {
		*outputImagePtr = getDefaultOutputFilename(*inputImagePtr)
	}

	// Save output in file
	fmt.Println("image: ", *inputImagePtr)
	fmt.Println("output: ", *outputImagePtr)
	fmt.Println("jsonOutput: ", *jsonOutputPtr)

	// Optionally output metadata as JSON
	data, _ := exif.Read(*inputImagePtr)
	for key, val := range data.Tags {
		fmt.Printf(".")
		fmt.Printf("%s = %s\n", key, val)
	}
	return nil
}

func main() {
	inputImagePtr := flag.String("image", "", "Image file to scrub metadata from")
	outputImagePtr := flag.String("output", "", "Output file of cleaned image (optional)")
	jsonOutputPtr := flag.Bool("json", false, "Print JSON metadata to stdout (optional)")
	flag.Parse()
	err := doCleaning(inputImagePtr, outputImagePtr, jsonOutputPtr)
	if err != nil {
		os.Exit(1)
	}
}
