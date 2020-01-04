package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readFile("input.txt")

	regex := regexp.MustCompile(`^To continue, please consult the code grid in the manual.  Enter the code at row (\d+), column (\d+).$`)

	match := regex.FindStringSubmatch(input)

	row := toInt(match[1])
	col := toInt(match[2])

	count := (row+col-1)*(row+col-2)/2 + col

	code := 20151125
	for iteration := 1; iteration < count; iteration++ {
		code = (code * 252533) % 33554393
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(code)
	}
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
