package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func main() {
	file, err := ioutil.ReadFile("./input.txt")
	check(err)
	groups := strings.Split(string(file), "\n\n")
	// count := uniqueYesAnswers(groups)
	count := allYesAnswers(groups)
	fmt.Println(count)
}

func uniqueYesAnswers(groups []string) int {
	count := 0
	for _, group := range groups {
		answers := strings.Join(strings.FieldsFunc(group, splitByNewLineAndSpace), "")
		uniqueAnswers := map[rune]bool{}
		for _, answer := range answers {
			if !uniqueAnswers[answer] {
				uniqueAnswers[answer] = true
			}
		}
		count += len(uniqueAnswers)
	}
	return count
}

func allYesAnswers(groups []string) int {
	count := 0
	for _, group := range groups {
		splitGroup := strings.Split(group, "\n")
		fmt.Println(len(splitGroup))
		groupAnswers := map[rune]int{}
		for _, individual := range splitGroup {
			uniqueIndividualAnswers := map[rune]bool{}
			for _, answer := range individual {
				if _, exists := uniqueIndividualAnswers[answer]; !exists {
					uniqueIndividualAnswers[answer] = true
				}
			}
			for uniqueIndividualAnswer := range uniqueIndividualAnswers {
				if _, exists := groupAnswers[uniqueIndividualAnswer]; !exists {
					groupAnswers[uniqueIndividualAnswer] = 1
				} else {
					groupAnswers[uniqueIndividualAnswer]++
				}
			}
			for _, numberOfAnswers := range groupAnswers {
				fmt.Printf("Number of answers: %d", numberOfAnswers)
				fmt.Printf("- Number of people: %d", len(splitGroup))
				fmt.Println()
				if numberOfAnswers == len(splitGroup) {
					count++
				}
			}
		}

	}
	return count
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func splitByNewLineAndSpace(r rune) bool {
	return unicode.IsSpace(r)
}
