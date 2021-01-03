package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	ingredients := [][]string{}
	allergens := []map[string]bool{}
	allAllergens := map[string]bool{}
	for _, line := range strings.Split(string(file), "\n") {
		splitIdx := strings.Index(line, "(")
		ingredientsArr := strings.Fields(strings.Trim(line[:splitIdx], " "))
		ingredients = append(ingredients, ingredientsArr)
		allergensArr := strings.Split(strings.Trim(line[splitIdx+9:len(line)-1], " "), ", ")
		allergensMap := map[string]bool{}
		for _, a := range allergensArr {
			allergensMap[a] = true
			allAllergens[a] = true
		}
		allergens = append(allergens, allergensMap)
	}

	intersections := map[string][]string{}
	for allergen := range allAllergens {
		intersection := []string{}
		for i := range ingredients {
			if allergens[i][allergen] {
				if len(intersection) == 0 {
					intersection = append(intersection, ingredients[i]...)
				} else {
					intersection = getIntersection(intersection, ingredients[i])
				}
			}
		}
		intersections[allergen] = intersection
	}
	notDefinite := true

	for notDefinite {
		var singleIngredient string
		for _, ingredients := range intersections {
			if len(ingredients) == 1 {
				singleIngredient = ingredients[0]
				break
			}
		}
		for allergen := range intersections {
			if len(intersections[allergen]) > 1 {
				idx := -1
				for i := range intersections[allergen] {
					if intersections[allergen][i] == singleIngredient {
						idx = i
						break
					}
				}
				if idx != -1 {
					before := intersections[allergen][:idx]
					after := intersections[allergen][idx+1:]
					reducedIngredients := append(before, after...)
					intersections[allergen] = reducedIngredients
				}
			}
		}
		notDefinite = false
		for allergen := range intersections {
			if len(intersections[allergen]) > 1 {
				notDefinite = true
				break
			}
		}
	}

	allergenicIngredients := map[string]bool{}

	for _, value := range intersections {
		allergenicIngredients[value[0]] = true
	}

	count := 0
	for _, i := range ingredients {
		for _, elem := range i {
			if !allergenicIngredients[elem] {
				count++
			}
		}
	}
	sorted := [][]string{}
	for key, value := range intersections {
		sorted = append(sorted, []string{string(key), string(value[0])})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] < sorted[j][0]
	})
	for _, s := range sorted {
		fmt.Printf("%s,", s[1])
	}
	fmt.Println(count)
	fmt.Println(sorted)
}

func getIntersection(a []string, b []string) []string {
	c := map[string]bool{}
	result := []string{}
	for _, ax := range a {
		c[ax] = true
	}
	for _, bx := range b {
		if c[bx] {
			result = append(result, bx)
		}
	}
	return result
}
