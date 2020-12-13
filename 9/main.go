package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./input.txt")
	check(err)
	lines := strings.Split(string(file), "\n")
	buffer := []int{}
	for _, line := range lines[:25] {
		num, err := strconv.Atoi(line)
		check(err)
		buffer = append(buffer, num)
	}
	linesWithoutBuffer := lines[25:]
	noPair := findOneWithNoPair(linesWithoutBuffer[:], buffer[:])
	continuous := findContinuousSet(lines[:], noPair)
	fmt.Println(continuous)
}

func findOneWithNoPair(lines []string, buffer []int) int {
	noPair := 0
	for i := 0; i < len(lines); i++ {
		target, err := strconv.Atoi(lines[i])
		check(err)
		found := false
		for j := 0; j < 24; j++ {
			for k := j + 1; k < 25; k++ {
				if buffer[j]+buffer[k] == target {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			noPair = target
			break
		}
		buffer = append(buffer[1:], target)
	}
	return noPair
}

func findContinuousSet(lines []string, target int) int {
	sum := 0
	set := []int{}
	for i := 0; i < len(lines)-1; i++ {
		sum = 0
		set = set[:0]
		elemI, err := strconv.Atoi(lines[i])
		check(err)
		sum += elemI
		set = append(set, elemI)
		for j := i + 1; j < len(lines); j++ {
			elemJ, err := strconv.Atoi(lines[j])
			check(err)
			sum += elemJ
			set = append(set, elemJ)
			if sum > target {
				break
			}
			if sum == target {
				fmt.Println(target)
				min := set[0]
				max := set[0]
				for _, elem := range set {
					if elem > max {
						max = elem
					} else if elem < min {
						min = elem
					}
				}
				fmt.Println(target)
				fmt.Println(set)
				fmt.Println(min)
				fmt.Println(max)
				return min + max
			}
		}
	}
	return 0
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
