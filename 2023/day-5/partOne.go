package day5

import (
	"bufio"
	"math"
)

func PartOne(scanner *bufio.Scanner) int {

	// create input array
	input, seeds := createArray(scanner)

	seedsCopy := make([]int, len(seeds))
	copy(seedsCopy, seeds)

	minLoc := math.MaxInt

	for _, seed := range seedsCopy {
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

	return minLoc
}
