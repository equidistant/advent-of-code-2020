package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./input.txt")
	check(err)
	lines := strings.Split(string(file), "\n")
	instructions := []instruction{}
	for i, line := range lines {
		name := line[:3]
		value, err := strconv.Atoi(strings.Replace(line[4:], " ", "", -1))
		check(err)
		instruction := instruction{name: name, value: value, done: false, index: i}
		instructions = append(instructions, instruction)
	}
	// first part
	// acc := onlyOnePass(instructions)
	// fmt.Println(acc)
	// second part

	for i := range instructions {
		tempInstructions := []instruction{}
		for _, currentInstruction := range instructions {
			tempInstructions = append(tempInstructions, instruction{name: currentInstruction.name, value: currentInstruction.value, done: false, index: currentInstruction.index})
		}
		if tempInstructions[i].name == "jmp" {
			tempInstructions[i].name = "nop"
		} else if tempInstructions[i].name == "nop" {
			tempInstructions[i].name = "jmp"
		} else {
			continue
		}
		acc, err := checkIfTerminating(tempInstructions)
		if err == nil {
			fmt.Println(acc)
			break
		}
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type instruction struct {
	name  string
	value int
	done  bool
	index int
}

func onlyOnePass(instructions []instruction) int {
	acc := 0
	currentInstruction := instructions[0]
	for currentInstruction.done == false {
		instructions[currentInstruction.index].done = true
		if currentInstruction.name == "acc" {
			acc += currentInstruction.value
		}
		if currentInstruction.name == "jmp" {
			currentInstruction = instructions[currentInstruction.index+currentInstruction.value]
		} else {
			currentInstruction = instructions[currentInstruction.index+1]
		}
	}
	return acc
}

func checkIfTerminating(instructions []instruction) (int, error) {
	acc := 0
	currentInstruction := instructions[0]

	for currentInstruction.done == false {
		instructions[currentInstruction.index].done = true
		if currentInstruction.name == "acc" {
			acc += currentInstruction.value
		}
		if currentInstruction.name == "jmp" {
			if currentInstruction.index+currentInstruction.value == len(instructions) {
				return acc, nil
			}
			currentInstruction = instructions[currentInstruction.index+currentInstruction.value]
		} else {
			if currentInstruction.index+1 == len(instructions) {
				return acc, nil
			}
			currentInstruction = instructions[currentInstruction.index+1]
		}
	}
	return 0, errors.New("indefinite loop")
}
