package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"go/types"
)

const (
	ACC = "acc"
	JMP = "jmp"
	NOP = "nop"
)

type InstructionType string

type Instruction struct {
	Type InstructionType
	Value int
	literal string
}

func (i *Instruction) tokenize() error {
	tmp := strings.Split(i.literal, " ")

	switch tmp[0] {
	case ACC:
		i.Type = ACC
	case JMP:
		i.Type = JMP
	case NOP:
		i.Type = NOP
	}

	v, err := strconv.Atoi(tmp[1])
	if err != nil {
		return err
	}

	i.Value = v
	return nil
}

func parse(value string) (Instruction, error) {
	i := Instruction{literal: value}
	err := i.tokenize()
	if err != nil {
		return Instruction{}, err
	}
	return i, nil
}

func eval(tokens []Instruction) (int, error) {
	values := make(map[int]types.Nil)
	accumulator, pointer := 0, 0

	for {
		if _, ok := values[pointer]; ok {
			return accumulator, fmt.Errorf("Infinite loop.")
		}

		values[pointer] = types.Nil{}

		switch tokens[pointer].Type {
		case ACC:
			accumulator += tokens[pointer].Value
			pointer++
		case JMP:
			pointer += tokens[pointer].Value
		case NOP:
			pointer++
		}

		if pointer >= len(tokens) {
			return accumulator, nil
		}
	}
	return 0, nil
}

func swapInstruction(target int, tokens []Instruction) []Instruction {
	result := make([]Instruction, len(tokens))
	tmp := 0

	copy(result, tokens)

	for i := 0; i < len(tokens); i ++ {
		if tokens[i].Type == JMP || tokens[i].Type == NOP {
			if tmp == target {
				if tokens[i].Type == JMP {
					result[i].Type = swapValues(tokens[i].Type)
					return result
				}
			}
			tmp++
		}
	}
	return result
}


func swapValues(value InstructionType) InstructionType {
	if value == JMP {
		return NOP
	} else if value == NOP {
		return JMP
	}
	return value
}

func evalHalt(tokens []Instruction) int {
	tries := 0

	for {
		result, err := eval(swapInstruction(tries, tokens))
		if err != nil {
			tries++
		} else {
			return result
		}
	}
}

func main() {
	data, err := getInput("input.txt")
	if err != nil {
		fmt.Printf("Failed to open the file %v\n", err)
	}

	var tokens []Instruction

	for _, v := range data {
		token, err := parse(v)
		if err != nil {
			panic(err)
		}
		tokens = append(tokens, token)
	}

	result, _ := eval(tokens)
	result2 := evalHalt(tokens)
	fmt.Println(result, result2)
}

func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	result := strings.Split(string(content), "\n")
	return result[:len(result)-1], nil
}
