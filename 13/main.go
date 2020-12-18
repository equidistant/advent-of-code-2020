package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(file), "\n")
	target, _ := strconv.Atoi(lines[0])
	buses := []int{}
	for _, letter := range strings.Split(lines[1], ",") {
		if letter != "x" {
			bus, _ := strconv.Atoi(letter)
			buses = append(buses, bus)
		}
	}
	earliest := findEarliestBus(target, buses)
	chinese := findChineseBus(buses)
	fmt.Println(earliest)
	fmt.Println(chinese)
}

func findEarliestBus(target int, buses []int) int {
	minTarget := 10000000000000
	result := 0
	for _, bus := range buses {
		integer := target / bus
		integerTarget := integer * bus
		for integerTarget < target {
			integerTarget += bus
		}
		if integerTarget < minTarget {
			minTarget = integerTarget
			result = (integerTarget - target) * bus
		}
	}
	return result
}

func findChineseBus(buses []int) int {

	result := 0

	return result
}
