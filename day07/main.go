package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
)

func countBags(b map[string]map[string]int, target string) int {
	contained := map[string]bool{}
	contained[target] = true
	isFinished := false

	for !isFinished {
		isFinished = true
		for k, v := range b {
			if contained[k] {
				continue
			}

			for i := range v {
				if contained[i] {
					contained[k] = true
					isFinished = false
				}
			}
		}
	}
	return len(contained)-1
}

func sumBags(b map[string]map[string]int, t map[string]int, target string) int {
	if v, ok := t[target]; ok {
		return v
	}

	count := 0
	for k, v := range b[target] {
		count += v * (1 + sumBags(b, t, k))
	}

	t[target] = count
	return count
}

func main() {
	data, err := getInput("input.txt")
	if err != nil {
		panic(err)
	}

	bags := map[string]map[string]int{}
	parseRules(bags, data)
	result1 := countBags(bags, "shiny gold")

	// Part Two
	totalBags := map[string]int{}
	_ = sumBags(bags, totalBags, "shiny gold")

	fmt.Println(result1, totalBags["shiny gold"])
}

func parseRules(b map[string]map[string]int, data []string) {
	for _, text := range data {
		words := strings.Split(text, " ")
		parentBag := words[0] + " " + words[1]
		b[parentBag] = map[string]int{}

		bags := strings.Split(text, words[3])[1:]
		bags = strings.Split(bags[0], ", ")

		for _, bag := range bags {
			if strings.Contains(bag, "no other") {
				continue
			}

			if strings.Contains(bag, "1") {
				bag = bag[:len(bag)-4]
			} else {
				bag = bag[:len(bag)-5]
			}
			bag = strings.Trim(bag, " .")
			count, err := strconv.Atoi(string(bag[0]))
			if err != nil {
				panic(err)
			}

			b[parentBag][bag[2:]] = count
		}
	}
}

func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	result := strings.Split(string(content), "\n")
	return result[:len(result)-1], nil
}
