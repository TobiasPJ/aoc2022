package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(func1())
	fmt.Println(func2())
}
func func1() string {
	input := getInput()
	columns, numberRow := getColumns(input[0])
	instructions := strings.Split(input[1], "\r\n")

	for _, ins := range instructions {
		amount, from, to := parseInstruction(ins)

		for _, box := range columns[from][0:amount] {
			columns[to] = string(box) + columns[to]
		}

		columns[from] = columns[from][amount:]
	}

	return getTops(columns, numberRow)
}

func func2() string {
	input := getInput()
	columns, numberRow := getColumns(input[0])
	instructions := strings.Split(input[1], "\r\n")

	for _, ins := range instructions {
		amount, from, to := parseInstruction(ins)
		columns[to] = columns[from][0:amount] + columns[to]
		columns[from] = columns[from][amount:]
	}

	return getTops(columns, numberRow)
}

func getColumns(c string) (map[string]string, string) {
	rows := strings.Split(c, "\r\n")
	rows, numberRow := rows[0:len(rows)-1], rows[len(rows)-1]
	columns := make(map[string]string)

	for i, number := range numberRow {
		if string(number) != " " {
			var t string
			for _, row := range rows {
				if string(row[i]) != " " {
					t = t + string(row[i])
				}
			}
			columns[string(number)] = t
		}
	}

	return columns, numberRow
}

func parseInstruction(ins string) (int, string, string) {
	v := strings.Split(strings.NewReplacer("move ", "", " from ", " ", " to ", " ").Replace(ins), " ")
	amount, _ := strconv.ParseInt(v[0], 10, 64)

	return int(amount), v[1], v[2]
}

func getTops(columns map[string]string, numberRow string) string {
	var tops string
	for _, n := range strings.ReplaceAll(numberRow, " ", "") {
		tops += string(columns[string(n)][0])
	}
	return tops
}

func getInput() []string {
	input, _ := os.ReadFile("./input")

	return strings.Split(string(input), "\r\n\r\n")
}
