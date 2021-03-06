package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lines := readLines("input.txt")

	locations := make(map[string]bool)
	distances := make(map[string]int)

	regex := regexp.MustCompile(`^(\w+) to (\w+) = (\d+)$`)
	for _, line := range lines {
		match := regex.FindStringSubmatch(line)
		locations[match[1]] = true
		locations[match[2]] = true
		distances[fmt.Sprintf("%s-%s", match[1], match[2])] = toInt(match[3])
		distances[fmt.Sprintf("%s-%s", match[2], match[1])] = toInt(match[3])
	}

	var locationsList []string
	for location := range locations {
		locationsList = append(locationsList, location)
	}

	routes := allPermutations(locationsList)

	shortest, longest := math.MaxInt32, 0
	for _, route := range routes {
		distance := calculateDistance(route, distances)
		shortest = min(distance, shortest)
		longest = max(distance, longest)
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(shortest)
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(longest)
	}
}

func allPermutations(values []string) (result [][]string) {
	if len(values) == 1 {
		result = append(result, values)
		return
	}
	for i, current := range values {
		others := make([]string, 0, len(values)-1)
		others = append(others, values[:i]...)
		others = append(others, values[i+1:]...)
		for _, route := range allPermutations(others) {
			result = append(result, append(route, current))
		}
	}
	return
}

func calculateDistance(route []string, distances map[string]int) (distance int) {
	for i := 0; i+1 < len(route); i++ {
		distance += distances[fmt.Sprintf("%s-%s", route[i], route[i+1])]
	}
	return
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

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
