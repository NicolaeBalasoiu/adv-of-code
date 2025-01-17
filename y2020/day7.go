package y2020

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day7() {
	file, err := os.Open("./y2020/input_day7.txt")
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

	graph := parseRules(lines)
	count := countBagsInside(graph, "shiny gold")
	fmt.Println(count)
}

func parseRules(rules []string) map[string][][2]interface{} {
	graph := make(map[string][][2]interface{})
	ruleRegex := regexp.MustCompile(`(\d+) ([a-z]+ [a-z]+) bag`)

	for _, rule := range rules {
		parts := strings.Split(rule, " bags contain ")
		container := parts[0]
		contents := parts[1]
		matches := ruleRegex.FindAllStringSubmatch(contents, -1)
		for _, match := range matches {
			count, _ := strconv.Atoi(match[1])
			color := match[2]
			graph[container] = append(graph[container], [2]interface{}{count, color})
		}
	}

	return graph
}

func countBagsThatCanContain(graph map[string][]string, target string) int{
	visited := make(map[string]bool)
	stack := []string{target}

	for len(stack) > 0 {
		current := stack[len(stack) - 1]
		stack = stack[:len(stack)-1]

		for _, parent := range graph[current] {
			if !visited[parent] {
				visited[parent] = true
				stack = append(stack, parent)
			}
		}
	}
	return len(visited)
}

func countBagsInside(graph map[string][][2]interface{}, bag string) int {
	total := 0

	for _, content := range graph[bag] {
		count := content[0].(int)
		childBag := content[1].(string)

		total += count + count*countBagsInside(graph, childBag)
	}

	return total
}