package day4

import (
	"bufio"
	"strconv"
	"strings"
)

func ScratchCardPartOne(scanner *bufio.Scanner) int {
	ans := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, ": ")
		cards := lineSplit[1]
		cardsSplit := strings.Split(cards, " | ")
		winningList := cardsSplit[0]
		listOfCards := cardsSplit[1]

		winningMap := make(map[int]bool)
		for idx := 0; idx < len(winningList); idx = idx + 3 {
			currNum, _ := strconv.Atoi(strings.TrimSpace(winningList[idx : idx+2]))
			winningMap[currNum] = true
		}

		currPoint := 0
		for idx := 0; idx < len(listOfCards); idx = idx + 3 {
			currNum, _ := strconv.Atoi(strings.TrimSpace(listOfCards[idx : idx+2]))
			if _, exist := winningMap[currNum]; exist {
				if currPoint == 0 {
					currPoint = 1
				} else {
					currPoint *= 2
				}
			}
		}
		ans += currPoint
	}

	return ans
}
