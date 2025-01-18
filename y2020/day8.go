package y2020

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	inst string
	no int
	visited bool
}

func Day8() {
	lines := readInput("./y2020/input_day8.txt")
	data := structureData(lines)
	acc := searchFault(data)
	fmt.Println(acc)
}

func structureData(data []string) []instruction {
	instructions := make([]instruction, 0)
	for _, line := range data {
		s := strings.Split(line, " ")
		insType := s[0]
		insNo, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Printf("Can't convert %s to int", s[1])
		}
		instructions = append(instructions, instruction{inst: insType, no: insNo, visited: false})
	}

	return instructions
}

func runInstructions(data *[]instruction) int {
	var acc, i int
	for {
		if (*data)[i].visited {
			return acc
		}
		(*data)[i].visited = true
		if (*data)[i].inst == "acc" {
			acc += (*data)[i].no
			i++
		} else if (*data)[i].inst == "jmp" {
			i += (*data)[i].no
		} else {
			i++
		}
	}
}

func runInstructionsSecondHalf(data []instruction) (int, bool) {
	var acc, i int
	visited := make(map[int]bool)

	for {
		if visited[i] {
			return 0, false
		}
		visited[i] = true
		if (data)[i].inst == "acc" {
			acc += (data)[i].no
			i++
		} else if (data)[i].inst == "jmp" {
			i += (data)[i].no
		} else {
			i++
		}
		if i == len(data) - 1 {
			return acc, true
		}
	}
}

func searchFault(data []instruction) int {
	
	for i, inst := range data {
		if inst.inst == "acc" {
			continue
		}
		originalOp := data[i].inst
		if originalOp == "jmp" {
			data[i].inst = "nop"
		} else {
			data[i].inst = "jmp"
		}
		acc, ok := runInstructionsSecondHalf(data)

		if ok {
			return acc
		}
		data[i].inst = originalOp
	}
	return 0
}