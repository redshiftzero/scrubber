package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
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

func doCleaning(inputImagePtr *string, cleanImagePtr *bool, outputImagePtr *string, jsonOutputPtr *bool) error {
	// Check user provided an input file
	if *inputImagePtr == "" {
		err := errors.New("user did not provide an image")
		flag.Usage()
		return err
	}

	// Check user provided an actual file
	if _, err := os.Stat(*inputImagePtr); os.IsNotExist(err) {
		return err
	}

	inputFile, err := os.Open(*inputImagePtr)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	imageData, imageType, err := image.Decode(inputFile)
	if err != nil {
		return err
	}
	inputFile.Seek(0, 0)

	// Name output file if no name provided by user
	if *outputImagePtr == "" {
		*outputImagePtr = getDefaultOutputFilename(*inputImagePtr)
	}

	if *cleanImagePtr == true {
		// See if image type is not supported
		if imageType != "png" && imageType != "jpeg" && imageType != "jpg" {
			fmt.Println(imageType)
			err := errors.New("image type not yet supported")
			return err
		}

		// Create output file
		outputFile, err := os.Create(*outputImagePtr)
		if err != nil {
			return err
		}

		if imageType == "png" {
			png.Encode(outputFile, imageData)
		}
		if imageType == "jpeg" || imageType == "jpg" {
			jpeg.Encode(outputFile, imageData, nil)
		}
		fmt.Println("Clean version of", *inputImagePtr, "saved in the current directory as", *outputImagePtr)
		outputFile.Close()
	}

	// Optionally output metadata as JSON
	if *jsonOutputPtr == false {
		return nil
	}
	data, _ := exif.Read(*inputImagePtr)
	jsonObj := gabs.New()
	jsonObj.Set(*inputImagePtr, "Filename")
	//fmt.Println(data)
	if data != nil {
		for key, val := range data.Tags {
			jsonObj.Set(val, key)
		}
	}
	fmt.Println(jsonObj.StringIndent("", "  "))
	return nil
}

func main() {
	inputImagePtr := flag.String("input", "", "Image file to scrub metadata from")
	cleanImagePtr := flag.Bool("clean", true, "Generate cleaned image (optional)")
	outputImagePtr := flag.String("output", "", "Output file of cleaned image (optional)")
	jsonOutputPtr := flag.Bool("json", false, "Print JSON metadata to stdout (optional)")
	flag.Parse()
	err := doCleaning(inputImagePtr, cleanImagePtr, outputImagePtr, jsonOutputPtr)
	if err != nil {
		fmt.Println("[!]", err)
		os.Exit(1)
	}
}
