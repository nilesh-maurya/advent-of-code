package day9

import (
	"bufio"
)

func MirageMaintenancePartTwo(scanner *bufio.Scanner) int {
	ans := 0

	for scanner.Scan() {
		nums := convertNum(scanner.Text())
		// fmt.Println("Current line: ", nums)

		firstValue := make([]int, 0)
		isAllZero := false

		for !isAllZero {
			firstValue = append(firstValue, nums[0])
			diffArr := make([]int, 0)

			for idx := 1; idx < len(nums); idx++ {
				diffArr = append(diffArr, nums[idx]-nums[idx-1])
			}
			nums = diffArr

			isAllZero = true // set initial to true
			for _, val := range diffArr {
				if val != 0 {
					isAllZero = false
					break
				}
			}
		}

		diff := 0
		for idx := len(firstValue) - 1; idx >= 0; idx-- {
			diff = firstValue[idx] - diff
		}

		ans += diff
	}

	return ans
}
