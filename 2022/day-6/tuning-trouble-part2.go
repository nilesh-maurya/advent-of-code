package main

import (
	"bufio"
	"log"
	"os"
)

func tuningTroublePart2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Unable to open file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		return findFirstStartMarker(line, 14)
	}

	return 0
}
