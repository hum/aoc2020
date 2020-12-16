package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func convRulesToIntArr(data []string) []int {
	result := make([]int, 0, 4)
	for _, v := range data {
		v := strings.Split(v, "-")
		result = append(result, toInt(v[0]))
		result = append(result, toInt(v[1]))
	}
	return result
}

func toInt(v string) int {
	result, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return result
}

func isValid(rules map[string][]int, value int) bool {
	for _, v := range rules {
		if value >= v[0] && value <= v[1] || value >= v[2] && value <= v[3] {
			return true
		}
	}
	return false
}

func parseTickets(m map[int][]int, data string) {
	for i, v := range strings.Split(data, "\n")[1:] {
		result := make([]int, 0)
		split := strings.Split(v, ",")

		for _, num := range split {
			result = append(result, toInt(num))
		}
		m[i] = result
	}
}

func parseRules(m map[string][]int, data string) {
	for _, v := range strings.Split(data, "\n") {
		split := strings.Split(v, ":")
		rule := split[0]

		numberSlice := strings.Split(strings.Trim(split[1], " "), " or ")
		m[rule] = convRulesToIntArr(numberSlice)
	}
}

func main() {
	data := getInput("input.txt")

	rules := make(map[string][]int)
	parseRules(rules, data[0])

	tickets := make(map[int][]int)
	parseTickets(tickets, data[2])

	result1 := 0

	for _, v := range tickets {
		for _, field := range v {
			if !isValid(rules, field) {
				result1 += field
			}
		}
	}
	fmt.Println(result1)
}
func getInput(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	result := strings.Split(string(content), "\n\n")
	lastElement := result[len(result)-1]
	result[len(result)-1] = lastElement[:len(lastElement)-1]

	return result
}
