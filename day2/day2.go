package main

import (
	"fmt"
	"os"
	"strings"
)

var scoreMap = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}
var winMap = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}
var tieMap = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}
var loseMap = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}
var lostScore = 0
var winScore = 6
var tieScore = 3

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	input := getInput()

	var score = 0
	for _, round := range input {
		moves := strings.Split(round, " ")
		opponent, me := moves[0], moves[1]
		if tieMap[opponent] == me {
			score += scoreMap[me] + tieScore
		} else if winMap[opponent] == me {
			score += scoreMap[me] + winScore
		} else {
			score += scoreMap[me] + lostScore
		}
	}

	return score
}

func part2() int {
	input := getInput()
	var score = 0

	for _, round := range input {
		moves := strings.Split(round, " ")
		opponent, me := moves[0], moves[1]
		switch me {
		case "Y":
			score += scoreMap[tieMap[opponent]] + tieScore
		case "X":
			score += scoreMap[loseMap[opponent]] + lostScore
		case "Z":
			score += scoreMap[winMap[opponent]] + winScore
		}
	}

	return score
}

func getInput() []string {
	input, _ := os.ReadFile("./input")

	return strings.Split(string(input), "\r\n")
}
