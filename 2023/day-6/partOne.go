package day6

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func convertPartOne(str string) []int {
	result := make([]int, 0)
	strSplit := strings.Split(str, " ")
	for _, num := range strSplit {
		if len(strings.TrimSpace(num)) > 0 {
			convertedNum, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Unable to convert string to int, Error: ", err)
			}
			result = append(result, convertedNum)
		}
	}

	return result
}

func BoatRacePartOne(scanner *bufio.Scanner) int {
	scanner.Scan()
	times := convertPartOne(strings.Split(scanner.Text(), ":")[1])
	scanner.Scan()
	distances := convertPartOne(strings.Split(scanner.Text(), ":")[1])

	ans := 1

	for timeIdx, time := range times {
		count := 0
		for idx := 0; idx <= time; idx++ {
			initialHold := idx
			currDistance := (time - idx) * initialHold
			if currDistance > distances[timeIdx] {
				count++
			}
		}
		ans *= count
	}
	return ans
}
