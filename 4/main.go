package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := ioutil.ReadFile("./input.txt")
	check(err)
	splitByDoubleNewLine := strings.Split(string(file), "\n\n")
	count := 0
	validPassports := [][]string{}
	for i := 0; i < len(splitByDoubleNewLine); i++ {
		passport := strings.FieldsFunc(splitByDoubleNewLine[i], splitByNewLineAndSpace)
		// valid := firstCheck(passport)
		valid := secondCheck(passport)
		if valid {
			count++
			validPassports = append(validPassports, passport)
		}
	}
	fmt.Println(count)
	// for _, pass := range validPassports {
	// 	fmt.Println(pass)
	// }
}

func firstCheck(passport []string) bool {
	if len(passport) == 8 {
		return true
	}
	if len(passport) == 7 {
		for _, elem := range passport {
			elemSplit := strings.Split(elem, ":")
			if elemSplit[0] == "cid" {
				return false
			}
		}
		return true
	}
	return false
}

func secondCheck(passport []string) bool {
	if len(passport) == 7 || len(passport) == 8 {
		for _, elem := range passport {
			elemSplit := strings.Split(elem, ":")
			key := elemSplit[0]
			value := elemSplit[1]
			switch key {
			case "cid":
				continue
			case "byr":
				byr, err := strconv.Atoi(value)
				check(err)
				if byr < 1920 || byr > 2002 {
					// fmt.Printf("++++NO: %d\n", byr)
					return false
				}
				// fmt.Printf("YES: %d\n", byr)
			case "iyr":
				iyr, err := strconv.Atoi(value)
				check(err)
				if iyr < 2010 || iyr > 2020 {
					// fmt.Printf("++++NO: %d\n", iyr)
					return false
				}
				// fmt.Printf("YES: %d\n", iyr)
			case "eyr":
				eyr, err := strconv.Atoi(value)
				check(err)
				if eyr < 2020 || eyr > 2030 {
					return false
				}
			case "hgt":
				measure := value[len(value)-2:]
				if measure != "in" && measure != "cm" {
					return false
				}
				if measure == "cm" {
					hgt, err := strconv.Atoi(value[:len(value)-2])
					check(err)
					if hgt < 150 || hgt > 193 {
						// fmt.Printf("+++++NO: %s\n", value)
						return false
					}
				}
				if measure == "in" {
					hgt, err := strconv.Atoi(value[:len(value)-2])
					check(err)
					if hgt < 59 || hgt > 76 {
						// fmt.Printf("+++++NO: %s\n", value)
						return false
					}
				}
			case "hcl":
				match, _ := regexp.MatchString("^#[0-9a-fA-f]{6}$", value)
				if !match {
					// fmt.Printf("+++++NO: %s\n", value)
					return false
				}
				// fmt.Printf("YES: %s\n", value)
			case "ecl":
				match, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth){1}$", value)
				if !match {
					// fmt.Printf("+++++NO: %s\n", value)
					return false
				}
				// fmt.Printf("YES: %s\n", value)
			case "pid":
				match, _ := regexp.MatchString("^[0-9]{9}$", value)
				if !match {
					// fmt.Printf("NOO: %s\n", value)
					return false
				}
				// fmt.Printf("YES: %s\n", value)
			}
		}
		return true
	}
	return false
}

func splitByNewLineAndSpace(r rune) bool {
	return unicode.IsSpace(r)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
