package y2020

import (
	"bufio"
	"fmt"
	"os"
)

func Day6() {
	file, err := os.Open("./y2020/input_day6.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close();
	var sum int
	chunk := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text();
		if line == "" {
			sum += processGroup(chunk)
			chunk = nil
			} else {
				chunk = append(chunk, line)
			}
		}
		
		if len(chunk) > 0 {
			sum += processGroup(chunk)
		}
		fmt.Println(sum)
}

func processGroup(chunk []string) int {
	// questions := make(map[string]struct{})
	sets := make([]map[string]struct{},0)
	for _, v := range chunk {
		temp := map[string]struct{}{}
		for _, q := range v {
			temp[string(q)] = struct{}{}
		}
		sets = append(sets, temp)
	}
	return len(intersectSets(sets))
}

func intersectSets(sets []map[string]struct{}) map[string]struct{} {

	result := sets[0]

	for _, set := range sets[1:] {
			temp := map[string]struct{}{}
			for key := range result {
					if _, exists := set[key]; exists {
							temp[key] = struct{}{}
					}
			}
			result = temp
	}
	return result
}