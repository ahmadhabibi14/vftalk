package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"strings"
)

func GenerateRandomID(size int) string {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		log.Panicln(err)
	}
	// Encode the random number to a base64 string
	encode := base64.StdEncoding.EncodeToString(b)
	replacer := strings.NewReplacer(
		"&", "",
		"-", "",
		"+", "",
		"=", "",
		"!", "",
		"/", "",
		`\`, "",
		"#", "",
		"*", "",
		"%", "",
	)
	id := replacer.Replace(encode)
	return id
}
