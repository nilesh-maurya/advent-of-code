package main

import (
	"bufio"
	"log"
	"os"
)

func rucksackPart2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Unable to open file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var score = 0
	for scanner.Scan() {
		first := scanner.Text()
		scanner.Scan()
		second := scanner.Text()
		scanner.Scan()
		third := scanner.Text()

		commonStr := getCommonSubStringIn3(first, second, third)
		priority := getPriority(commonStr)
		score += priority

	}
	return score
}

func getCommonSubStringIn3(first, second, third string) string {
	for i := 0; i < len(first); i++ {
		for j := 0; j < len(second); j++ {
			for k := 0; k < len(third); k++ {
				if first[i] == second[j] && second[j] == third[k] {
					return string(first[i])
				}
			}
		}
	}

	return ""
}
