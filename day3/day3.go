package main

import (
	"aoc2022/utils"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var lowerOffset = 96
var capitalOffset = 38

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	input := getInput()
	var totalSum = 0

	for _, bag := range input {
		var l = len([]rune(bag))
		c1, c2 := bag[0:l/2], bag[l/2:l]
		same := getStringSame([]string{c1, c2})
		totalSum += getSum(same)
	}
	return totalSum
}

func part2() int {
	input := getInput()
	var totalSum = 0

	for i := 0; i < len(input); i += 3 {
		same := getStringSame(input[i : i+3])
		totalSum += getSum(same)

	}

	return totalSum
}

func getStringSame(ss []string) []rune {
	same := []rune{}
	s1, rest := ss[0], ss[1:]

	for _, r := range s1 {
		allContains := []bool{}
		for _, s := range rest {
			allContains = append(allContains, strings.ContainsRune(s, r))
		}

		if utils.AllTrue(allContains) && !utils.SliceContains(same, r) {
			same = append(same, r)
		}
	}

	return same
}

func getSum(items []rune) int {
	sum := 0

	for _, r := range items {
		if unicode.IsUpper(r) {
			sum += int(r) - capitalOffset
		} else {
			sum += int(r) - lowerOffset
		}
	}

	return sum
}

func getInput() []string {
	input, _ := os.ReadFile("./input")

	return strings.Split(string(input), "\r\n")
}
