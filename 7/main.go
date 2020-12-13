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
	bags := buildBagMap(lines)
	count := howMuchShinyGolds(bags)
	count2 := countBags(bags, "shiny gold", 1)
	fmt.Println(count)
	// deduct original shiny gold bag
	fmt.Println(count2 - 1)
}

func buildBagMap(lines []string) map[string][]bagInfo {
	bags := map[string][]bagInfo{}
	for _, line := range lines {
		bagLines := strings.Split(line, ",")
		var bagColor string
		var insideBagColor string
		var insideBagAmount int
		var err error
		for j, bagString := range bagLines {
			bagSplit := strings.Fields(strings.Trim(bagString, ". "))
			if j == 0 {
				bagColor = strings.Join(bagSplit[:2], " ")
				insideBagAmount, err = strconv.Atoi(bagSplit[4])
				if err != nil {
					insideBagAmount = 0
				}
				insideBagColor = strings.Join(bagSplit[5:7], " ")
			} else {
				insideBagAmount, err = strconv.Atoi(bagSplit[0])
				if err != nil {
					insideBagAmount = 0
				}
				insideBagColor = strings.Join(bagSplit[1:3], " ")
			}
			if _, exists := bags[bagColor]; !exists {
				bags[bagColor] = []bagInfo{}
			}
			if insideBagAmount != 0 {
				bags[bagColor] = append(bags[bagColor], bagInfo{color: insideBagColor, amount: insideBagAmount})
			}
		}
	}
	return bags
}

func howMuchShinyGolds(bags map[string][]bagInfo) int {
	count := 0
	for color := range bags {
		currCount := containsShinyGold(bags, color)
		if currCount > 0 {
			count++
		}
	}
	return count
}

func containsShinyGold(bags map[string][]bagInfo, color string) int {
	for _, item := range bags[color] {
		if item.color == "shiny gold" {
			return 1
		}
	}
	if len(bags[color]) == 0 {
		return 0
	}
	count := 0
	for _, item := range bags[color] {
		count += containsShinyGold(bags, item.color)
	}
	return count
}

func countBags(bags map[string][]bagInfo, color string, amount int) int {
	if len(bags[color]) == 0 {
		return amount
	}
	count := 0
	for _, item := range bags[color] {
		count += countBags(bags, item.color, item.amount)
	}
	return amount*count + amount
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type bagInfo struct {
	color  string
	amount int
}
