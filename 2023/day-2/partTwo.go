package day2

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func CubeConundrumPartTwo(scanner *bufio.Scanner) int {
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameline := strings.Split(line, ": ")
		// gameId := gameline[0]

		draws := strings.Split(gameline[1], "; ")
		maxCount := map[string]int{"blue": 0, "red": 0, "green": 0}
		for _, draw := range draws {
			balls := strings.Split(draw, ", ")
			for _, ball := range balls {
				temp := strings.Split(ball, " ")
				val, err := strconv.Atoi(temp[0])
				if err != nil {
					fmt.Println("Unable to convert string to int, Error: ", err)
				}
				if maxCount[temp[1]] < val {
					maxCount[temp[1]] = val
				}
			}
		}

		ans := 1
		for _, value := range maxCount {
			ans *= value
		}

		result += ans
	}

	return result
}
