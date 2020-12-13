package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func spaceSplit(c rune) bool {
	return unicode.IsSpace(c)
}

func stringsToInts(s []string) []int {
	result := []int{}
	for _, i := range s {
		j, err := strconv.Atoi(i)
		check(err)
		result = append(result, j)
	}
	return result
}

func twoNumbers(input []int) int {
	m := make(map[int]int)
	for _, i := range input {
		complement := 2020 - i
		j, found := m[complement]
		if found {
			return i * j
		}
		m[i] = i
	}
	return 0
}

func threeNumbers(input []int) int {
	m := make(map[int]int)
	length := len(input)

	for i := 0; i < length; i++ {
		mi := input[i]
		for j := 0; j < length; j++ {
			mj := input[j]
			complement := 2020 - mi - mj
			mk, found := m[complement]
			if found {
				return mi * mj * mk
			}
			m[mj] = mj
		}
	}
	return 0
}

func main() {
	dat, err := ioutil.ReadFile("./input.txt")
	check(err)
	input := stringsToInts(strings.FieldsFunc(string(dat), spaceSplit))
	twoNumbersProduct := twoNumbers(input[:])
	fmt.Printf("Result is %d", twoNumbersProduct)
	threeNumbersProduct := threeNumbers(input[:])
	fmt.Printf("Result is %d", threeNumbersProduct)
}
