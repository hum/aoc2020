package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	MaxInt = 2147483647
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getInt(s string) (int, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return -1, err
	}
	return v, nil
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	data, err := getInput("input.txt")
	if err != nil {
		panic(err)
	}

	timestamp, _ := getInt(data[0])
	busIDs := make([]int, 0)
	minutes := make([]int, 0)

	for i, v := range strings.Split(data[1], ",") {
		if v != "x" {
			id, err := getInt(v)
			if err != nil {
				panic(err)
			}
			busIDs = append(busIDs, id)
			minutes = append(minutes, i)
		}
	}

	earliestDeparture := MaxInt
	result := 0

	for _, v := range busIDs {
		tmp := 0
		for tmp <= timestamp {
			tmp += v
		}

		if tmp < earliestDeparture {
			result = v
			earliestDeparture = tmp
		}
	}
	result1 := abs(timestamp - earliestDeparture) * result

	// Part Two
	departure := busIDs[0]
	sum := busIDs[0] + minutes[0]

	for i, v := range busIDs {
		for (sum + minutes[i]) % v != 0 {
			sum += departure
		}
		departure = (departure * v) / gcd(departure, v)
	}

	fmt.Println(result1, sum)
}

func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	result := strings.Split(string(content), "\n")
	return result[:len(result)-1], nil
}
