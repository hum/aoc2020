package main

import (
	"strings"
	"io/ioutil"
	"log"
)

func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

func binarySearch(low int, high int, data string, target string) int {
	result := 0
	for _, char := range data {
		if string(char) == target {
			low, result = low + ((high - low) / 2 + 1), high
		} else {
			high, result = high - ((high - low) / 2 + 1), low
		}
	}
	return result
}

func main() {
	data, err := getInput("input.txt")
	if err != nil {
		log.Fatalf("Failed to open the file %v\n", err)
	}

	data = data[:len(data)-1]
	leftSeats := make(map[int]string)
	highest, ourSeat := 0, 0

	for _, seat := range data {
		row := binarySearch(0, 127, seat[:7], "B")
		column := binarySearch(0, 7, seat[7:], "R")
		result := (row * 8) + column

		if result > highest {
			highest = result
		}
		leftSeats[result] = ""
	}

	// Part Two
	for i := 7; i < len(data); i++ {
		if _, ok := leftSeats[i]; !ok {
			ourSeat = i
		}
	}
	log.Println(highest, ourSeat)
}
