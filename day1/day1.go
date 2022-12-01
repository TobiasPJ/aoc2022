package main

import (
	"aoc2022/utils"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int64 {
	input := getInput()
	sumArray := make([]int64, len(input))

	for i, nums := range input {
		sumArray[i] = utils.SumStringArray(strings.Split(nums, "\r\n"))
	}

	return utils.ArrayMax(sumArray)
}

func part2() int64 {
	input := getInput()
	sumArray := make([]int64, len(input))

	for i, nums := range input {
		sumArray[i] = utils.SumStringArray(strings.Split(nums, "\r\n"))
	}

	sort.Slice(sumArray, func(i, j int) bool {
		return sumArray[i] > sumArray[j]
	})

	return utils.SumInt64Array(sumArray[0:3])
}

func getInput() []string {
	data, _ := os.ReadFile("./input")

	return strings.Split(string(data), "\r\n\r\n")
}
