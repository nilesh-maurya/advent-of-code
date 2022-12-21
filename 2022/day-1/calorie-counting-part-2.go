package main

import (
	"bufio"
	"strconv"
)

func checkAndAddInTop3(top3Ptr *[3]int, currentCal int) {
  if currentCal > top3Ptr[0] {
    top3Ptr[2] = top3Ptr[1]
    top3Ptr[1] = top3Ptr[0]
    top3Ptr[0] = currentCal
    } else if currentCal > top3Ptr[1] {
    top3Ptr[2] = top3Ptr[1]
    top3Ptr[1] = currentCal
    } else if currentCal > top3Ptr[2] {
    top3Ptr[2] = currentCal
  }
}

func calCountingPartTwo(scanner bufio.Scanner) [3]int {
  top3 := [3]int{0, 0 , 0}
  var currentCal = 0

  for scanner.Scan() {
    line := scanner.Text()
    if len(line) > 0 {
      num, _ := strconv.Atoi(line)
      currentCal += num
    } else {
      checkAndAddInTop3(&top3, currentCal)
      currentCal = 0
    }
  }

  if (currentCal != 0) {
    checkAndAddInTop3(&top3, currentCal)
  }

  return top3
}
