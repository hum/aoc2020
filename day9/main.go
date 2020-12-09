package main

import (
	"fmt"
	"go/types"
	"io/ioutil"
	"strconv"
	"strings"
)

func isValid(target int, preamble []int) bool {
	tmp := make(map[int]types.Nil, len(preamble))
	for _, v := range preamble {
		if _, ok := tmp[target-v]; ok {
			return true
		}
		tmp[v] = types.Nil{}
	}
	return false
}

func findSum(data []int, target int) int {
	for i, _ := range data {
		min, max, sum := data[i], data[i], data[i]

		for j := i + 1; j < len(data); j++ {
			if sum == target {
				return min + max
			}

			if data[j] > max {
				max = data[j]
			} else if data[j] < min {
				min = data[j]
			}
			sum += data[j]
		}
	}
	return 0
}

func main() {
	data, err := getInput("input.txt")
	if err != nil {
		fmt.Printf("Failed to open the file %v\n", err)
	}
	values, err := sliceAtoi(data)
	if err != nil {
		fmt.Printf("Failed to convert values to int. %v\n", err)
	}

	for i, v := range values[25:] {
		preamble := values[i : i+25]

		if !isValid(v, preamble) {
			result2 := findSum(values, v)
			fmt.Println(v, result2)
			break
		}
	}
}

func sliceAtoi(data []string) ([]int, error) {
	result := make([]int, 0, len(data))

	for _, v := range data {
		num, err := strconv.Atoi(v)
		if err != nil {
			return result, err
		}
		result = append(result, num)
	}
	return result, nil
}

func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	result := strings.Split(string(content), "\n")
	return result[:len(result)-1], nil
}
