package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	weights := readNumbers("input.txt")

	total := 0
	for _, weight := range weights {
		total += weight
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(findQuantumEntanglementOfFirstGroup(weights, total/3))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(findQuantumEntanglementOfFirstGroup(weights, total/4))
	}
}

func findQuantumEntanglementOfFirstGroup(weights []int, target int) int {
	bestCount, bestEntanglement := math.MaxInt32, math.MaxInt32

	for _, group := range findGroups(weights, target) {
		count := len(group)

		entanglement := 1
		for _, weight := range group {
			entanglement *= weight
		}

		if count < bestCount || (count == bestCount && entanglement < bestEntanglement) {
			bestCount = count
			bestEntanglement = entanglement
		}
	}

	return bestEntanglement
}

func findGroups(weights []int, target int) [][]int {
	if len(weights) == 0 {
		return nil
	}
	groups := findGroups(weights[1:], target)
	if weights[0] == target {
		groups = append(groups, []int{weights[0]})
	} else if weights[0] < target {
		for _, group := range findGroups(weights[1:], target-weights[0]) {
			group = append(group, weights[0])
			groups = append(groups, group)
		}
	}
	return groups
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
