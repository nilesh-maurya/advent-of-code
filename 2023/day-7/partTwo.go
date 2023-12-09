package day7

import (
	"bufio"
	"strings"
)

// High Card             =>return : 0             => all will be distincts                            => distinct: 5
// One Pair              =>return : 1             => one will be repeated, rest will be distincts     => distinct: 4
// Two Pair              =>return : 2             => two will be repeated, one will be distinct       => distinct: 3
// Three of a kind       =>return : 3             => one will repeated, rest are distict              => distinct: 3
// Full House            =>return : 4             => one will be repeated trice, one is repeated once => distinct: 2
// Four of a Kind        =>return : 5             => one will be repeated four times, remaining one   => distinct: 2
// Five of a kind        =>return : 6             => one will be repeated five times                  => distinct: 1

// Joker can pretend to be any card which can be Best
func findTypeOfCardPartTwo(h *hand) int {
	countMap := make(map[string]int)
	var index int
	for idx := 0; idx < len(h.cards); idx++ {
		countMap[string(h.cards[idx])]++
	}

	if countMap["J"] > 0 && len(countMap) > 1 {
		var hKey string
		hValue := 0
		for key, value := range countMap {
			if value > hValue && key != "J" {
				hValue = value
				hKey = key
			}
		}

		countMap[hKey] += countMap["J"]
		delete(countMap, "J")
	}

	length := len(countMap)

	if length == 1 { // Five of kind
		index = 6
	} else if length == 4 { // One pair
		index = 1
	} else if length == 5 { // High card
		index = 0
	} else if length == 2 {
		higherValue := 0
		for _, value := range countMap {
			if value > higherValue {
				higherValue = value
			}
		}
		if higherValue == 4 { // Four of a kind
			index = 5
		} else { // Full house
			index = 4
		}
	} else if length == 3 {
		higherValue := 0
		for _, value := range countMap {
			if value > higherValue {
				higherValue = value
			}
		}
		if higherValue == 3 { // Three of a kind
			index = 3
		} else { // Two pair
			index = 2
		}
	}

	return index
}

func CamelCardsPartTwo(scanner *bufio.Scanner) int {
	cardMap := map[rune]int{
		'A': 13, 'K': 12, 'Q': 11, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2, 'J': 1,
	}

	input := make([][]hand, 7)
	rank := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		temp := hand{
			cards: line[0],
			bid:   convertNum(line[1]),
		}

		index := findTypeOfCardPartTwo(&temp)
		input[index] = append(input[index], temp)
		rank++
	}

	// sort
	for idx := range input {
		bubbleSort(&input[idx], &cardMap)
	}

	// calculate winnings
	ans := 0
	for i := 6; i >= 0; i-- {
		for _, h := range input[i] {
			ans += rank * h.bid
			rank--
		}
	}

	return ans
}
