package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

func parseLine(line string) (int, int, string, string) {
	substrings := strings.Split(line, " ")

	limits := strings.Split(substrings[0], "-")
	low, _ := strconv.Atoi(limits[0])
	high, _ := strconv.Atoi(limits[1])

	target := string(substrings[1][0])
	password := substrings[2]

	return low, high, target, password
}

func countAppearance(target string, text string) int {
	count := 0

	for _, ch := range text {
		if string(ch) == target {
			count++
		}
	}
	return count
}

func validateCount(low int, high int, count int) bool {
	return count >= low && count <= high
}

// Part Two
func checkPositionalAppearance(pos1 int, pos2 int, target string, text string) bool {
	check1 := false
	check2 := false

	if string(text[pos1-1]) == target {	check1 = true	}
	if string(text[pos2-1]) == target {	check2 = true	}

	if check1 && !check2 {
		return true
	} else if check2 && !check1 {
		return true
	}
	return false
}

func main() {
	data, err := getInput("input.txt")
	if err != nil {
		log.Fatalf("Failed to open the file: %v\n", err)
	}

	validCount := 0
	validPosition := 0

	for _, line := range data[:len(data)-1] {
		low, high, target, password := parseLine(line)
		count := countAppearance(target, password)

		if isValid := validateCount(low, high, count); isValid {
			validCount++
		}

		// Part Two
		position1, position2 := low, high
		isValid := checkPositionalAppearance(position1, position2, target, password)
		if isValid {
			validPosition++
		}
	}

	log.Println(validCount, validPosition)
}
