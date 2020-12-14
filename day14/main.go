package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
)

func replaceValue(text, replace string, index int) string {
	result := ""
	for i := 0; i < len(text); i++ {
		if i == index {
			result += replace
			continue
		}
		result += string(text[i])
	}
	return result
}

func addPadding(binary string, pad int) string {
	for len(binary) < pad {
		binary = "0" + binary
	}
	return binary
}

func parseMem(line string) (int, int) {
	line = strings.Replace(line, "mem[", "", 1)
	line = strings.Replace(line, "]", "", 1)
	result := strings.Split(line, " = ")

	addr, err := strconv.Atoi(result[0])
	if err != nil {
		panic(err)
	}

	value, err := strconv.Atoi(result[1])
	if err != nil {
		panic(err)
	}
	return addr, value
}

func perm(s1 string, s2 []string) []string {
	result := make([]string, 0)

	for _, v := range s2 {
		tmp, j := s1, 0
		for i := 0; i < len(tmp); i++ {
			if string(tmp[i]) == "X" {
				tmp = replaceValue(tmp, string(v[j]), i)
				j++
			}
		}
		result = append(result, tmp)
	}
	return result
}

func getAll(length int) []string {
	result := make([]string, 0)

	val := ""
	for i := 0; i < length; i++ {
		val += "1"
	}

	bin := ""
	for i := 0; bin != val; i++ {
		bin = strconv.FormatInt(int64(i), 2)
		result = append(result, addPadding(bin, length))
	}
	return result
}

func main() {
	data := getInput("input.txt")
	memory, mask := make(map[int64]int), ""

	for _, v := range data {
		if strings.HasPrefix(v, "mask") {
			mask = strings.Split(v, "mask = ")[1]
			continue
		}

		addr, value := parseMem(v)
		binaryVal := addPadding(strconv.FormatInt(int64(value), 2), 36)

		for i, v := range mask {
			if string(v) != "X" {
				binaryVal = replaceValue(binaryVal, string(v), i)
			}
		}

		val, err := strconv.ParseInt(binaryVal, 2, 64)
		if err != nil {
			panic(err)
		}

		memory[int64(addr)] = int(val)
	}

	result1 := 0
	for _, v := range memory {
		result1 += v
	}

	// Part Two
	memPerm, mask := make(map[int64]int), ""
	for _, v := range data {
		if strings.HasPrefix(v, "mask") {
			mask = strings.Split(v, "mask = ")[1]
			continue
		}

		addr, value := parseMem(v)
		binaryVal := addPadding(strconv.FormatInt(int64(addr), 2), 36)

		for i, v := range mask {
			if string(v) != "0" {
				binaryVal = replaceValue(binaryVal, string(v), i)
			}
		}

		bins := getAll(strings.Count(mask, "X"))
		for _, v := range perm(binaryVal, bins) {
			memAddr, err := strconv.ParseInt(v, 2, 64)
			if err != nil {
				panic(err)
			}

			memPerm[memAddr] = value
		}
	}
	result2 := 0
	for _, v := range memPerm {
		result2 += v
	}

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
