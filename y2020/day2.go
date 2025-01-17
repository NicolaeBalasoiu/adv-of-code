package y2020

import (
	"fmt"
	"strconv"
	"strings"
)

type line struct {
	limit [2]int
	letter rune
	input string
}

func Day2() {
	file := readFile("./y2020/input_day2.txt")
	lines := splitIntoLines(file)
	var data = splitLines(lines)
	// first_half(data)
	second_half(data)
}

func first_half (data []line) int {
	var valid int
	for _, v := range data {
		count := 0
		for _, char :=  range v.input {
			if char == v.letter {
				count ++
			}
		}
		if count >= v.limit[0] && count <= v.limit[1] {
			valid ++
		}
	}
	fmt.Println(valid)
	return valid
}

func second_half(data[] line) int {
	var valid int
	for _, v := range data {
		count := 0
		for i, char := range v.input {
			if (i == v.limit[0]-1 || i == v.limit[1]-1) && char == v.letter  {
				count ++
			}
		}
		if count == 1 {
			valid ++
		}
	}
	fmt.Println(valid)
	return valid
}

func splitIntoLines(file []byte) []string {
	lines := strings.Split(string(file), "\n")
	return lines
}

func splitLines(lines []string) []line {
	result := make([]line, 0)
	for _, v := range lines {
		splitbySpace := strings.Split(v, " ")
		splitbyDash := strings.Split(splitbySpace[0], "-")
		lowerLimit, err := strconv.Atoi(splitbyDash[0])
		if err != nil {
			panic(err)
		}
		upperlimit, err := strconv.Atoi(splitbyDash[1])
		if err != nil {
			panic(err)
		}

		limit := [2]int{lowerLimit, upperlimit}
		letter := rune(splitbySpace[1][0])
		input := splitbySpace[2]
		result = append(result, line{limit: limit, letter: letter, input: input})
	}
	return result
}