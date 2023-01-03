package main

import "fmt"

func main() {
	fmt.Println("Part1: ", tuningTroublePart1())
	fmt.Println("Part2: ", tuningTroublePart2())
}

func findFirstStartMarker(str string, windowSize int) int {
	size := len(str)
	i := 0
	j := 0
	isAllUnique := 0
	mp := make(map[string]int)

	for j < size {
		// calculation
		jChar := string(str[j])
		if _, ok := mp[jChar]; ok {
			mp[jChar] = mp[jChar] + 1
			isAllUnique++
		} else {
			mp[jChar] = 1
		}

		if j-i+1 == windowSize {
			if isAllUnique == 0 {
				return j + 1
			}

			iChar := string(str[i])
			if mp[iChar] > 1 {
				mp[iChar]--
				isAllUnique--

			} else {
				delete(mp, iChar)
			}
			i++
		}
		j++
	}

	return -1
}
