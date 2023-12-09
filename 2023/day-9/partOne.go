package day9

import (
	"bufio"
)

func MirageMaintenancePartOne(scanner *bufio.Scanner) int {
	ans := 0

	for scanner.Scan() {
		nums := convertNum(scanner.Text())

		lastValue := make([]int, 0)
		isAllZero := false

		for !isAllZero {
			lastValue = append(lastValue, nums[len(nums)-1])
			diffArr := make([]int, 0)

			for idx := 1; idx < len(nums); idx++ {
				diffArr = append(diffArr, nums[idx]-nums[idx-1])
			}

			isAllZero = true // set initial to true
			for _, val := range diffArr {
				if val != 0 {
					isAllZero = false
					break
				}
			}

			nums = diffArr
		}

		diff := 0
		for idx := len(lastValue) - 1; idx >= 0; idx-- {
			diff = lastValue[idx] + diff
		}

		ans += diff
	}

	return ans
}
