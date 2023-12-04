package day2

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func CubeConundrumPartOne(scanner *bufio.Scanner) int {

	config := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameline := strings.Split(line, ": ")
		gameId := gameline[0]

		draws := strings.Split(gameline[1], "; ")
		isPossible := true
		for _, draw := range draws {
			balls := strings.Split(draw, ", ")
			for _, ball := range balls {
				temp := strings.Split(ball, " ")
				if val, _ := strconv.Atoi(temp[0]); val > config[temp[1]] {
					isPossible = false
				}
			}
		}

		if isPossible {
			val, err := strconv.Atoi(strings.Split(gameId, " ")[1])
			if err != nil {
				fmt.Println("Unable to convert string to int, Error: ", err)
			}
			ans += val
		}
	}

	return ans
}
