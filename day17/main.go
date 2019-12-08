package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	capacities := readNumbers("input.txt")

	{
		fmt.Println("--- Part One ---")
		fmt.Println(countCombinations(capacities, 150))
	}

	{
		fmt.Println("--- Part Two ---")
	}
}

func countCombinations(capacities []int, volume int) int {
	if len(capacities) == 0 {
		if volume == 0 {
			return 1
		} else {
			return 0
		}
	}
	combinations := countCombinations(capacities[1:], volume)
	if capacities[0] <= volume {
		combinations += countCombinations(capacities[1:], volume-capacities[0])
	}
	return combinations
}

func readNumbers(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, toInt(scanner.Text()))
	}
	return numbers
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
