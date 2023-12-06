package day5

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func convert(str string) []int {
	result := make([]int, 0)
	strSplit := strings.Split(str, " ")
	for _, num := range strSplit {
		convertedNum, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Unable to convert string to int, Error: ", err)
		}
		result = append(result, convertedNum)
	}

	return result
}

func createArray(scanner *bufio.Scanner) ([][][]int, []int) {
	input := make([][][]int, 0)

	scanner.Scan()
	seedSPlit := strings.Split(scanner.Text(), ": ")
	seeds := convert(seedSPlit[1])
	scanner.Scan()

	temp := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			input = append(input, temp)
			temp = nil
		} else if line[len(line)-1] == ':' {
			// current = strings.Split(strings.Split(line, " ")[0], "-to-")[1]
		} else {
			temp = append(temp, convert(line))
		}
	}

	if len(temp) > 0 {
		input = append(input, temp)

	}

	return input, seeds
}
