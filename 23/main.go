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
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	cups := []int{}
	for _, l := range lines[0] {
		i, _ := strconv.Atoi(string(l))
		cups = append(cups, i)
	}
	for round := 0; round < 100; round++ {
		current := cups[0]
		next := cups[4]
		pickup := cups[1:4]
		cups = append([]int{cups[0]}, cups[4:]...)
		destinationIdx := -1
		destination := current - 1
		if destination == 0 {
			destination = 9
		}
		inPickup := inArray(destination, pickup)
		for inPickup {
			destination--
			if destination == 0 {
				destination = 9
			}
			inPickup = inArray(destination, pickup)
		}
		for i := 0; i < len(cups); i++ {
			if destination == cups[i] {
				destinationIdx = i
				break
			}
		}

		after := make([]int, len(cups[destinationIdx+1:]))
		copy(after, cups[destinationIdx+1:])
		newCups := append(cups[0:destinationIdx+1], pickup...)
		newCups = append(newCups, after...)
		cups = newCups

		nextIdx := 0
		for i, c := range cups {
			if c == next {
				nextIdx = i
				break
			}
		}

		cups = append(cups[nextIdx:], cups[:nextIdx]...)
	}

	oneIndex := 0
	for i, c := range cups {
		if c == 1 {
			oneIndex = i
			break
		}
	}
	cups = append(cups[oneIndex:], cups[:oneIndex]...)

	fmt.Printf("%s\n", arrayToString(cups[1:], ""))
}

func part2(lines []string) {
	cups := []int{}
	for _, l := range lines[0] {
		i, _ := strconv.Atoi(string(l))
		cups = append(cups, i)
	}
	for i := 10; i <= 1000000; i++ {
		cups = append(cups, i)
	}
	for round := 0; round < 10000000; round++ {
		if round%10000 == 0 {
			fmt.Println(round)
		}
		current := cups[0]
		next := cups[4]
		pickup := cups[1:4]
		cups = append([]int{cups[0]}, cups[4:]...)
		destinationIdx := -1
		destination := current - 1
		if destination == 0 {
			destination = 9
		}
		inPickup := inArray(destination, pickup)
		for inPickup {
			destination--
			if destination == 0 {
				destination = 9
			}
			inPickup = inArray(destination, pickup)
		}
		for i := 0; i < len(cups); i++ {
			if destination == cups[i] {
				destinationIdx = i
				break
			}
		}

		after := make([]int, len(cups[destinationIdx+1:]))
		copy(after, cups[destinationIdx+1:])
		newCups := append(cups[0:destinationIdx+1], pickup...)
		newCups = append(newCups, after...)
		cups = newCups

		nextIdx := 0
		for i, c := range cups {
			if c == next {
				nextIdx = i
				break
			}
		}

		cups = append(cups[nextIdx:], cups[:nextIdx]...)
	}

	oneIndex := 0
	for i, c := range cups {
		if c == 1 {
			oneIndex = i
			break
		}
	}
	cups = append(cups[oneIndex:], cups[:oneIndex]...)

	fmt.Printf("%d\n", cups[1]*cups[2])
}

func inArray(a int, b []int) bool {
	for _, i := range b {
		if a == i {
			return true
		}
	}
	return false
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
