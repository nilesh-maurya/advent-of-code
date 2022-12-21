package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func main() {
  file, err :=  os.Open("./calorie-counting-part-1-data.txt");
  if err != nil {
    log.Fatal("Unable to read file, Err: ", err)
  }

  scanner := bufio.NewScanner(file);
  // maxCal := calCountingPartOne(*scanner)
  // fmt.Printf("Calorie counting part1: %v\n", maxCal)

	top3 := calCountingPartTwo(*scanner)
  fmt.Printf("Calorie counting part2: %v\n", top3)
	fmt.Println(top3[0] + top3[1] + top3[2])
}
