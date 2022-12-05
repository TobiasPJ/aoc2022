package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	func1()

}
func func1() string {
	input := getInput()
	columns, _ := strings.Split(input[0], "\r\n"), input[1]

	fmt.Println(columns)
	return ""
}

func getColumns(c string) []string {
	columns := strings.Split(c, "\r\n")

	ret := []string{}
	for _, c := range columns {
		replacer := strings.NewReplacer("[", "", "]", "")
		ret = append(ret, replacer.Replace(c))
		fmt.Println(len(replacer.Replace(c)))
	}

	return ret
}

func getInput() []string {
	input, _ := os.ReadFile("./test_input")

	return strings.Split(string(input), "\r\n\r\n")
}
