package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	players := strings.Split(string(file), "\n\n")
	player1 := []int{}
	player2 := []int{}

	nums1 := strings.Split(players[0][1:], "\n")
	nums2 := strings.Split(players[1][1:], "\n")

	for i := 1; i < len(nums1); i++ {
		num1, _ := strconv.Atoi(string(nums1[i]))
		num2, _ := strconv.Atoi(string(nums2[i]))
		player1 = append(player1, num1)
		player2 = append(player2, num2)
	}

	score := part1(player1, player2)
	score2 := part2([][]int{player1, player2})
	fmt.Println(score)
	fmt.Println(score2)
}

func part1(player1, player2 []int) int {
	for len(player1) > 0 && len(player2) > 0 {

		if player1[0] > player2[0] {
			after1 := player1[1:]
			newPlayer1 := append(after1, []int{player1[0], player2[0]}...)
			player2 = player2[1:]
			player1 = newPlayer1
		} else {
			after2 := player2[1:]
			newPlayer2 := append(after2, []int{player2[0], player1[0]}...)
			player1 = player1[1:]
			player2 = newPlayer2
		}
	}
	var winner []int
	if len(player1) > len(player2) {
		winner = player1
	} else {
		winner = player2
	}
	sum := 0
	for i := 0; i < len(winner); i++ {
		sum += winner[len(winner)-i-1] * (i + 1)
	}
	return sum
}

func part2(players [][]int) int {
	winner := 0
	round := 1
	for len(players[0]) > 0 && len(players[1]) > 0 {
		// fmt.Printf("******** %d *********\n", round)
		// fmt.Printf("Player 1's deck: %v len: %d\n", players[0], len(players[0]))
		// fmt.Printf("Player 2's deck: %v len: %d\n", players[1], len(players[1]))
		// fmt.Printf("Player 1 plays: %d\n", players[0][0])
		// fmt.Printf("Player 2 plays: %d\n", players[1][0])
		winner = 0
		// fmt.Printf("%t %t\n", len(players[0]) >= players[0][0], len(players[1]) >= players[1][0])
		if len(players[0][1:]) >= players[0][0] && len(players[1][1:]) >= players[1][0] {
			copy1 := make([]int, len(players[0][1:1+players[0][0]]))
			copy(copy1, players[0][1:1+players[0][0]])
			copy2 := make([]int, len(players[1][1:1+players[1][0]]))
			copy(copy2, players[1][1:1+players[1][0]])
			winner = recursivePart2([][]int{copy1, copy2}, map[string]bool{})
		} else {
			if players[1][0] > players[0][0] {
				winner = 1
			}
		}

		// fmt.Printf("Player %d wins round %d\n\n", winner+1, round)
		next := players[winner][1:]
		newWinner := append(next, []int{players[winner][0], players[xor(winner)][0]}...)
		players[xor(winner)] = players[xor(winner)][1:]
		players[winner] = newWinner
		round++
	}
	sum := 0
	fmt.Println(players[winner])
	for i := 0; i < len(players[winner]); i++ {
		sum += players[winner][len(players[winner])-i-1] * (i + 1)
	}
	return sum
}

func recursivePart2(players [][]int, history map[string]bool) int {
	winner := 0
	for len(players[0]) > 0 && len(players[1]) > 0 {

		// fmt.Printf("\nxxxxxxxxxxxxxx %d xxxxxxxxxxxxxxn\n", len(history)+1)
		// fmt.Printf("Player 1's deck: %v len: %d\n", players[0], len(players[0]))
		// fmt.Printf("Player 2's deck: %v len: %d\n", players[1], len(players[1]))
		// fmt.Printf("Player 1 plays: %d\n", players[0][0])
		// fmt.Printf("Player 2 plays: %d\n", players[1][0])

		winner = 0
		s1 := arrayToString(players[0], ",")
		s2 := arrayToString(players[1], ",")
		s := s1 + ":" + s2
		if history[s] {
			return 0
		}
		history[s] = true
		if len(players[0][1:]) >= players[0][0] && len(players[1][1:]) >= players[1][0] {
			copy1 := make([]int, len(players[0][1:1+players[0][0]]))
			copy(copy1, players[0][1:1+players[0][0]])
			copy2 := make([]int, len(players[1][1:1+players[1][0]]))
			copy(copy2, players[1][1:1+players[1][0]])
			winner = recursivePart2([][]int{copy1, copy2}, history)
		} else {
			if players[1][0] > players[0][0] {
				winner = 1
			}
		}
		next := players[winner][1:]
		newWinner := append(next, []int{players[winner][0], players[xor(winner)][0]}...)
		players[xor(winner)] = players[xor(winner)][1:]
		players[winner] = newWinner
	}
	return winner
}

func xor(i int) int {
	if i == 0 {
		return 1
	}
	return 0
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
