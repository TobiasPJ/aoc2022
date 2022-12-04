package main

import (
	"aoc2022/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	input := getInput()
	total := 0

	for _, row := range input {
		assigns := strings.Split(row, ",")
		a1, a2 := createRange(assigns[0]), createRange(assigns[1])

		if (len(a1) >= len(a2) && checkAllIn(a2, a1)) || (len(a1) < len(a2) && checkAllIn(a1, a2)) {
			total += 1
		}
	}

	return total
}

func part2() int {
	input := getInput()
	total := 0

	for _, row := range input {
		assigns := strings.Split(row, ",")
		a1, a2 := createRange(assigns[0]), createRange(assigns[1])

		if (len(a1) >= len(a2) && checkAnyIn(a2, a1)) || (len(a1) < len(a2) && checkAnyIn(a1, a2)) {
			total += 1
		}
	}

	return total
}

func checkAllIn(s1 string, s2 string) bool {
	bools := []bool{}
	for _, r := range s1 {
		bools = append(bools, strings.ContainsRune(s2, r))
	}

	return utils.AllTrue(bools)
}

func checkAnyIn(s1 string, s2 string) bool {
	for _, r := range s1 {
		if strings.ContainsRune(s2, r) {
			return true
		}
	}

	return false
}

func createRange(s string) string {
	r := strings.Split(s, "-")
	r1, _ := strconv.ParseInt(r[0], 10, 64)
	r2, _ := strconv.ParseInt(r[1], 10, 64)

	ns := []rune{}
	for i := r1; i <= r2; i++ {
		ns = append(ns, rune(i))
	}

	return string(ns)
}

func getInput() []string {
	input, _ := os.ReadFile("./input")

	return strings.Split(string(input), "\r\n")
}
