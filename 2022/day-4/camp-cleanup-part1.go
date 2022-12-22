package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func campCleanupPart1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Unable to open file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	overlapCount := 0

	for scanner.Scan() {
		inputLine := strings.Split(scanner.Text(), ",")
		x1, x2 := getBoundary(inputLine[0])
		y1, y2 := getBoundary(inputLine[1])

		if x1-y1 == 0 || x2-y2 == 0 {
			overlapCount++
		} else if x1-y1 >= 0 && x2-y2 < 0 {
			overlapCount++
		} else if x1-y1 < 0 && x2-y2 >= 0 {
			overlapCount++
		}
	}

	return overlapCount
}

func getBoundary(input string) (int, int) {
	x := strings.Split(input, "-")
	x1, _ := strconv.Atoi(string(x[0]))
	x2, _ := strconv.Atoi(string(x[1]))

	return x1, x2
}
