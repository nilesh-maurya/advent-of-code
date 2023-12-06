package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	day1 "github.com/nilesh-maurya/advent-of-code/2023/day-1"
	day2 "github.com/nilesh-maurya/advent-of-code/2023/day-2"
	day4 "github.com/nilesh-maurya/advent-of-code/2023/day-4"
	day5 "github.com/nilesh-maurya/advent-of-code/2023/day-5"
)

func main() {
	fmt.Printf("Day 1 - Part1: %d\n", TakeInput("./2023/day-1/input.txt", day1.CalculateCaliberationPartOne))
	fmt.Printf("Day 1 - Part2: %d\n\n", TakeInput("./2023/day-1/input.txt", day1.CalculateCaliberationPartTwo))

	fmt.Printf("Day 2 - Part1: %d\n", TakeInput("./2023/day-2/input.txt", day2.CubeConundrumPartOne))
	fmt.Printf("Day 2 - Part2: %d\n\n", TakeInput("./2023/day-2/input.txt", day2.CubeConundrumPartTwo))

	// fmt.Printf("Day 3 - Part1: %d\n", TakeInput("./2023/day-3/input.txt", day3.GearRatiosPartOne))
	// fmt.Printf("Day 3 - Part2: %d\n\n", TakeInput("./2023/day-3/input.txt", day3.GearRatiosPartTwo))

	fmt.Printf("Day 4 - Part1: %d\n", TakeInput("./2023/day-4/input.txt", day4.ScratchCardPartOne))
	fmt.Printf("Day 4 - Part2: %d\n\n", TakeInput("./2023/day-4/input.txt", day4.ScratchCardPartTwo))

	fmt.Printf("Day 5 - Part1: %d\n", TakeInput("./2023/day-5/input.txt", day5.PartOne))
	// fmt.Printf("Day 5 - Part2: %d\n\n", TakeInput("./2023/day-5/input.txt", day5.PartTwo)) // Takes long time to calculate

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
