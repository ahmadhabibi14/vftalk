package utils

import (
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}
