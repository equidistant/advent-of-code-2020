package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(file), "\n")
	tiles := part1(lines)
	part2(tiles)
}

func part1(lines []string) map[coordinate]bool {
	tiles := map[coordinate]bool{}
	for _, line := range lines {
		x := 0
		y := 0
		var previous rune
		for _, c := range line {
			// se
			if c == 'e' && previous == 's' {
				x++
				y--
			} else if c == 'w' && previous == 's' {
				x--
				y--
			} else if c == 'e' && previous == 'n' {
				x++
				y++
			} else if c == 'w' && previous == 'n' {
				x--
				y++
			} else if c == 'w' {
				x -= 2
			} else if c == 'e' {
				x += 2
			}
			previous = c
		}
		tiles[coordinate{x, y}] = !tiles[coordinate{x, y}]
	}
	count := 0
	for _, black := range tiles {
		if black {
			count++
		}
	}
	// fmt.Println(count)

	return tiles
}

func part2(tiles map[coordinate]bool) {
	for i := 0; i < 100; i++ {
		//printTiles(tiles, -5, 5)
		nextTiles := map[coordinate]bool{}
		for coordinate, black := range tiles {
			neighbours := getNeighbours(coordinate)
			n := countBlackNeighbours(coordinate, tiles)
			if black && (n == 0 || n > 2) {
				nextTiles[coordinate] = false
			} else if !black && n == 2 {
				nextTiles[coordinate] = true
			} else {
				nextTiles[coordinate] = tiles[coordinate]
			}
			for _, nCoordinate := range neighbours {
				black := tiles[nCoordinate]
				n := countBlackNeighbours(nCoordinate, tiles)
				if black && (n == 0 || n > 2) {
					nextTiles[nCoordinate] = false
				} else if !black && n == 2 {
					nextTiles[nCoordinate] = true
				} else {
					nextTiles[nCoordinate] = tiles[nCoordinate]
				}
			}
		}
		tiles = nextTiles
		count := 0
		for _, black := range tiles {
			if black {
				count++
			}
		}
		fmt.Println(count)
	}
	// count := 0
	// for _, black := range tiles {
	// 	if black {
	// 		count++
	// 	}
	// }
	// fmt.Println(count)
}

func countBlackNeighbours(c coordinate, tiles map[coordinate]bool) int {
	count := 0
	if tiles[coordinate{c.x + 1, c.y + 1}] {
		count++
	}
	if tiles[coordinate{c.x + 2, c.y}] {
		count++
	}
	if tiles[coordinate{c.x + 1, c.y - 1}] {
		count++
	}
	if tiles[coordinate{c.x - 1, c.y - 1}] {
		count++
	}
	if tiles[coordinate{c.x - 2, c.y}] {
		count++
	}
	if tiles[coordinate{c.x - 1, c.y + 1}] {
		count++
	}
	return count
}

func getNeighbours(c coordinate) []coordinate {
	return []coordinate{coordinate{c.x + 1, c.y + 1}, coordinate{c.x + 2, c.y}, coordinate{c.x + 1, c.y - 1}, coordinate{c.x - 1, c.y - 1}, coordinate{c.x - 2, c.y}, coordinate{c.x - 1, c.y + 1}}
}

type coordinate struct {
	x int
	y int
}

func printTiles(tiles map[coordinate]bool, start int, end int) {
	for j := start; j <= end; j++ {
		if j == start {
			fmt.Printf("  ")
			for k := start; k <= end; k++ {
				fmt.Printf("% 0.f", math.Abs(float64(k)))
			}
			fmt.Println()
		}
		for i := start; i <= end; i++ {
			if i == start {
				fmt.Printf(" %0.f", math.Abs(float64(j)))
			}
			if tiles[coordinate{i, j}] {
				fmt.Printf(" #")
			} else {
				fmt.Printf(" .")
			}
		}
		fmt.Println()
		fmt.Println()
	}
	fmt.Println()
}
