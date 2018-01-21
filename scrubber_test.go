package main

import "testing"

func TestImage(t *testing.T) {
	test_image := "test/test.jpg"
	imagePtr := &test_image
	cleanImage(imagePtr)
}
