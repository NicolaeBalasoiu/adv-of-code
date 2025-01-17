package y2020

import (
	"fmt"
	"os"
	"strings"
)

func readFile(path string) []byte {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("error when opening the file")
		panic(err)
	}

	return file
}

func splitFile(file []byte) []string {
	return strings.Split(strings.Trim(string(file), " "), "\n")
}

