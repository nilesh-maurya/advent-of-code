package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func getLine() string {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to read input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
	}

	return line
}

func partOne() {
	line := getLine()
	count := 0
	for _, ch := range line {
		if ch == '(' {
			count++
		} else if ch == ')' {
			count--
		}
	}

	fmt.Println("partOne: ", count)

}

func partTwo() {
	line := getLine()
	count := 0
  position := 0
	for i, ch := range line {
		if ch == '(' {
			count++
		} else if ch == ')' {
			count--
		}

		if count == -1 {
      position = i + 1
			break
		}
	}

  fmt.Println("partTwo: ", position)

}

func measureTime(run func()) {
  start := time.Now()
  run()
  elapsed := time.Since(start)
  fmt.Printf("Time took to execute: %s\n", elapsed)
}

func main() {
	measureTime(partOne)
	measureTime(partTwo)
}
