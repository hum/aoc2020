package main

import (
	"log"
	"io/ioutil"
	"strings"
	"strconv"
)

func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n\n"), nil
}

func populate(m map[string]string, data string) {
	split := strings.Split(data, " ")

	for _, value := range split {
		if value == "" { continue }

		tmp := strings.Split(value, ":")
		m[tmp[0]] = tmp[1]
	}
}

func isValid(m map[string]string) bool {
	fields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	for _, field := range fields {
		if _, ok := m[field]; !ok {
			return false
		}
	}
	return true
}

// Part Two
func hasValidValues(m map[string]string) bool {
	validEyeColours := []string{
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth",
	}

	for k, v := range m {
		switch k {
		case "byr", "iyr", "eyr":
			tmp, err := strconv.Atoi(v)
			if err != nil {
				return false
			}

			if k == "byr" && (tmp < 1920 || tmp > 2002) {
				return false
			} else if k == "iyr" && (tmp < 2010 || tmp > 2020) {
				return false
			} else if k == "eyr" && (tmp < 2020 || tmp > 2030) {
				return false
			}
		case "hgt":
			tmp, err := strconv.Atoi(v[:len(v)-2])
			if err != nil {
				return false
			}
			unit := v[len(v)-2:]
			if unit == "in" {
				if tmp < 59 || tmp > 76 {
					return false
				}
			} else {
				if tmp < 150 || tmp > 193 {
					return false
				}
			}
		case "hcl":
			if string(v[0]) != "#" || len(v) != 7 {
				return false
			}
			_, err := strconv.ParseUint(v[1:], 16, 64)
			if err != nil {
				return false
			}
		case "ecl":
			found := false
			for _, colour := range validEyeColours {
				if colour == v {
					found = true
				}
			}
			if !found {
				return false
			}
		case "pid":
			if len(v) != 9 {
				return false
			}
			if _, err := strconv.Atoi(v); err != nil {
				return false
			}
		}
	}
	return true
}

func main() {
	data, err := getInput("input.txt")
	if err != nil {
		log.Fatalln("Failed to open the file %v\n", err)
	}

	validPassport, validData := 0, 0
	for _, text := range data {
		passport := make(map[string]string)
		text = strings.Replace(text, "\n", " ", -1)

		if populate(passport, text); isValid(passport) {
			validPassport++

			// Part Two
			if hasValidValues(passport) {
				validData++
			}
		}
	}
	log.Println(validPassport, validData)
}
