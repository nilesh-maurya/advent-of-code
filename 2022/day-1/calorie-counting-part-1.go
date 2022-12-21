package main

import (
	"bufio"
	"strconv"
)


func calCountingPartOne(scanner bufio.Scanner) int {
var maxCal int = 0
  var currentCal = 0
  for scanner.Scan() {
    line := scanner.Text();
    if line != "" {
      num, _ := strconv.Atoi(line);
      currentCal += num;
    } else {
      if maxCal < currentCal {
        maxCal = currentCal
      }
      currentCal = 0
    }
  }

  return maxCal
}
