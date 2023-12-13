package day8

import (
	"bufio"
	"strings"
)

func parseInputForPartOne(scanner *bufio.Scanner) (string, map[string][]string) {
	scanner.Scan()
	instruction := scanner.Text()
	nodesMap := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			lineSplit := strings.Split(line, " = ")
			key := lineSplit[0]
			values := make([]string, 2)
			neighbor := strings.Split(lineSplit[1], ", ")
			values[0] = neighbor[0][1:]
			values[1] = neighbor[1][:len(neighbor[1])-1]

			nodesMap[key] = values
		}
	}

	return instruction, nodesMap
}

func HauntedWastelandPartOne(scanner *bufio.Scanner) int {
	instruction, nodesMap := parseInputForPartOne(scanner)

	currentNode := "AAA"
	idx := 0
	count := 0

	for currentNode != "ZZZ" && idx < len(instruction) {
		if instruction[idx] == 'R' {
			currentNode = nodesMap[currentNode][1]
		} else if instruction[idx] == 'L' {
			currentNode = nodesMap[currentNode][0]
		}

		idx++
		count++

		if currentNode == "ZZZ" && idx == len(instruction) {
			break
		} else if idx == len(instruction) {
			idx = 0
		}
	}

	return count
}
