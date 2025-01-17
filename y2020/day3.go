package y2020

import (
	"fmt"
)

func Day3() {
	file := readFile("./y2020/input_day3.txt")
	data := splitFile(file)
	fmt.Println(
		first_half_day3(data, 1, 1) *
		first_half_day3(data, 3, 1) *
		first_half_day3(data, 5, 1) *
		first_half_day3(data, 7, 1) * 
		first_half_day3(data, 1, 2),
	)

}

func first_half_day3(data []string, right int, down int) int {
	var trees int
	var i, j int
	for j < len(data) {
		if data[j][i] == '#' {
			trees ++
		}
		i = (i+right) % len(data[0])
		j+= down
	}
	return trees
}