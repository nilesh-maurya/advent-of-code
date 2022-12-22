package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func campCleanupPart2() int {
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

		if x2-y1 >= 0 && x1-y2 <= 0 {
			overlapCount++
		} else if x2-y1 <= 0 && x1-y2 >= 0 {
			overlapCount++
		}
	}

	return overlapCount
}
