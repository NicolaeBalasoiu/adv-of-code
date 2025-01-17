package y2020

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day4() {
	file, err := os.Open("./y2020/input_day4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	validCount := 0
	var chunk []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			if len(chunk) > 0 {
				if isValidPassport(parsePassport(chunk)) {
					validCount++
				}
				chunk = nil
			}
		} else {
			chunk = append(chunk, line)
		}
	}

	// Process the last chunk
	if len(chunk) > 0 {
		if isValidPassport(parsePassport(chunk)) {
			validCount++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Valid passports:", validCount)
}

// parsePassport combines the lines of a chunk into a single map
func parsePassport(chunk []string) map[string]string {
	passport := make(map[string]string)
	combined := strings.Join(chunk, " ") // Combine all lines into one string
	fields := strings.Fields(combined)  // Split into key-value pairs

	for _, field := range fields {
		parts := strings.Split(field, ":")
		passport[parts[0]] = parts[1]
	}

	return passport
}

// isValidPassport checks if a passport has all required fields and valid data
func isValidPassport(passport map[string]string) bool {
	validators := map[string]func(string) bool{
		"byr": validateBYR,
		"iyr": validateIYR,
		"eyr": validateEYR,
		"hgt": validateHGT,
		"hcl": validateHCL,
		"ecl": validateECL,
		"pid": validatePID,
	}

	// Check required fields
	for field, validate := range validators {
		value, ok := passport[field]
		if !ok || !validate(value) {
			return false
		}
	}

	return true
}

// Validation Functions

func validateBYR(value string) bool {
	year, err := strconv.Atoi(value)
	return err == nil && year >= 1920 && year <= 2002
}

func validateIYR(value string) bool {
	year, err := strconv.Atoi(value)
	return err == nil && year >= 2010 && year <= 2020
}

func validateEYR(value string) bool {
	year, err := strconv.Atoi(value)
	return err == nil && year >= 2020 && year <= 2030
}

func validateHGT(value string) bool {
	if strings.HasSuffix(value, "cm") {
		height, err := strconv.Atoi(strings.TrimSuffix(value, "cm"))
		return err == nil && height >= 150 && height <= 193
	} else if strings.HasSuffix(value, "in") {
		height, err := strconv.Atoi(strings.TrimSuffix(value, "in"))
		return err == nil && height >= 59 && height <= 76
	}
	return false
}

func validateHCL(value string) bool {
	hclRegex := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	return hclRegex.MatchString(value)
}

func validateECL(value string) bool {
	eclRegex := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	return eclRegex.MatchString(value)
}

func validatePID(value string) bool {
	pidRegex := regexp.MustCompile(`^\d{9}$`)
	return pidRegex.MatchString(value)
}