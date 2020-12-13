package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./input.txt")
	check(err)
	seats := strings.Split(string(file), "\n")
	max := 0
	ids := map[int]bool{}
	for _, seat := range seats {
		rowFrom := 0
		rowTo := 127
		for _, r := range seat[:len(seat)-3] {
			if r == 'F' {
				rowTo = ((rowFrom + rowTo + 1) / 2) - 1
			} else if r == 'B' {
				rowFrom = (rowFrom + rowTo + 1) / 2
			}
		}
		columnFrom := 0
		columnTo := 7
		for _, c := range seat[len(seat)-3:] {
			if c == 'L' {
				columnTo = ((columnFrom + columnTo + 1) / 2) - 1
			} else if c == 'R' {
				columnFrom = (columnFrom + columnTo + 1) / 2
			}
		}
		id := rowFrom*8 + columnFrom
		ids[id] = true
		if id > max {
			max = rowFrom*8 + columnFrom
		}
	}
	keys := make([]int, 0, len(ids))
	for k := range ids {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)
	targetSeat := -1
	previous := keys[0]
	for i := 1; i < len(keys)-1; i++ {
		if previous+2 == keys[i] {
			targetSeat = previous + 1
		}
		previous = keys[i]
	}
	fmt.Printf("Max: %d", max)
	fmt.Printf("Target seat: %d", targetSeat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
