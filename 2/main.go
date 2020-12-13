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

func newLineSplit(c rune) bool {
	return unicode.IsSpace(c)
}

type Line struct {
	lowerBound int
	upperBound int
	letter     rune
	password   string
}

func parseFile(file []byte) []Line {
	textLines := strings.Split(string(file), "\n")
	lines := []Line{}
	for _, i := range textLines {
		textLine := strings.Fields(i)
		bounds := strings.Split(textLine[0], "-")
		lowerBound, err := strconv.Atoi(bounds[0])
		check(err)
		upperBound, err := strconv.Atoi(bounds[1])
		check(err)
		letter := []rune(textLine[1])[0]
		password := textLine[2]
		line := Line{
			lowerBound: lowerBound,
			upperBound: upperBound,
			letter:     letter,
			password:   password}
		lines = append(lines, line)
	}
	return lines
}

func checkPasswords1(lines []Line) int {
	validPasswordCount := 0
	for _, line := range lines {
		letterCount := 0
		for _, letter := range line.password {
			if letter == line.letter {
				letterCount++
			}
		}
		if letterCount >= line.lowerBound && letterCount <= line.upperBound {
			validPasswordCount++
		}
	}
	return validPasswordCount
}

func checkPasswords2(lines []Line) int {
	validPasswordCount := 0
	for _, line := range lines {
		if (rune(line.password[line.lowerBound-1]) == line.letter && rune(line.password[line.upperBound-1]) != line.letter) ||
			(rune(line.password[line.lowerBound-1]) != line.letter && rune(line.password[line.upperBound-1]) == line.letter) {
			validPasswordCount++
		}
	}
	return validPasswordCount
}

func main() {
	file, err := ioutil.ReadFile("./input.txt")
	check(err)
	lines := parseFile(file)
	validPasswordsCount1 := checkPasswords1(lines[:])
	fmt.Println(validPasswordsCount1)
	validPasswordsCount2 := checkPasswords2(lines[:])
	fmt.Println(validPasswordsCount2)
}
