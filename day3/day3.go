package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
}

func part1() {
	input := getInput()

	for _, bag := range input {
		var l = len([]rune(bag))
		c1, c2 := bag[0:l/2], bag[l/2:l]
		same := getStringSame(c1, c2)

		fmt.Println(same)
	}
}

func getInput() []string {
	input, _ := os.ReadFile("./test_input")

	return strings.Split(string(input), "\r\n")
}

func getStringSame(s1 string, s2 string) []rune {
	same := []rune{}

	for _, r := range s1 {
		if strings.ContainsRune(s2, r) {
			same = append(same, r)
		}
	}

	return same
}
