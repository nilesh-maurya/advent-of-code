package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

/*
	approach

	read line from file
		x, y = split in half
		find the common substring(as per question it will be len of 1 always)
		commonStr = lcs(x, y)

		find priority for commonStr
			a-z => 1-26
			A-Z => 27-52

			ascii of commonStr - 'a' + 1 (if commonStr is lowercase letter)
			ascii of commonStr - 'A' + 1 + 26 (if commonStr is uppercase letter)

			totalScore += priority
*/

func rucksackPart1() int {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Unable to open file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var score = 0
	for scanner.Scan() {
		text := scanner.Text()
		first, second := text[:len(text)/2], text[len(text)/2:]

		commonStr := getCommonString(first, second)

		priority := getPriority(commonStr)
		// log.Println(priority)
		score += priority
	}
	return score
}

func getPriority(str string) int {
	// uppercase
	if strings.ToUpper(str) == str {
		return int(rune(str[0]) - 'A' + 27)
	}

	return int(rune(str[0]) - 'a' + 1)
}

func getCommonString(first, second string) string {
	var dp = make([][]int, len(first)+1)

	for i := range dp {
		dp[i] = make([]int, len(second)+1)
		for j := range dp[i] {
			dp[i][j] = 0
		}
	}

	// find Longest common substring
	var max int = math.MinInt
	var iMax = 0
	var jMax = 0

	for i := 1; i <= len(first); i++ {
		for j := 1; j <= len(second); j++ {
			if first[i-1] == second[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]

				if dp[i][j] > max {
					max = dp[i][j]
					iMax = i
					jMax = j
				}
			} else {
				dp[i][j] = 0
			}
		}
	}

	// from dp array, derive commonSubString
	var tempRes = make([]byte, max)
	for i := len(tempRes) - 1; iMax > 0 && jMax > 0; {
		if dp[iMax][jMax] == 0 {
			break
		}
		tempRes[i] = first[iMax-1]
		i--
		iMax--
		jMax--
	}

	result := string(tempRes)

	return result
}

func reverse(str string) string {
	res := []rune(str)
	fmt.Println(res)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}
