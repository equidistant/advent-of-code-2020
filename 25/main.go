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
	publicKeys := []int{}
	for _, l := range lines {
		publicKey, _ := strconv.Atoi(l)
		publicKeys = append(publicKeys, publicKey)
	}
	part1(publicKeys)
}

func part1(publicKeys []int) {
	loops := getLoops(publicKeys)
	fmt.Println(loops)
	for l1 := range loops[0] {
		for l2 := range loops[1] {
			privateKey1 := getPrivateKey(publicKeys[0], l2)
			privateKey2 := getPrivateKey(publicKeys[1], l1)
			if privateKey1 == privateKey2 {
				fmt.Println(privateKey1)
			}
		}
	}
	// privateKey1 := getPrivateKey(publicKeys[0], loops[1])
	// privateKey2 := getPrivateKey(publicKeys[1], loops[0])
	// fmt.Println(privateKey1)
	// fmt.Println(privateKey2)

}

func getLoops(publicKeys []int) []map[int]bool {
	loops := []map[int]bool{map[int]bool{}, map[int]bool{}}
	for i := 1; i < 1000000; i++ {
		if i%10000 == 0 {
			fmt.Printf("%d/%d\n", i/10000, 100)
		}
		value := 1
		for j := 0; j < 100000; j++ {
			if value == publicKeys[0] {
				loops[0][j] = true
			} else if value == publicKeys[1] {
				loops[1][j] = true
			}
			value *= i
			value %= 20201227
		}
	}
	return loops
}

func getPrivateKey(publicKey int, loop int) int {
	privateKey := 1
	for i := 0; i < loop; i++ {
		privateKey *= publicKey
		privateKey %= 20201227
	}
	return privateKey
}
