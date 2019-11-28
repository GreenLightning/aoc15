package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("input.txt")

	wrappingPaper, ribbon := 0, 0
	for _, line := range lines {
		parts := strings.Split(line, "x")
		x, y, z := toInt(parts[0]), toInt(parts[1]), toInt(parts[2])

		wrappingPaper += 2*x*y + 2*y*z + 2*z*x + min(x*y, min(y*z, z*x))
		ribbon += min(2*(x+y), min(2*(y+z), 2*(z+x))) + x*y*z
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(wrappingPaper)
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(ribbon)
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

func min(x, y int) int {
	if y < x {
		return y
	}
	return x
}
