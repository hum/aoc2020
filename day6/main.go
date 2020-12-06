package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	result := strings.Split(string(content), "\n\n")
	lastElement := result[len(result)-1]
	result[len(result)-1] = lastElement[:len(lastElement)-1]

	return result, nil
}

func countGroup(data string) (int, int) {
	lines := strings.Split(data, "\n")
	letterMap := make(map[rune]int)

	for _, line := range lines {
		for _, letter := range line {
			letterMap[letter]++
		}
	}

	// Part Two
	count := 0
	for _, v := range letterMap {
		if v == len(lines) {
			count++
		}
	}
	return len(letterMap), count
}

func main() {
	data, err := getInput("input.txt")
	if err != nil {
		log.Fatalf("Failed to open the file %v\n", err)
	}

	sum1, sum2 := 0, 0
	for _, group := range data {
		result1, result2 := countGroup(group)
		sum1 += result1
		sum2 += result2
	}

	log.Println(sum1, sum2)
}
