package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func evalAdapter(num1, num2 int) (int, bool) {
	if v := num2 - num1; v <= 3 {
		return v, true
	}
	return 0, false
}

func countDistinctWays(values map[int][]int, cache map[int]int, target, index int) int {
	if v, ok := cache[index]; ok {
		return v
	}

	value := 0
	for _, v := range values[index] {
		if v != target {
			value += countDistinctWays(values, cache, target, v)
			continue
		}
		value++
	}
	cache[index] = value
	return value
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

	sort.Ints(values)
	adapters := make([]int, 0, len(values)+1)
	adapters = append(adapters, 0)

	for _, v := range values {
		adapters = append(adapters, v)
	}

	deviceAdapter := values[len(values)-1] + 3
	adapters = append(adapters, deviceAdapter)

	jolts := make(map[int]int)
	for i := 0; i < len(adapters)-1; i++ {
		if diff, ok := evalAdapter(adapters[i], adapters[i+1]); ok {
			jolts[diff]++
		}
	}

	result1 := jolts[1] * jolts[3]
	// Part Two
	distinct := make(map[int][]int)
	cache := make(map[int]int)

	for _, v := range adapters[:len(adapters)-1] {
		distinct[v] = []int{v + 3, v + 2, v + 1}
	}
	result2 := countDistinctWays(distinct, cache, adapters[len(adapters)-1], 0)
	fmt.Println(result1, result2)
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
