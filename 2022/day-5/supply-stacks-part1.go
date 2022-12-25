package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func supplyStacksPart1() string {
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

		for i := 1; i <= block; i++ {
			lastEle := matrix[from-1][len(matrix[from-1])-1]
			matrix[to-1] = append(matrix[to-1], lastEle)
			matrix[from-1] = matrix[from-1][:len(matrix[from-1])-1]
		}
	}

	var answer strings.Builder
	for i := 0; i < len(matrix); i++ {
		answer.WriteString(matrix[i][len(matrix[i])-1])
	}
	return answer.String()
}

func print2DSlice(slice [][]string) {
	fmt.Println("[")
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i])
		fmt.Print(",\n")
	}
	fmt.Println("]")
}

func getMeasureOfStackCrates(file *os.File) (int, int) {
	count := 0
	lengthOfStack := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, " 1") {
			for _, char := range line {
				c := string(char)
				if c != " " {
					count++
				}
			}
			break
		}
		lengthOfStack++
	}

	file.Seek(0, io.SeekStart)
	return count, lengthOfStack
}

func createMatrixFromInput(file *os.File, fileScanner *bufio.Scanner) [][]string {
	length, lengthOfStack := getMeasureOfStackCrates(file)
	var matrix = make([][]string, length)
	for i := 0; i < len(matrix); i++ {
		inner := make([]string, 0, lengthOfStack)
		matrix[i] = inner
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(line, " 1") {
			fileScanner.Scan()
			break
		}
		for i, inner := 1, 0; i < len(line) || inner < len(matrix); i, inner = i+4, inner+1 {
			if string(line[i]) != " " {
				matrix[inner] = append([]string{string(line[i])}, matrix[inner]...)
			}
		}
	}

	return matrix
}
