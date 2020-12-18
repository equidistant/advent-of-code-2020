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
	sequences := []sequence{}
	var currentSequence sequence
	for _, line := range lines {
		lineSplit := strings.Fields(line)
		if lineSplit[0] == "mask" {
			if currentSequence.mask != lineSplit[2] && len(currentSequence.mask) > 0 {
				sequences = append(sequences, currentSequence)
			}
			currentSequence = sequence{mask: lineSplit[2], operations: []operation{}}
		} else {
			address, _ := strconv.Atoi(lineSplit[0][4 : len(lineSplit[0])-1])
			value, _ := strconv.Atoi(lineSplit[2])
			currentSequence.operations = append(currentSequence.operations, operation{address: uint64(address), value: uint64(value)})
		}
	}
	sequences = append(sequences, currentSequence)
	sum := firstPart(sequences)
	sum2 := secondPart(sequences)
	fmt.Printf("%v\n", sum)
	fmt.Printf("%v\n", sum2)
}

func getAndMask(mask string) uint64 {
	var andMask uint64
	for i := len(mask) - 1; i >= 0; i-- {
		bit := mask[i]
		switch bit {
		case '0':
			continue
		case '1':
			andMask += uint64(math.Pow(2, float64(len(mask)-1-i)))
		case 'X':
			andMask += uint64(math.Pow(2, float64(len(mask)-1-i)))
		}
	}
	return andMask
}

func getOrMask(mask string) uint64 {
	var orMask uint64
	for i := len(mask) - 1; i >= 0; i-- {
		bit := mask[i]
		switch bit {
		case '0':
			continue
		case '1':
			orMask += uint64(math.Pow(2, float64(len(mask)-1-i)))
		case 'X':
			continue
		}
	}
	return orMask
}

func firstPart(sequences []sequence) uint64 {
	addresses := map[uint64]uint64{}

	for _, sequence := range sequences {
		andMask := getAndMask(sequence.mask)
		orMask := getOrMask(sequence.mask)
		// fmt.Printf("%.36b\n", andMask)
		// fmt.Printf("%.36b\n", orMask)
		for _, operation := range sequence.operations {
			// fmt.Printf("%.36b\n", operation.value)
			maskedValue1 := operation.value & andMask
			// fmt.Printf("%.36b\n", maskedValue1)
			maskedValue2 := maskedValue1 | orMask
			// fmt.Printf("%.36b\n", maskedValue2)
			addresses[operation.address] = maskedValue2
		}
	}
	var sum uint64
	for _, value := range addresses {
		sum += value
	}
	return sum
}

func getOrMask2(mask string) uint64 {
	var orMask uint64
	for i := len(mask) - 1; i >= 0; i-- {
		bit := mask[i]
		switch bit {
		case '0':
			continue
		case '1':
			orMask += uint64(math.Pow(2, float64(len(mask)-1-i)))
		case 'X':
			continue
		}
	}
	return orMask
}

func getXorMasks(mask string) []uint64 {
	floatingMasks := []uint64{}
	for i := len(mask) - 1; i >= 0; i-- {
		bit := mask[i]
		switch bit {
		case '0':
			continue
		case '1':
			continue
		case 'X':
			if len(floatingMasks) == 0 {
				floatingMasks = append(floatingMasks, 0)
				floatingMasks = append(floatingMasks, 1)
			} else {
				newFloatingMasks := []uint64{}
				for _, floatingMask := range floatingMasks {
					newFloatingMask := floatingMask + uint64(math.Pow(2, float64(len(mask)-1-i)))
					newFloatingMasks = append(newFloatingMasks, floatingMask)
					newFloatingMasks = append(newFloatingMasks, newFloatingMask)
				}
				floatingMasks = newFloatingMasks
			}
		}
	}
	return floatingMasks
}

func getXAddress(address string, mask string) string {
	uintAddress, _ := strconv.Atoi(address)
	uint64Address := uint64(uintAddress)
	stringAddress := fmt.Sprintf("%.36b", uint64Address)
	var xAddress string
	for i := len(mask) - 1; i >= 0; i-- {
		bit := mask[i]
		switch bit {
		case '0':
			// fmt.Printf("%s", string(stringAddress[len(mask)-1-i]))
			xAddress = fmt.Sprintf("%s%s", string(stringAddress[i]), xAddress)
		case '1':
			xAddress = fmt.Sprintf("1%s", xAddress)
		case 'X':
			xAddress = fmt.Sprintf("X%s", xAddress)
		}
	}
	return xAddress
}

func getFloatingAddresses(mask string) []uint64 {
	floatingMasks := []uint64{}
	for i := len(mask) - 1; i >= 0; i-- {
		bit := mask[i]
		switch bit {
		case '0':
			continue
		case '1':
			if len(floatingMasks) == 0 {
				floatingMasks = append(floatingMasks, 1)
			} else {
				newFloatingMasks := []uint64{}
				for _, floatingMask := range floatingMasks {
					newFloatingMasks = append(newFloatingMasks, floatingMask+uint64(math.Pow(2, float64(len(mask)-1-i))))
				}
				floatingMasks = newFloatingMasks
			}
		case 'X':
			if len(floatingMasks) == 0 {
				floatingMasks = append(floatingMasks, 0)
				floatingMasks = append(floatingMasks, 1)
			} else {
				newFloatingMasks := []uint64{}
				for _, floatingMask := range floatingMasks {
					newFloatingMask := floatingMask + uint64(math.Pow(2, float64(len(mask)-1-i)))
					newFloatingMasks = append(newFloatingMasks, floatingMask)
					newFloatingMasks = append(newFloatingMasks, newFloatingMask)
				}
				floatingMasks = newFloatingMasks
			}
		}
	}
	return floatingMasks
}

func secondPart(sequences []sequence) uint64 {
	addresses := map[uint64]uint64{}
	for _, sequence := range sequences {
		for _, operation := range sequence.operations {
			xAddress := getXAddress(fmt.Sprint(operation.address), fmt.Sprint(sequence.mask))
			floatingAddresses := getFloatingAddresses(xAddress)
			for _, floatingAddress := range floatingAddresses {
				addresses[floatingAddress] = operation.value
			}
		}
	}
	var sum uint64
	for _, value := range addresses {
		sum += value
	}
	return sum
}

func thirdPart(sequences []sequence) uint64 {
	addresses := map[uint64]uint64{}
	for _, sequence := range sequences {
		orMask := getOrMask2(sequence.mask)
		xorMasks := getXorMasks(sequence.mask)
		for _, operation := range sequence.operations {
			orAddress := operation.address | orMask
			for _, xorMask := range xorMasks {
				actualAddress := orAddress ^ xorMask
				fmt.Printf("%.6b\n", orAddress)
				fmt.Printf("%.6b\n", xorMask)
				fmt.Printf("%.6b %v\n", actualAddress, actualAddress)
				addresses[actualAddress] = operation.value
			}
		}
	}
	var sum uint64
	for _, value := range addresses {
		sum += value
	}
	return sum
}

type sequence struct {
	mask       string
	operations []operation
}

type operation struct {
	address uint64
	value   uint64
}
