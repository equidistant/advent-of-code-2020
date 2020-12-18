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

	cubes := map[cube]bool{}

	for y, line := range lines {
		for x, c := range line {
			if c == '.' {
				cubes[cube{x, y, 0}] = false
			} else {
				cubes[cube{x, y, 0}] = true
			}

		}
	}

	printCubes(cubes, 3, 3, 0, 0)

	for i := 0; i < 6; i++ {
		newCubes := map[cube]bool{}
		for c := range cubes {
			n := getNeighbours(c)
			for _, neighbour := range n {
				active := cubes[neighbour]
				neighbours := getNeighbours(neighbour)
				activeNeighbours := countActiveNeighbours(cubes, neighbours)
				if active {
					activeNeighbours--
				}
				if active && (activeNeighbours == 2 || activeNeighbours == 3) {
					newCubes[cube{neighbour.x, neighbour.y, neighbour.z}] = true
				} else if !active && activeNeighbours == 3 {
					newCubes[cube{neighbour.x, neighbour.y, neighbour.z}] = true
				} else {
					newCubes[cube{neighbour.x, neighbour.y, neighbour.z}] = false
				}
			}
		}
		cubes = newCubes
		// printCubes(cubes, 3+i*2, 3+i*2, 1+i, i+1)
	}
	count := 0
	for _, active := range cubes {
		if active {
			count++
		}
	}
	fmt.Println(count)

	ncubes := map[ncube]bool{}
	for y, line := range lines {
		for x, c := range line {
			if c == '.' {
				ncubes[ncube{x, y, 0, 0}] = false
			} else {
				ncubes[ncube{x, y, 0, 0}] = true
			}
		}
	}
	for i := 0; i < 6; i++ {
		fmt.Println(i)
		newNCubes := map[ncube]bool{}
		for c := range ncubes {
			n := getNNeighbours(c)
			for _, neighbour := range n {
				active := ncubes[neighbour]
				neighbours := getNNeighbours(neighbour)
				activeNeighbours := countActiveNNeighbours(ncubes, neighbours)
				if active {
					activeNeighbours--
				}
				if active && (activeNeighbours == 2 || activeNeighbours == 3) {
					newNCubes[ncube{neighbour.x, neighbour.y, neighbour.z, neighbour.w}] = true
				} else if !active && activeNeighbours == 3 {
					newNCubes[ncube{neighbour.x, neighbour.y, neighbour.z, neighbour.w}] = true
				} else {
					newNCubes[ncube{neighbour.x, neighbour.y, neighbour.z, neighbour.w}] = false
				}
			}
		}
		ncubes = newNCubes
		// printCubes(cubes, 3+i*2, 3+i*2, 1+i, i+1)
	}
	ncount := 0
	for _, active := range ncubes {
		if active {
			ncount++
		}
	}
	fmt.Println(ncount)
}

func getNeighbours(c cube) []cube {
	neighbours := []cube{}
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				neighbours = append(neighbours, cube{c.x + i, c.y + j, c.z + k})
			}
		}
	}
	return neighbours
}

func getNNeighbours(c ncube) []ncube {
	neighbours := []ncube{}
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				for l := -1; l < 2; l++ {
					neighbours = append(neighbours, ncube{c.x + i, c.y + j, c.z + k, c.w + l})
				}
			}
		}
	}
	return neighbours
}

func countActiveNeighbours(cubes map[cube]bool, neighbours []cube) int {
	n := 0
	for _, neighbour := range neighbours {
		if cubes[neighbour] {
			n++
		}
	}
	return n
}

func countActiveNNeighbours(cubes map[ncube]bool, neighbours []ncube) int {
	n := 0
	for _, neighbour := range neighbours {
		if cubes[neighbour] {
			n++
		}
	}
	return n
}

type cube struct {
	x, y, z int
}

type ncube struct {
	x, y, z, w int
}

func printCubes(cubes map[cube]bool, x int, y int, z int, n int) {
	fmt.Printf("**************** %d **************\n\n", n)
	for zdx := -z; zdx <= z; zdx++ {
		fmt.Printf("%.0f ", math.Abs(float64(zdx)))
		for i := -x; i <= x; i++ {
			fmt.Printf("%.0f ", math.Abs(float64(i)))
		}
		fmt.Println()
		for ydx := -y; ydx <= y; ydx++ {
			fmt.Printf("%.0f ", math.Abs(float64(ydx)))
			for xdx := -x; xdx <= x; xdx++ {
				if cubes[cube{xdx, ydx, zdx}] {
					fmt.Printf("# ")
				} else {
					fmt.Printf(". ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
	}
}
