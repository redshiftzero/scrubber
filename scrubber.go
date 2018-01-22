package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Jeffail/gabs"
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

	// Name output file if no name provided by user
	if *outputImagePtr == "" {
		*outputImagePtr = getDefaultOutputFilename(*inputImagePtr)
	}

	// Save output in file
	fmt.Println("Clean version of", *inputImagePtr, "saved in the current directory as", *outputImagePtr)

	// Optionally output metadata as JSON
	if *jsonOutputPtr == false {
		return nil
	}
	data, _ := exif.Read(*inputImagePtr)
	jsonObj := gabs.New()
	jsonObj.Set(*inputImagePtr, "Filename")
	for key, val := range data.Tags {
		jsonObj.Set(val, key)
	}
	fmt.Println(jsonObj.StringIndent("", "  "))
	return nil
}

func main() {
	inputImagePtr := flag.String("input", "", "Image file to scrub metadata from")
	outputImagePtr := flag.String("output", "", "Output file of cleaned image (optional)")
	jsonOutputPtr := flag.Bool("json", false, "Print JSON metadata to stdout (optional)")
	flag.Parse()
	err := doCleaning(inputImagePtr, outputImagePtr, jsonOutputPtr)
	if err != nil {
		os.Exit(1)
	}
}
