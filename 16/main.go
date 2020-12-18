package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	parts := strings.Split(string(file), "\n\n")
	intervals := []interval{}
	for _, line := range strings.Split(parts[0], "\n") {
		lineSplit := strings.Split(line, ":")
		name := lineSplit[0]
		split := strings.Split(strings.Trim(lineSplit[1], " "), "or")
		for _, split2 := range split {
			split3 := strings.Split(strings.Trim(split2, " "), "-")
			start, _ := strconv.Atoi(split3[0])
			end, _ := strconv.Atoi(split3[1])
			intervals = append(intervals, interval{start: start, end: end, name: name})
		}
	}
	intervalsName := []intervalName{}
	for _, line := range strings.Split(parts[0], "\n") {
		lineSplit := strings.Split(line, ":")
		name := lineSplit[0]
		split := strings.Split(strings.Trim(lineSplit[1], " "), "or")
		split3 := strings.Split(strings.Trim(split[0], " "), "-")
		split4 := strings.Split(strings.Trim(split[1], " "), "-")
		start1, _ := strconv.Atoi(split3[0])
		end1, _ := strconv.Atoi(split3[1])
		start2, _ := strconv.Atoi(split4[0])
		end2, _ := strconv.Atoi(split4[1])
		intervalsName = append(intervalsName, intervalName{start1: start1, end1: end1, start2: start2, end2: end2, name: name})
	}
	ticketLines := strings.Split(parts[2], "\n")
	ticketLines = ticketLines[1:]
	tickets := [][]int{}
	for _, line := range ticketLines {
		split := strings.Split(line, ",")
		ticket := []int{}
		for _, s := range split {
			n, _ := strconv.Atoi(s)
			ticket = append(ticket, n)
		}
		tickets = append(tickets, ticket)
	}

	yourTicketLines := strings.Split(parts[1], "\n")
	yourTicketLine := yourTicketLines[1]
	yourTicketSplit := strings.Split(yourTicketLine, ",")
	yourTicket := []int{}
	for _, s := range yourTicketSplit {
		num, _ := strconv.Atoi(s)
		yourTicket = append(yourTicket, num)
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i].start < intervals[j].start })
	unions := []interval{interval{start: intervals[0].start, end: intervals[0].end}}
	// fmt.Printf("Number of intervals: %d", len(intervals))
	for i := 1; i < len(intervals); i++ {
		// fmt.Printf("%d. INTERVAL: %d %d\n", i, intervals[i].start, intervals[i].end)
		if intervals[i].end < unions[0].start {
			// prepend
			newUnions := []interval{interval{start: intervals[i].start, end: intervals[i].end}}
			unions = append(newUnions, unions...)
		} else if intervals[i].end <= unions[0].end {
			// front merge
			unions[0].start = min(intervals[i].start, unions[0].start)
		} else if intervals[i].start <= unions[0].end {
			// back merge
			unions[0].start = min(intervals[i].start, unions[0].start)
			unions[0].end = max(intervals[i].end, unions[0].end)
		} else {
			// append
			unions = append(unions, intervals[i])
		}
	}
	sum := 0
	for _, ticket := range tickets {
		for _, num := range ticket {
			if num < unions[0].start || num > unions[0].end {
				sum += num
			}
		}
	}
	fmt.Println(sum)
	validTickets := [][]int{}
	for _, ticket := range tickets {
		invalid := false
		for _, num := range ticket {
			if num < unions[0].start || num > unions[0].end {
				invalid = true
				break
			}
		}
		if !invalid {
			validTickets = append(validTickets, ticket)
		}
	}

	possibilities := make([]possibility, len(intervalsName))
	for i := range possibilities {
		possibilities[i].values = map[string]bool{}
		possibilities[i].index = i
		for _, j := range intervals {
			possibilities[i].values[j.name] = true
		}
	}
	for _, ticket := range validTickets {
		for ndx, num := range ticket {
			for _, i := range intervalsName {
				if !((num >= i.start1 && num <= i.end1) || (num >= i.start2 && num <= i.end2)) {
					delete(possibilities[ndx].values, i.name)
				}
			}
		}
	}

	sort.Slice(possibilities, func(i, j int) bool { return len(possibilities[i].values) < len(possibilities[j].values) })
	// for _, p := range possibilities {
	// 	fmt.Println(p)
	// }
	namedIntervals := map[string]int{}
	excluded := map[string]bool{}
	for i, p := range possibilities {
		for px := range p.values {
			if !excluded[px] {
				excluded[px] = true
				namedIntervals[px] = possibilities[i].index
				break
			}
		}
	}

	sum2 := 1
	for name, index := range namedIntervals {
		if strings.HasPrefix(name, "departure") {
			sum2 *= yourTicket[index]
		}
	}
	fmt.Println(sum2)
	// possibilities := make([]map[string]bool, len(intervalsName))
	// for i := range possibilities {
	// 	possibilities[i] = map[string]bool{}
	// 	for _, j := range intervals {
	// 		possibilities[i][j.name] = true
	// 	}
	// }
	// for _, ticket := range validTickets {
	// 	for ndx, num := range ticket {
	// 		for _, i := range intervalsName {
	// 			if !((num >= i.start1 && num <= i.end1) || (num >= i.start2 && num <= i.end2)) {
	// 				delete(possibilities[ndx], i.name)
	// 			}
	// 		}
	// 	}
	// }

	// sort.Slice(possibilities, func(i, j int) bool { return len(possibilities[i]) < len(possibilities[j]) })

	// intervals = intervals[i:]
	// for idx := range intervals {
	// 	ci := intervals[idx]
	// 	fmt.Printf("INTERVAL: %d %d\n", ci.start, ci.end)
	// 	for udx := 0; udx < len(unions); udx++ {
	// 		cu := unions[udx]
	// 		fmt.Printf("UNION %d: %d %d\n", udx+1, cu.start, cu.end)
	// 		if intervals[idx].end < unions[udx].start {
	// 			// prepend
	// 			newUnions := []interval{interval{start: intervals[idx].start, end: intervals[idx].end}}
	// 			unions = append(newUnions, unions...)
	// 			break
	// 		} else if intervals[idx].end <= unions[udx].end {
	// 			// front expand
	// 			unions[0].start = min(intervals[idx].start, unions[udx].start)
	// 			break
	// 		} else if intervals[idx].start <= unions[udx].end && intervals[idx].end > unions[idx].end {
	// 			// last element, back expand
	// 			if udx+1 == len(unions) {
	// 				unions[udx].end = max(intervals[idx].end, unions[udx].end)
	// 				break
	// 			} else {
	// 				// check where it ends
	// 				for nudx := udx + 1; nudx < len(unions); nudx++ {
	// 					if intervals[idx].end < unions[nudx].start {
	// 						// ends before
	// 						start := min(intervals[idx].start, unions[udx].start)
	// 						end := intervals[idx].end
	// 						leftUnions := unions[0:udx]
	// 						rightUnions := unions[nudx:]
	// 						newLeftUnions := append(leftUnions, interval{start: start, end: end})
	// 						unions = append(newLeftUnions, rightUnions...)
	// 					} else if intervals[idx].end <= unions[nudx].end {
	// 						// ends in next, front merge
	// 						start := min(intervals[idx].start, unions[udx].start)
	// 						end := unions[nudx].end
	// 						leftUnions := unions[0:udx]
	// 						rightUnions := unions[nudx+1:]
	// 						newLeftUnions := append(leftUnions, interval{start: start, end: end})
	// 						unions = append(newLeftUnions, rightUnions...)
	// 					} else if nudx+1 == len(unions) {
	// 						// bigger than last, back merge
	// 						start := min(intervals[idx].start, unions[udx].start)
	// 						end := intervals[idx].end
	// 						leftUnions := unions[0:udx]
	// 						rightUnions := unions[nudx+1:]
	// 						newLeftUnions := append(leftUnions, interval{start: start, end: end})
	// 						unions = append(newLeftUnions, rightUnions...)
	// 					}
	// 				}
	// 				break
	// 			}
	// 		} else if udx == len(unions)-1 {
	// 			// last elem, append
	// 			unions = append(unions, interval{start: intervals[idx].start, end: intervals[idx].end})
	// 			break
	// 		}
	// 	}
	// 	fmt.Printf("\nUNIONS: %v\n", unions)
	// 	fmt.Println()
	// }
}

type interval struct {
	name  string
	start int
	end   int
}

type intervalName struct {
	name   string
	start1 int
	end1   int
	start2 int
	end2   int
}

type possibility struct {
	values map[string]bool
	index  int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
