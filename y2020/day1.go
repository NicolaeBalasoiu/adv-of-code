package y2020

import (
	"fmt"
	"strconv"
	"strings"
)

func Day1(three_nos bool) (int, error) {
	file := readFile("./y2020/day1/input.txt")
	numbers := strings.Split(string(file), "\n")
	set_two :=  make(map[int]struct{})
	set_three :=  make(map[int]struct{})

	for i, v := range numbers {
		num, err := strconv.Atoi(v);
		if err != nil {
			fmt.Println("string to number conversion isnt possible")
		}
		if !three_nos {
			if _, diff := set_two[2020 - num]; diff {
				return num * (2020 - num), nil
			} else {
				set_two[num] = struct{}{}
			}
		} else {
			for j := i+1; j < len(numbers) - 1; j++ {
				convNr, _ := strconv.Atoi(numbers[j])
				needed := 2020 - (num + convNr)
				if _, exists := set_three[needed]; exists{
					fmt.Printf("Numbers: %d, %d, %d", needed, num, convNr )
					return needed * num * convNr, nil
				}
				set_three[convNr] = struct{}{}
			}
		}
	}
	return 0, fmt.Errorf("no numbers found in file %s", "input.txt" )
}

func sumArray(arr [2]int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}