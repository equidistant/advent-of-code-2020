package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(file), "\n")
	adapters := []int{0}
	for _, line := range lines {
		line, _ := strconv.Atoi(line)
		adapters = append(adapters, line)
	}
	quicksort(adapters, 0, len(adapters)-1)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	fmt.Println(adapters)
	dif1 := 0
	dif3 := 0
	for i := 1; i < len(adapters); i++ {
		switch dif := adapters[i] - adapters[i-1]; dif {
		case 1:
			dif1++
		case 3:
			dif3++
		}
	}
	fmt.Println(dif1 * dif3)
	permutations := allPermutations(adapters)
	fmt.Printf("%f", permutations)

}

func quicksort(arr []int, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			temp := arr[splitIndex]

			arr[splitIndex] = arr[i]
			arr[i] = temp

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	quicksort(arr, start, splitIndex-1)
	quicksort(arr, splitIndex+1, end)
}

func allPermutations(arr []int) float64 {
	var permutations float64 = 1
	var count float64
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i-1] <= 3 {
			count++
		} else {
			if count > 0 {
				if count == 1 {
					permutations *= 2
				} else {
					if arr[i]-arr[i-int(count)-1] > 3 {
						permutations = permutations * (math.Pow(2, count) - 1)
					} else {
						permutations = permutations * math.Pow(2, count)
					}
				}
				count = 0
			}
		}
	}
	return permutations
}
