package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lines := readLines("input.txt")
	regex := regexp.MustCompile(`^(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+)\.$`)

	attendees := make(map[string]bool)
	rules := make(map[string]int)

	for _, line := range lines {
		match := regex.FindStringSubmatch(line)
		one, direction, amount, two := match[1], match[2], toInt(match[3]), match[4]

		if direction == "lose" {
			amount = -amount
		}

		attendees[one] = true
		attendees[two] = true

		rules[fmt.Sprintf("%s-%s", one, two)] = amount
	}

	var attendeeList []string
	for attendee := range attendees {
		attendeeList = append(attendeeList, attendee)
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(findBestScore(attendeeList, rules))
	}

	{
		fmt.Println("--- Part Two ---")
		attendeeList = append(attendeeList, "")
		fmt.Println(findBestScore(attendeeList, rules))
	}
}

func findBestScore(attendees []string, rules map[string]int) int {
	bestScore := 0
	arrangements := allPermutations(attendees)
	for _, arrangement := range arrangements {
		score := 0
		for i := 0; i < len(arrangement); i++ {
			current := arrangement[i]
			next := arrangement[(i+1)%len(arrangement)]
			score += rules[fmt.Sprintf("%s-%s", current, next)]
			score += rules[fmt.Sprintf("%s-%s", next, current)]
		}
		bestScore = max(bestScore, score)
	}
	return bestScore
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

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
