package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
)

func fillWithInputNums(m map[int][]int, arr []int) {
	for i, v := range arr {
		m[v] = []int{i + 1}
	}
}

func getIntSlice(arr []string) []int {
	result := make([]int, 0, len(arr))
	for _, v := range arr {
		result = append(result, getInt(v))
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getInt(v string) int {
	result, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return result
}

func calculateNthNumber(val int, arr []int) int {
	spoken := make(map[int][]int)
	fillWithInputNums(spoken, arr)

	index := arr[len(arr) - 1]
	for i := len(arr); i < val; i++ {
		diff := 0

		if len(spoken[index]) >= 2 {
			diff = abs(spoken[index][len(spoken[index]) - 1] - spoken[index][len(spoken[index]) - 2])
		}

		spoken[diff] = append(spoken[diff], i + 1)
		index = diff
	}
	return index
}

func main() {
	data := getInput("input.txt")
	values := getIntSlice(strings.Split(data[0], ","))

	result1 := calculateNthNumber(2020, values)
	result2 := calculateNthNumber(30000000, values)

	fmt.Println(result1, result2)
}

func getInput(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	result := strings.Split(string(content), "\n")
	return result[:len(result)-1]
}
