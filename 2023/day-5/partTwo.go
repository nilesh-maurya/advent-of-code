package day5

import (
	"bufio"
	"math"
)

func PartTwo(scanner *bufio.Scanner) int {

	// create input array
	input, rangeForSeeds := createArray(scanner)

	minLoc := math.MaxInt

	for idx := 0; idx < len(rangeForSeeds); idx = idx + 2 {
		startValue := rangeForSeeds[idx]
		rangeValue := rangeForSeeds[idx+1]

		for rangeIdx := 0; rangeIdx < rangeValue; rangeIdx++ {
			seed := startValue + rangeIdx
			for category := 0; category < len(input); category++ {
				for i := 0; i < len(input[category]); i++ {
					arr := input[category][i]
					destination := arr[0]
					source := arr[1]
					rangeLen := arr[2] - 1

					if seed >= source && seed <= source+rangeLen {
						seed = destination + seed - source
						break
					}

				}

				if category == len(input)-1 {
					if seed < minLoc {
						minLoc = seed
					}
				}
			}

		}
	}

	return minLoc
}
