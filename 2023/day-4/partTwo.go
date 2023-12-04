package day4

import (
	"bufio"
	"strconv"
	"strings"
)

func getMatchingNumbers(winningList string, listOfCards string) int {
	winningMap := make(map[int]bool)
	for idx := 0; idx < len(winningList); idx = idx + 3 {
		currNum, _ := strconv.Atoi(strings.TrimSpace(winningList[idx : idx+2]))
		winningMap[currNum] = true
	}

	currPoint := 0
	for idx := 0; idx < len(listOfCards); idx = idx + 3 {
		currNum, _ := strconv.Atoi(strings.TrimSpace(listOfCards[idx : idx+2]))
		if _, exist := winningMap[currNum]; exist {
			currPoint++
		}
	}

	return currPoint
}

func ScratchCardPartTwo(scanner *bufio.Scanner) int {
	ans := 0

	countMap := make(map[int]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, ": ")
		cardId, cards := lineSplit[0], lineSplit[1]

		cardsSplit := strings.Split(cards, " | ")
		winningList := cardsSplit[0]
		listOfCards := cardsSplit[1]

		cardIdSplit := strings.Split(cardId, " ")
		cardNumber, _ := strconv.Atoi(strings.TrimSpace(cardIdSplit[len(cardIdSplit)-1]))

		currPoint := getMatchingNumbers(winningList, listOfCards)

		// adding in countMap (for original count(1) )
		countMap[cardNumber]++

		for idx := 1; idx <= currPoint; idx++ {
			checkValue := cardNumber + idx
			countMap[checkValue] += countMap[cardNumber]
		}

	}
	// fmt.Println("countMap: ", countMap)

	for _, value := range countMap {
		ans += value
	}

	return ans
}
