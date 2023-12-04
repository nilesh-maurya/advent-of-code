package day1

import (
	"bufio"
	"strconv"
)

func isNumeric(val byte) bool {
	_, err := strconv.Atoi(string(val))
	return err == nil
}

func CalculateCaliberation(scanner *bufio.Scanner) int {
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		forward := false
		backward := false
		digit := make([]byte, 2)
		for i, j := 0, len(line)-1; i < len(line); i, j = i+1, j-1 {
			if !forward && isNumeric(line[i]) {
				digit[0] = line[i]
				forward = true
			}
			if !backward && isNumeric(line[j]) {
				digit[1] = line[j]
				backward = true
			}
			if forward && backward {
				break
			}
		}
		num, _ := strconv.Atoi(string(digit))
		ans += num
	}
	return ans
}
