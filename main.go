package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	day3 "github.com/nilesh-maurya/advent-of-code/2023/day-3"
	day4 "github.com/nilesh-maurya/advent-of-code/2023/day-4"
)

func main() {
	fmt.Printf("Day3 - Part1: %d\n", TakeInput("./2023/day-3/input.txt", day3.GearRatiosPartOne))
	// fmt.Printf("Day3 - Part2: %d\n", TakeInput("./2023/day-3/input.txt", day3.GearRatiosPartTwo))

	fmt.Printf("Day4 - Part1: %d\n", TakeInput("./2023/day-4/input.txt", day4.ScratchCardPartOne))
	fmt.Printf("Day4 - Part2: %d\n", TakeInput("./2023/day-4/input.txt", day4.ScratchCardPartTwo))

}

func TakeInput(filename string, run func(*bufio.Scanner) int) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to read input file, Err: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	return run(scanner)
}
