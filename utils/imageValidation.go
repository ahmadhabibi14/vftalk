package utils

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"strings"
)

func ImageValidation(file *multipart.FileHeader) error {
	f, err := file.Open()
	if err != nil {
		return errors.New("Could not open uploaded file")
	}
	defer f.Close()

	// Check that file is an image
	img, format, err := image.DecodeConfig(f)
	if err != nil {
		return errors.New("Uploaded file is not an image")
	}

	// Check that the file extension is allowed
	if !allowedExtension(format) {
		return errors.New("Uploaded file extension is not allowed")
	}

	// Check that the image has the same dimension
	if img.Width != img.Height {
		return errors.New("Uploaded image does not have the same dimension, recommend to 500x500")
	}

	// Check that the image meets the minimum and maximum pixel dimensions
	if (img.Width < minPixelWidth) && (img.Height < minPixelHeight) {
		return errors.New(fmt.Sprintf("Uploaded image does not meet minimum pixel dimensions of %dx%d", minPixelWidth, minPixelHeight))
	}
	if (img.Width > maxPixelWidth) && (img.Height > maxPixelHeight) {
		return errors.New(fmt.Sprintf("Uploaded image exceeds maximum pixel dimensions of %dx%d", maxPixelWidth, maxPixelHeight))
	}

	// Check if the file size is not too large
	max := float64(maxFileSize) / 1000000
	if file.Size > maxFileSize {
		return errors.New(fmt.Sprintf("Uploaded file exceeds maximum size of %.2f MB", max))
	}

	return nil
}

const (
	minPixelWidth  = 100
	minPixelHeight = 100
	maxPixelWidth  = 1000
	maxPixelHeight = 1000
	maxFileSize    = 1024 * 1024 * 5 // 5 MB
)

func allowedExtension(format string) bool {
	for _, ext := range allowedExtensions {
		if strings.EqualFold(ext, format) {
			return true
		}
	}
	return false
}

var allowedExtensions = []string{
	"jpeg", "jpg", "png",
}
