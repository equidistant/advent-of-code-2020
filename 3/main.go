package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./input.txt")
	check(err)
	lines := strings.Split(string(file), "\n")
	treeCount1 := countTrees(1, 1, lines)
	treeCount2 := countTrees(3, 1, lines)
	treeCount3 := countTrees(5, 1, lines)
	treeCount4 := countTrees(7, 1, lines)
	treeCount5 := countTrees(1, 2, lines)
	product := treeCount1 * treeCount2 * treeCount3 * treeCount4 * treeCount5
	fmt.Println(product)
}

func countTrees(right int, down int, lines []string) int {
	modulo := len(lines[0])
	currentIndex := 0
	treeCount := 0
	for i := 0; i < len(lines); i += down {
		if lines[i][currentIndex] == '#' {
			treeCount++
		}
		currentIndex = (currentIndex + right) % modulo
	}
	return treeCount
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
