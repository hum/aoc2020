package main

import (
	"log"
	"io/ioutil"
	"strings"
)

func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

func countTrees(data []string, slopeX, slopeY int) int {
	result := 0

	for y, x := 0, 0;  y < len(data); y, x = y + slopeY, x + slopeX {
		location := string(data[y][x % len(data[y])])

		if location == "#" {
			result++
		}
	}
	return result
}

// Part Two
func multiplyTrees(data []string) int {
	slopes := [5][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	result := 1
	for _, slope := range slopes {
		result *= countTrees(data, slope[0], slope[1])
	}

	return result
}

func main() {
	data, err := getInput("input.txt")
	if err != nil {
		log.Fatalln("Failed to open the file: %v\n", err)
	}
	data = data[:len(data)-1]

	result := countTrees(data, 3, 1)
	log.Println(result)

	// Part Two
	result = multiplyTrees(data)
	log.Println(result)
}
