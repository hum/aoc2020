package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Instruction struct {
	direction string
	value     int
}

func tokenise(text string) Instruction {
	i := Instruction{}
	i.direction = string(text[0])

	value, err := strconv.Atoi(text[1:])
	if err != nil {
		fmt.Printf("Error converting text to int %v\n", err)
	}

	i.value = value
	return i
}

func main() {
	data, err := getInput("input.txt")
	if err != nil {
		panic(err)
	}

	instructions := make([]Instruction, 0, len(data))

	for _, v := range data {
		instr := tokenise(v)
		instructions = append(instructions, instr)
	}

	x, y, rot := 0, 0, 0
	for _, v := range instructions {
		switch v.direction {
		case "N":
			y += v.value
		case "E":
			x += v.value
		case "S":
			y -= v.value
		case "W":
			x -= v.value
		case "L":
			rot += v.value
		case "R":
			rot -= v.value
		case "F":
			angleX, angleY := rotate(x, y, rot, v.value)
			x += angleX
			y += angleY
		}
	}
	result1 := abs(x) + abs(y)

	// Part Two
	x, y, wx, wy := 0, 0, 10, 1
	for _, v := range instructions {
		switch v.direction {
		case "N":
			wy += v.value
		case "E":
			wx += v.value
		case "S":
			wy -= v.value
		case "W":
			wx -= v.value
		case "L":
			wx, wy = rotateWaypoint(wx, wy, v.value)
		case "R":
			wx, wy = rotateWaypoint(wx, wy, -v.value)
		case "F":
			x += wx * v.value
			y += wy * v.value
		}
	}
	result2 := abs(x) + abs(y)
	fmt.Println(result1, result2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func rotateWaypoint(wx, wy, value int) (int, int) {
	tmpX := int(math.Cos(radian(float64(value)))) * wx
	tmpY := int(math.Sin(radian(float64(value)))) * wy

	x := tmpX - tmpY

	tmpY = int(math.Cos(radian(float64(value)))) * wy
	tmpX = int(math.Sin(radian(float64(value)))) * wx

	y := tmpY + tmpX
	return x, y
}

func radian(degree float64) float64 {
	return degree * (math.Pi / 180)
}

func rotate(x, y, rot, value int) (int, int) {
	valueX := int(math.Cos(float64(rot)*(math.Pi/180))) * value
	valueY := int(math.Sin(float64(rot)*(math.Pi/180))) * value
	return valueX, valueY
}

func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	result := strings.Split(string(content), "\n")
	return result[:len(result)-1], nil
}
