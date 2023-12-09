package day7

import (
	"fmt"
	"strconv"
)

type hand struct {
	cards string
	bid   int
}

func convertNum(num string) int {
	convertedNum, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Unable to convert string to int, Error: ", err)
	}
	return convertedNum
}

func bubbleSort(arr *[]hand, cardMap *map[rune]int) {
	n := len(*arr)

	// Traverse through all array elements
	for i := 0; i < n-1; i++ {
		// Last i elements are already sorted, so we don't need to check them
		for j := 0; j < n-i-1; j++ {
			for idx := 0; idx < 5; idx++ {
				card1 := (*arr)[j].cards
				card2 := (*arr)[j+1].cards

				// Check card ranks using the cardMap
				if (*cardMap)[rune(card1[idx])] < (*cardMap)[rune(card2[idx])] {
					(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
					break // Break the loop if a swap is done
				} else if (*cardMap)[rune(card1[idx])] > (*cardMap)[rune(card2[idx])] {
					// No need to continue checking if ranks are different
					break
				}
			}

		}
	}
}
