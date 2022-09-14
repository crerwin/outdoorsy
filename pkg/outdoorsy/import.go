package outdoorsy

import (
	"fmt"
	"log"
	"os"
)

func readFile(filename string) []byte {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file %v with error %v", filename, err)
	}

	return content
}

func ReadFiles(filenames []string) {
	var content string
	for _, filename := range filenames {
		content += string(readFile(filename))
		content += "\n"
	}
	fmt.Println(content)
}
