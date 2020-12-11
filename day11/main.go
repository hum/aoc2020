package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	EMPTY    = "L"
	OCCUPIED = "#"
	FLOOR    = "."
)

var adjacentIndex = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func updateChairValue(x, y int, data [][]string, considerFirstChair bool) string {
	if data[y][x] == FLOOR {
		return FLOOR
	}

	if !considerFirstChair {
		nearby := 0
		for _, v := range adjacentIndex {
			adjX, adjY := x+v[0], y+v[1]
			if (adjX >= 0 && adjY >= 0) && (adjY < len(data) && adjX < len(data[adjY])) {
				if data[adjY][adjX] == OCCUPIED {
					nearby++
				}
			}
		}

		if data[y][x] == EMPTY && nearby == 0 {
			return OCCUPIED
		} else if data[y][x] == OCCUPIED && nearby >= 4 {
			return EMPTY
		}
		return data[y][x]
	} else {
		first := 0
		for _, v := range adjacentIndex {
			if firstChairValid(x, y, v[0], v[1], data) {
				first++
			}
		}

		if data[y][x] == EMPTY && first == 0 {
			return OCCUPIED
		} else if data[y][x] == OCCUPIED && first >= 5 {
			return EMPTY
		}
		return data[y][x]
	}
}

func firstChairValid(x, y, adjX, adjY int, data [][]string) bool {
	for {
		x += adjX
		y += adjY

		if (y >= len(data) || y < 0) || (x >= len(data[y]) || x < 0) {
			return false
		}
		if data[y][x] == EMPTY {
			return false
		}
		if data[y][x] == OCCUPIED {
			return true
		}
	}
}

func main() {
	data, err := getInput("input.txt")
	if err != nil {
		panic(err)
	}

	values := getGrid(data)
	for {
		unchanged := true
		grid := copyGrid(values)

		for y := 0; y < len(values); y++ {
			for x := 0; x < len(values[y]); x++ {
				result := updateChairValue(x, y, values, false)
				if values[y][x] != result {
					unchanged = false
				}
				grid[y][x] = result
			}
		}
		if !unchanged {
			values = grid
		} else {
			break
		}
	}

	result1 := countOccupied(values)
	// Part Two
	values = getGrid(data)
	for {
		unchanged := true
		grid := copyGrid(values)

		for y := 0; y < len(values); y++ {
			for x := 0; x < len(values[y]); x++ {
				result := updateChairValue(x, y, values, true)
				if values[y][x] != result {
					unchanged = false
				}
				grid[y][x] = result
			}
		}
		if !unchanged {
			values = grid
		} else {
			break
		}
	}

	result2 := countOccupied(values)
	fmt.Println(result1, result2)
}

func countOccupied(data [][]string) (result int) {
	for _, row := range data {
		for _, v := range row {
			if v == OCCUPIED {
				result++
			}
		}
	}
	return result
}

func copyGrid(data [][]string) [][]string {
	result := make([][]string, 0)

	for _, v := range data {
		row := make([]string, len(v))
		copy(row, v)
		result = append(result, row)
	}
	return result
}

func getGrid(data []string) [][]string {
	result := make([][]string, len(data))

	for _, row := range data {
		line := strings.Split(row, "")
		result = append(result, line)
	}
	return result
}

func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	result := strings.Split(string(content), "\n")
	return result[:len(result)-1], nil
}
