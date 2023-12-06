package day6

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func convertPartTwo(str string) int {
	strSplit := strings.Split(str, " ")

	convertedNum, err := strconv.Atoi(strings.Join(strSplit, ""))
	if err != nil {
		fmt.Println("Unable to convert string to int, Error: ", err)
	}

	return convertedNum
}

func BoatRacePartTwo(scanner *bufio.Scanner) int {
	scanner.Scan()
	times := convertPartTwo(strings.Split(scanner.Text(), ":")[1])
	scanner.Scan()
	distances := convertPartTwo(strings.Split(scanner.Text(), ":")[1])

	ans := 1
	count := 0

	for idx := 0; idx <= times; idx++ {
		initialHold := idx
		currDistance := (times - idx) * initialHold
		if currDistance > distances {
			count++
		}
	}

	ans *= count
	return ans
}
