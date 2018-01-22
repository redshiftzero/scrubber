package main

import (
	"testing"
)

var outputFilenameTests = []struct {
	in  string
	out string
}{
	{"test.jpg", "test_clean.jpg"},
	{"test.png", "test_clean.png"},
}

func TestGetDefaultOutputFilename(t *testing.T) {
	for _, tt := range outputFilenameTests {
		s := getDefaultOutputFilename(tt.in)
		if s != tt.out {
			t.Errorf("getDefaultOutputFilename(%q) => %q, want %q", tt.in, s, tt.out)
		}
	}
}

func TestProgramCompletesIfInputFile(t *testing.T) {
	testImage := "test/test.jpg"
	inputImagePtr := &testImage
	testOutput := ""
	outputImagePtr := &testOutput
	testJsonOutputFlag := false
	jsonOutputPtr := &testJsonOutputFlag

	err := doCleaning(inputImagePtr, outputImagePtr, jsonOutputPtr)
	if err != nil {
		t.Errorf("doCleaning(%q, _, _) => %q, want %q", testImage, err, "nil")
	}
}

func TestProgramBailsIfNoInputFile(t *testing.T) {
	testImage := ""
	inputImagePtr := &testImage
	testOutput := ""
	outputImagePtr := &testOutput
	testJsonOutputFlag := false
	jsonOutputPtr := &testJsonOutputFlag

	err := doCleaning(inputImagePtr, outputImagePtr, jsonOutputPtr)
	if err == nil {
		t.Errorf("doCleaning(%q, _, _) => %q, want %q", testImage, err, err)
	}
}

func TestProgramBailsIfInputFileDoesNotExist(t *testing.T) {
	testImage := "test/doesnotexist.jpg"
	inputImagePtr := &testImage
	testOutput := ""
	outputImagePtr := &testOutput
	testJsonOutputFlag := false
	jsonOutputPtr := &testJsonOutputFlag

	err := doCleaning(inputImagePtr, outputImagePtr, jsonOutputPtr)
	if err == nil {
		t.Errorf("doCleaning(%q, _, _) => %q, want %q", testImage, err, "nil")
	}
}

func TestProgramJsonOutput(t *testing.T) {
	testImage := "test/test.jpg"
	inputImagePtr := &testImage
	testOutput := ""
	outputImagePtr := &testOutput
	testJsonOutputFlag := true
	jsonOutputPtr := &testJsonOutputFlag

	err := doCleaning(inputImagePtr, outputImagePtr, jsonOutputPtr)
	if err != nil {
		t.Errorf("doCleaning(%q, _, _) => %q, want %q", testImage, err, "nil")
	}
}
