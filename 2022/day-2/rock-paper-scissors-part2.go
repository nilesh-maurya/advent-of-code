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


create a map with (A,B,C) with respective rock,paper,scissor's value
create a map with (X,Y,Z) with respective outcome
create outcome map with value

for each line of file
	opponent's choice
	result of the round

	myChoice = determine myChoice based on "result of the round"
	score = myChoice + outcome[result of the round]
	total score += score
*/

func possibleResult(results map[string]string) {
	results["X"] = "Lose"
	results["Y"] = "Draw"
	results["Z"] = "Won"
}

func createOption(symbol string) option {
	if symbol == "Rock" {
		return option{symbol: "Rock", value: 1}
	} else if symbol == "Paper" {
		return option{symbol: "Paper", value: 2}
	} else {
		return option{symbol: "Scissor", value: 3}
	}
}

func determineMyChoice(opponentCh option, result string) option {
	if result == "Won" {
		if opponentCh.symbol == "Rock" {
			return createOption("Paper")
		}
		if opponentCh.symbol == "Paper" {
			return createOption("Scissor")
		}
		if opponentCh.symbol == "Scissor" {
			return createOption("Rock")
		}
	} else if result == "Lose" {
		if opponentCh.symbol == "Rock" {
			return createOption("Scissor")
		}
		if opponentCh.symbol == "Paper" {
			return createOption("Rock")
		}
		if opponentCh.symbol == "Scissor" {
			return createOption("Paper")
		}
	}

	// Draw
	// return same as opponent
	return createOption(opponentCh.symbol)

}

func calculateScorePart2() int {
	totalScore := 0

	shapes := make(map[string]option)
	addAllOptions(shapes)

	results := make(map[string]string)
	possibleResult(results)

	outcomes := make(map[string]int)
	addOutcomes(outcomes)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Unable to open file: ", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		opponentCh := shapes[line[0]]
		result := results[line[1]]

		myChoice := determineMyChoice(opponentCh, result)
		score := myChoice.value + outcomes[result]
		totalScore += score
		// fmt.Println(opponentCh, myChoice, result)
	}

	return totalScore
}
