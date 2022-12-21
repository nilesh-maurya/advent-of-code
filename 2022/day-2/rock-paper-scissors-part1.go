package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

/*
shapeSelected (1-> rock, 2 -> paper, 3-> scissor)
Outcome of the round (0 -> lost, 3 -> draw, 6 -> Won)
each round = shapeSelected + Outcome of the round
total score = sum(each round)


create a map with (A,B,C,X,Y,Z) with respective rock,paper,scissor's value
create outcome map

for each line of file
	opponent's choice
	my choice

	result = determine who won?
	score = my choice + outcome[result]
	total score += score
*/

func determineResult(opponentCh option, myCh option) string {
	if opponentCh.symbol == myCh.symbol {
		return "Draw"
	}

	if opponentCh.symbol == "Rock" {
		if myCh.symbol == "Scissor" {
			return "Lose"
		}

		if myCh.symbol == "Paper" {
			return "Won"
		}
	}

	if opponentCh.symbol == "Paper" {
		if myCh.symbol == "Rock" {
			return "Lose"
		}

		if myCh.symbol == "Scissor" {
			return "Won"
		}
	}

	if opponentCh.symbol == "Scissor" {
		if myCh.symbol == "Rock" {
			return "Won"
		}

		if myCh.symbol == "Paper" {
			return "Lose"
		}
	}

	return "Draw"
}
func calculateScorePart1() int {
	totalScore := 0

	shapes := make(map[string]option)
	addAllOptions(shapes)

	outcomes := make(map[string]int)
	addOutcomes(outcomes)

	// read file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Unable to open file: ", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		opponentCh := shapes[line[0]]
		myCh := shapes[line[1]]

		res := determineResult(opponentCh, myCh)
		// fmt.Println(opponentCh, myCh, "result is: "+res)
		score := myCh.value + outcomes[res]
		totalScore += score
	}

	return totalScore
}
