package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func supplyStacksPart2() string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Unable to open file: ", err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	// first scan input stacks of crates
	matrix := createMatrixFromInput(file, fileScanner)

	// Scan instructions
	for fileScanner.Scan() {
		lineSlice := strings.Split(fileScanner.Text(), " ")
		lineSlice = []string{lineSlice[1], lineSlice[3], lineSlice[5]}

		var instructions []int
		for _, str := range lineSlice {
			temp, _ := strconv.Atoi(str)
			instructions = append(instructions, temp)
		}
		block, from, to := instructions[0], instructions[1], instructions[2]

		tempSlice := matrix[from-1][len(matrix[from-1])-block:]
		matrix[to-1] = append(matrix[to-1], tempSlice...)
		matrix[from-1] = matrix[from-1][:len(matrix[from-1])-block]
	}

	var answer strings.Builder
	for i := 0; i < len(matrix); i++ {
		answer.WriteString(matrix[i][len(matrix[i])-1])
	}
	return answer.String()
}
