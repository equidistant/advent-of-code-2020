package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	startingNumbers := strings.Split(string(file), ",")
	occurences := map[int][]int{}
	history := []int{}
	n := len(startingNumbers)
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(startingNumbers[0])
		history = append(history, num)
		if len(occurences[num]) == 0 {
			occurences[num] = []int{i + 1}
		}
		startingNumbers = startingNumbers[1:]
	}
	for i := n + 1; i <= 30000000; i++ {
		lastElem := history[len(history)-1]
		if len(occurences[lastElem]) == 1 {
			history = append(history, 0)
			if len(occurences[0]) == 0 {
				occurences[0] = []int{i}
			} else {
				occurences[0] = append(occurences[0], i)
			}
		} else if len(occurences[lastElem]) > 1 {
			l1 := occurences[lastElem][len(occurences[lastElem])-1]
			l2 := occurences[lastElem][len(occurences[lastElem])-2]
			nextElem := l1 - l2
			history = append(history, nextElem)
			occurences[nextElem] = append(occurences[nextElem], i)
		}
	}
	fmt.Println(history[30000000-1])

	// var lastNumber int
	// history := map[int]int{}
	// isNew := true
	// for i := 1; i <= 2020; i++ {
	// 	if len(startingNumbers) > 0 {
	// 		lastNumber, _ = strconv.Atoi(startingNumbers[0])
	// 		startingNumbers = startingNumbers[1:]
	// 		history[lastNumber] = i
	// 	} else if isNew {
	// 		lastNumber = 0
	// 		isNew = false
	// 	} else {
	// 		nextNumber := i - 1 - history[lastNumber]
	// 		history[lastNumber] = i - 1
	// 		if history[nextNumber] == 0 {
	// 			isNew = true
	// 		}
	// 		lastNumber = nextNumber
	// 	}
	// }
	// fmt.Println(lastNumber)

}
