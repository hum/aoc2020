package main

import (
	"log"
	"os"
	"bufio"
	"strconv"
	"io"
	"sort"
	"fmt"
)

func readIntsFromFile(r io.Reader) ([]int, error) {
	var result []int
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, value)
	}
	return result, nil
}

func findTwoSum(numbers []int, target int) (int, int) {
	result := make(map[int]int)

	for _, v := range numbers {
		result[target - v] = v
		if value, ok := result[v]; ok {
			return value, v
		}
	}
	return 0, 0
}

// Part Two
func findThreeSum(nums []int, target int) (int, int, int) {
	for i, _ := range nums[:len(nums) - 2] {
		start, end := i + 1, len(nums) - 1

		for start < end {
			sum := nums[i] + nums[start] + nums[end]

			if sum > target {
				end--
			} else if sum < target {
				start++
			} else {
				return nums[i], nums[start], nums[end]
			}
		}
	}
	return 0, 0, 0
}

func main() {
	d, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open the file: %v\n", err)
	}

	data, err := readIntsFromFile(d)
	if err != nil {
		log.Fatalf("Failed to convert file contents to int array: %v\n")
	}

	sort.Ints(data)
	value1, value2 := findTwoSum(data, 2020)
	fmt.Println("Part One result: ", value1 * value2)

	// Part Two
	value1, value2, value3 := findThreeSum(data, 2020)
	fmt.Println("Part Two result: ", value1 * value2 * value3)
}
