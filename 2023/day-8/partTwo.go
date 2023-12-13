package day8

import (
	"bufio"
	"math"
	"strings"
)

func parseInputForPartTwo(scanner *bufio.Scanner) (string, []string, map[string][]string) {
	scanner.Scan()
	instruction := scanner.Text()
	nodesMap := make(map[string][]string)
	startingNode := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			lineSplit := strings.Split(line, " = ")
			key := lineSplit[0]

			// check for starting node
			if key[len(key)-1] == 'A' {
				startingNode = append(startingNode, key)
			}

			values := make([]string, 2)
			neighbor := strings.Split(lineSplit[1], ", ")
			values[0] = neighbor[0][1:]
			values[1] = neighbor[1][:len(neighbor[1])-1]

			nodesMap[key] = values
		}
	}

	return instruction, startingNode, nodesMap
}

func findLoopLength(currentNode string, nodesMap map[string][]string, instruction string) int {
	idx := 0
	length := 0
	currentNodeCopy := currentNode
	for {
		currentInstr := instruction[idx]
		if currentInstr == 'R' {
			currentNode = nodesMap[currentNode][1]
		} else if currentInstr == 'L' {
			currentNode = nodesMap[currentNode][0]
		}

		idx++
		length++

		if currentNode == currentNodeCopy {
			break
		}
		if idx == len(instruction) {
			idx = 0
		}
	}

	return length
}

func lcm(a int, b int) int {
	greater := int(math.Max(float64(a), float64(b)))
	smaller := int(math.Min(float64(a), float64(b)))

	for idx := greater; ; idx += greater {
		if idx%smaller == 0 {
			return idx
		}
	}
}
func findLCM(loopLengths []int) int {
	a := loopLengths[0]

	for idx := 1; idx < len(loopLengths); idx++ {
		b := loopLengths[idx]
		a = lcm(a, b)
	}
	return a
}

func HauntedWastelandPartTwo(scanner *bufio.Scanner) int {
	instruction, startingNode, nodesMap := parseInputForPartTwo(scanner)
	loopLengths := make([]int, len(startingNode))

	for nodeIdx, node := range startingNode {
		currentNode := node
		instrIndx := 0
		for {
			currentInstr := instruction[instrIndx]
			if currentInstr == 'R' {
				currentNode = nodesMap[currentNode][1]
			} else if currentInstr == 'L' {
				currentNode = nodesMap[currentNode][0]
			}
			instrIndx++

			if currentNode[len(currentNode)-1] == 'Z' && instrIndx == len(instruction) {
				// we fount node ending with Z
				// find loop length
				loopLengths[nodeIdx] = findLoopLength(currentNode, nodesMap, instruction)

				break
			} else if instrIndx == len(instruction) {
				instrIndx = 0
			}
		}
	}

	return findLCM(loopLengths)
}
