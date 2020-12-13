package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(file), "\n")
	instructions := []instruction{}
	for _, line := range lines {
		name := line[:1]
		amount, _ := strconv.Atoi(line[1:])
		instructions = append(instructions, instruction{name: name, amount: amount})
	}
	manhattan1 := navigate(instructions)
	manhattan2 := navigateWaypoint(instructions)
	fmt.Println(manhattan1)
	fmt.Println(manhattan2)
}

func navigate(instructions []instruction) int {
	direction := 90
	x := 0
	y := 0
	for _, instruction := range instructions {
		name := instruction.name
		amount := instruction.amount
		switch name {
		case "F":
			switch direction {
			case 90:
				x += amount
			case 180:
				y -= amount
			case 270:
				x -= amount
			case 0:
				y += amount
			}
		case "R":
			direction += amount
			direction %= 360
		case "L":
			direction += 360 - amount
			direction %= 360
		case "N":
			y += amount
		case "S":
			y -= amount
		case "E":
			x += amount
		case "W":
			x -= amount
		}
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func navigateWaypoint(instructions []instruction) int {
	x := 0
	y := 0
	wx := 10
	wy := 1
	for _, instruction := range instructions {
		name := instruction.name
		amount := instruction.amount
		switch name {
		case "F":
			x += amount * wx
			y += amount * wy
		case "R":
			steps := amount / 90
			for i := 0; i < steps; i++ {
				wx, wy = wy, -wx
			}
		case "L":
			steps := amount / 90
			for i := 0; i < steps; i++ {
				wx, wy = -wy, wx
			}
		case "N":
			wy += amount
		case "S":
			wy -= amount
		case "E":
			wx += amount
		case "W":
			wx -= amount
		}
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

type instruction struct {
	name   string
	amount int
}
