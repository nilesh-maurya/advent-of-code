package day9

import (
	"fmt"
	"strconv"
	"strings"
)

func convertNum(line string) []int {
	numStr := strings.Split(line, " ")
	nums := make([]int, len(numStr))
	for idx, num := range numStr {
		convertedNum, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Unable to convert string to int, Error: ", err)
		}
		nums[idx] = convertedNum
	}
	return nums
}
