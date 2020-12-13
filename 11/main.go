package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	originalSeats := strings.Split(string(file), "\n")
	newSeats := originalSeats[:]
	previousOccupied := countOccupied(originalSeats)
	occupied := 0

	for {
		newSeats = seatingRound2(newSeats)
		// for _, row := range newSeats {
		// 	fmt.Println(row)
		// }
		// fmt.Println()
		occupied = countOccupied(newSeats)
		// fmt.Printf("Occupied seats: %d\n", occupied)
		// fmt.Println()
		// fmt.Println()
		if previousOccupied == occupied {
			break
		}
		previousOccupied = occupied
	}
	fmt.Printf("%d", occupied)
}

func seatingRound(seats []string) []string {
	x := len(seats)
	y := len(seats[0])
	newSeats := []string{}
	for i := 0; i < x; i++ {
		var newRow string
		for j := 0; j < y; j++ {
			if seats[i][j] == '.' {
				newRow += "."
				continue
			}
			occupied := 0
			row := i - 1
			col := j - 1
			if isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
				}
			}
			row = i - 1
			col = j
			if isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
				}
			}
			row = i - 1
			col = j + 1
			if isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
				}
			}
			row = i
			col = j + 1
			if isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
				}
			}
			row = i + 1
			col = j + 1
			if isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
				}
			}
			row = i + 1
			col = j
			if isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
				}
			}
			row = i + 1
			col = j - 1
			if isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
				}
			}
			row = i
			col = j - 1
			if isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
				}
			}
			if occupied >= 4 {
				newRow += "L"
			} else if occupied == 0 {
				newRow += "#"
			} else {
				newRow += string(seats[i][j])
			}
		}
		newSeats = append(newSeats, newRow)
	}
	seats = newSeats
	return newSeats
}

func seatingRound2(seats []string) []string {
	x := len(seats)
	y := len(seats[0])
	newSeats := []string{}
	for i := 0; i < x; i++ {
		var newRow string
		for j := 0; j < y; j++ {
			if seats[i][j] == '.' {
				newRow += "."
				continue
			}
			occupied := 0
			row := i - 1
			col := j - 1
			for isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
					break
				} else if seats[row][col] == 'L' {
					break
				}
				row--
				col--
			}
			row = i - 1
			col = j
			for isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
					break
				} else if seats[row][col] == 'L' {
					break
				}
				row--
			}
			row = i - 1
			col = j + 1
			for isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
					break
				} else if seats[row][col] == 'L' {
					break
				}
				row--
				col++
			}
			row = i
			col = j + 1
			for isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
					break
				} else if seats[row][col] == 'L' {
					break
				}
				col++
			}
			row = i + 1
			col = j + 1
			for isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
					break
				} else if seats[row][col] == 'L' {
					break
				}
				row++
				col++
			}
			row = i + 1
			col = j
			for isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
					break
				} else if seats[row][col] == 'L' {
					break
				}
				row++
			}
			row = i + 1
			col = j - 1
			for isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
					break
				} else if seats[row][col] == 'L' {
					break
				}
				row++
				col--
			}
			row = i
			col = j - 1
			for isInBounds(row, col, x, y) {
				if seats[row][col] == '#' {
					occupied++
					break
				} else if seats[row][col] == 'L' {
					break
				}
				col--
			}
			if occupied >= 5 {
				newRow += "L"
			} else if occupied == 0 {
				newRow += "#"
			} else {
				newRow += string(seats[i][j])
			}
		}
		newSeats = append(newSeats, newRow)
	}
	seats = newSeats
	return newSeats
}

func countOccupied(seats []string) int {
	x := len(seats)
	y := len(seats[0])
	occupied := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if seats[i][j] == '#' {
				occupied++
			}
		}
	}
	return occupied
}

func isInBounds(i int, j int, x int, y int) bool {
	return i >= 0 && j >= 0 && i < x && j < y
}
