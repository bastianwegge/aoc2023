package util

import (
	"log"
	"os"
)

func ReadFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
