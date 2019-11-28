package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input := readFile("input.txt")

	floor, basementPosition := 0, -1
	for index, char := range input {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
		if basementPosition == -1 && floor == -1 {
			basementPosition = index + 1
		}
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(floor)
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(basementPosition)
	}
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
