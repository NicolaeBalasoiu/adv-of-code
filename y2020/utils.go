package y2020

import (
	"bufio"
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

func readInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

