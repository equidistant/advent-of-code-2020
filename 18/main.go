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

	sum := firstPart(lines)
	sum2 := secondPart(lines)

	fmt.Println(sum)
	fmt.Println(sum2)
}

func firstPart(lines []string) int {
	sum := 0
	for _, line := range lines {
		containsParenthesis := strings.Contains(line, "(")
		// noSpaceLine := strings.ReplaceAll(line, " ", "")
		for containsParenthesis {
			start, end := depestParenthesis(line)
			result := fmt.Sprint(calc(line[start+1 : end]))

			before := make([]byte, len(line[:start]))
			after := make([]byte, len(line[end+1:]))

			middle := []byte(result)
			copy(before, line[:start])
			copy(after, line[end+1:])

			before = append(before, middle...)
			before = append(before, after...)

			newLine := string(before)
			line = newLine
			containsParenthesis = strings.Contains(line, "(")
		}
		lineResult := calc(line)
		sum += lineResult
	}
	return sum
}

func secondPart(lines []string) int {
	sum := 0
	for _, line := range lines {
		containsParenthesis := strings.Contains(line, "(")
		for containsParenthesis {
			start, end := depestParenthesis(line)
			var part = line[start+1 : end]
			if strings.Contains(part, "+") {
				part = calcAdditions(part)
			}
			if strings.Contains(part, "*") {
				part = calcMultiplications(part)
			}

			before := make([]byte, len(line[:start]))
			after := make([]byte, len(line[end+1:]))

			middle := []byte(part)
			copy(before, line[:start])
			copy(after, line[end+1:])

			before = append(before, middle...)
			before = append(before, after...)

			newLine := string(before)
			line = newLine
			containsParenthesis = strings.Contains(line, "(")
		}
		if strings.Contains(line, "+") {
			line = calcAdditions(line)
		}
		if strings.Contains(line, "*") {
			line = calcMultiplications(line)
		}
		lineResult, _ := strconv.Atoi(line)
		sum += lineResult
	}
	return sum
}

func depestParenthesis(s string) (int, int) {
	start := strings.Index(s, "(")
	end := strings.Index(s, ")")
	nextStart := strings.Index(s[start+1:], "(") + start + 1
	for end > nextStart {
		if start == nextStart {
			break
		}
		start = nextStart
		nextStart = strings.Index(s[start+1:], "(") + start + 1
	}
	return start, end
}

func calc(s string) int {
	first := true
	result := 0
	var operator string
	for _, c := range strings.Fields(s) {
		num, err := strconv.Atoi(c)
		if err == nil {
			if first {
				first = false
				result = num
			} else {
				switch operator {
				case "+":
					result += num
				case "*":
					result *= num
				}
			}
		} else {
			operator = c
		}
	}
	return result
}

func getMultiplicationIndex(elements []string) int {
	for i, elem := range elements {
		if elem == "*" {
			return i
		}
	}
	return -1
}

func getAdditionIndex(elements []string) int {
	for i, elem := range elements {
		if elem == "+" {
			return i
		}
	}
	return -1
}

func calcMultiplications(s string) string {
	elements := strings.Fields(s)
	index := getMultiplicationIndex(elements)
	for index != -1 {
		elem1, _ := strconv.Atoi(elements[index-1])
		elem2, _ := strconv.Atoi(elements[index+1])
		result := strconv.Itoa(elem1 * elem2)
		before := elements[:index-1]
		after := elements[index+2:]

		before = append(before, result)
		before = append(before, after...)

		elements = before
		index = getMultiplicationIndex(elements)
	}
	return strings.Join(elements, " ")
}

func calcAdditions(s string) string {
	elements := strings.Fields(s)
	index := getAdditionIndex(elements)
	for index != -1 {
		elem1, _ := strconv.Atoi(elements[index-1])
		elem2, _ := strconv.Atoi(elements[index+1])
		result := strconv.Itoa(elem1 + elem2)

		before := elements[:index-1]
		after := elements[index+2:]

		before = append(before, result)
		before = append(before, after...)

		elements = before
		index = getAdditionIndex(elements)
	}
	return strings.Join(elements, " ")
}
