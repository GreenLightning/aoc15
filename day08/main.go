package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines := readLines("input.txt")

	metas, literals, values := 0, 0, 0
	for _, literal := range lines {
		meta := strconv.Quote(literal)
		value, err := strconv.Unquote(literal)
		check(err)
		metas += len(meta)
		literals += len(literal)
		values += len(value)
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(literals - values)
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(metas - literals)
	}
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
