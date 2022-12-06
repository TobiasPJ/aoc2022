package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	return solution(3)
}

func part2() int {
	return solution(13)
}

func solution(start int) int {
	input := getInput()

	for i := start; i < len([]rune(input)); i++ {
		look_back := input[i-start : i+1]

		if allUnique(look_back) {
			return i + 1
		}
	}

	return 0
}

func allUnique(s string) bool {
	for i, r := range s {
		if strings.ContainsRune(s[i+1:], r) {
			return false
		}
	}

	return true
}

func getInput() string {
	input, _ := os.ReadFile("./input")

	return string(input)
}
