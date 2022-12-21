package main

import "fmt"

type option struct {
	symbol string
	value  int
}

func addAllOptions(shapes map[string]option) {
	shapes["A"] = option{symbol: "Rock", value: 1}
	shapes["B"] = option{symbol: "Paper", value: 2}
	shapes["C"] = option{symbol: "Scissor", value: 3}
	shapes["X"] = option{symbol: "Rock", value: 1}
	shapes["Y"] = option{symbol: "Paper", value: 2}
	shapes["Z"] = option{symbol: "Scissor", value: 3}
}

func addOutcomes(outcomes map[string]int) {
	outcomes["Lost"] = 0
	outcomes["Draw"] = 3
	outcomes["Won"] = 6
}

func main() {
	part1Score := calculateScorePart1()
	part2Score := calculateScorePart2()
	fmt.Println(part1Score)
	fmt.Println(part2Score)
}
