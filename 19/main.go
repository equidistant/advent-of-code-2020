package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	parts := strings.Split(string(file), "\n\n")
	rules := parseRules(strings.Split(parts[0], "\n"))
	// regex := buildRegex(rules, "0")
	regex1 := buildRegex(rules, "42")
	regex2 := buildRegex(rules, "31")
	messages := strings.Split(parts[1], "\n")
	// count := 0
	// for _, m := range messages {
	// 	for _, r := range regex {
	// 		if r == m {
	// 			count++
	// 		}
	// 	}
	// }
	// for each message
	count := 0
	for _, m := range messages {
		l1 := len(regex1[0])
		// for each slice
		idx := 0
		count1 := 0
		msgLen := len(m)
		for i := 0; i < msgLen; i += l1 {
			valid := false
			c := m[i : i+l1]
			for _, r := range regex1 {
				if r == c {
					count1++
					valid = true
					break
				}
			}
			if !valid {
				break
			}
			idx += l1
		}
		// aaaaaaaa babbabab abaaaaba baaaabba
		l2 := len(regex2[0])
		count2 := 0
		for i := idx; i < len(m); i += l2 {
			valid := false
			c := m[i : i+l2]
			for _, r := range regex2 {
				if r == c {
					count2++
					valid = true
					break
				}
			}
			if !valid {
				break
			}
		}
		if count1 > count2 {
			count++
			fmt.Println(m)
		}
	}

	fmt.Println(count)

}

func parseRules(lines []string) map[string]string {
	rules := map[string]string{}
	for _, line := range lines {
		split := strings.Split(line, ":")
		ruleValue := strings.Trim(split[1], " ")
		rules[split[0]] = ruleValue
	}
	return rules
}

func buildRegex(rules map[string]string, rule string) []string {

	if strings.Contains(rule, "\"") {
		return []string{string(rule[1])}
	}

	if !strings.Contains(rule, " ") {
		return buildRegex(rules, rules[rule])
	}

	if strings.Contains(rule, "|") {

		sides := strings.Split(rule, "|")
		left := buildRegex(rules, strings.Trim(sides[0], " "))
		right := buildRegex(rules, strings.Trim(sides[1], " "))

		// for _, l := range left {
		// 	for _, r := range right {
		// 		result = append(result, strings.Join([]string{l, r}, ""))
		// 	}
		// }

		return append(left, right...)
	}
	elements := strings.Split(strings.Trim(rule, " "), " ")
	result := []string{}
	leftover := [][]string{}
	for _, e := range elements {
		leftover = append(leftover, buildRegex(rules, e))
	}

	for _, elem1 := range leftover[0] {
		for _, elem2 := range leftover[1] {
			result = append(result, strings.Join([]string{elem1, elem2}, ""))
		}
	}

	for _, l := range leftover[2:] {
		newResult := []string{}
		for _, elem2 := range l {
			for _, elem1 := range result {
				newResult = append(newResult, strings.Join([]string{elem1, elem2}, ""))
			}
		}
		result = newResult
	}

	return result

}
