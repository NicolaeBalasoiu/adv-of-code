package y2020

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const MINROW = 0
const MAXROW = 127
const MINCOLUMN = 0
const MAXCOLUMN = 7

func Day5() {
	file, err := os.Open("./y2020/input_day5.txt")
	seatIds := make([]int, 0)
	if err != nil {
		fmt.Println("cannot open file")
	}
	defer file.Close()
	chunkSize := 5 
	chunk := make([]string, 0, chunkSize)
	var result int
	scanner := bufio.NewScanner(file) 
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		chunk = append(chunk, line)

		if len(chunk) == chunkSize {
			aux := processChunk(chunk, &seatIds)
			if aux > result {
				result = aux
			}
			chunk = chunk[:0] // Clear chunk for reuse
		}
	}

	if len(chunk) > 0 {
		aux := processChunk(chunk, &seatIds)
		if aux > result {
			result = aux
		}	
	}

	fmt.Println(result, findMySeat(&seatIds))
}

func findMySeat(seatIds *[]int) int {
	sort.Ints(*seatIds)
	// fmt.Println(*seatIds)
	for i := 0; i < len(*seatIds)-1; i++ {
		if (*seatIds)[i+1]-(*seatIds)[i] > 1 {
			candidate := (*seatIds)[i] + 1
			// Check if both adjacent seats exist
			if contains(*seatIds, candidate-1) && contains(*seatIds, candidate+1) {
				return candidate
			}
		}
	}
	return 0 // No valid seat found
}

func contains(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}


func processChunk(chunk []string, seatIds *[]int) int {
	var max int
	for _, line := range chunk {
		row := line[:7]
		col := line[7:]
		rowNo := processRow(row)
		colNo := processCol(col)
		id := rowNo * 8 + colNo
		*seatIds = append(*seatIds, id)
		if id > max {
			max = id
		}
	}
	return max
}

func processRow(row string) int {
	var left, right = MINROW, MAXROW
	for _, v := range row {
		if (v == 'F') {
			right = (right + left) / 2
		} else {
			left = (right + left + 1) / 2
		}
	}
	return left
}

func processCol(col string) int {
	var left, right = MINCOLUMN, MAXCOLUMN
	for _, v := range col {
		if(v == 'R') {
			left = (right + left + 1) / 2
		} else {
			right = (right + left) / 2
		}
	}
	return left
}