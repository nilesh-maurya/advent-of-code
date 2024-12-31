package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sort"
	"time"
)

func measureTime(run func()) {
	start := time.Now()

	run()

	elapsed := time.Since(start)
	fmt.Printf("Time took to execute: %s\n", elapsed)
}

func getScanner() *bufio.Scanner {
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error occured while opening the file", err)
	}

	scanner := bufio.NewScanner(file)
	return scanner
}

func partOne() {
	scanner := getScanner()
	var ans int = 0
	for scanner.Scan() {
		line := scanner.Text()
		var arr [3]int
		for i, num := range strings.Split(line, "x") {
			intnum, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal("error occured while conversion to integer", err)
			}

			arr[i] = intnum
		}

		l, w, h := arr[0], arr[1], arr[2]

		area := 2*l*w + 2*w*h + 2*h*l
		extraPaper := int(math.Min(math.Min(float64(l*w), float64(w*h)), float64(h*l)))
		ans += (area + extraPaper)
	}

	fmt.Println("partOne: ", ans)
}

func partTwo() {
	scanner := getScanner()
	var ans int = 0
	for scanner.Scan() {
		line := scanner.Text()
		var arr [3]int
		for i, num := range strings.Split(line, "x") {
			intnum, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal("error occured while conversion to integer", err)
			}

			arr[i] = intnum
		}

		l, w, h := arr[0], arr[1], arr[2]

		bow := l * w * h
		
		slice := arr[:]
		sort.Ints(slice)
		ribbon := slice[0] * 2 + slice[1] * 2
		
		ans += (bow + ribbon)
	}

	fmt.Println("partTwo: ", ans)
}

func main() {
	measureTime(partOne)
	measureTime(partTwo)
}
