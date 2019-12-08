package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Aunt map[string]int

func main() {
	lines := readLines("input.txt")

	aunts := make(map[int]Aunt)

	lineRegex := regexp.MustCompile(`^Sue (\d+): (.*)$`)
	infoRegex := regexp.MustCompile(`^(\w+): (\d+)$`)

	for _, line := range lines {
		match := lineRegex.FindStringSubmatch(line)
		number, infos := toInt(match[1]), match[2]

		aunt := make(Aunt)

		for _, info := range strings.Split(infos, ", ") {
			match := infoRegex.FindStringSubmatch(info)
			thing, amount := match[1], toInt(match[2])
			aunt[thing] = amount
		}

		aunts[number] = aunt
	}

	results := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	{
		fmt.Println("--- Part One ---")
	partOne:
		for number, aunt := range aunts {
			for thing, amount := range aunt {
				if amount != results[thing] {
					continue partOne
				}
			}
			fmt.Println(number)
		}
	}

	{
		fmt.Println("--- Part Two ---")
	partTwo:
		for number, aunt := range aunts {
			for thing, amount := range aunt {
				switch thing {
				case "cats", "trees":
					if amount <= results[thing] {
						continue partTwo
					}
				case "pomeranians", "goldfish":
					if amount >= results[thing] {
						continue partTwo
					}
				default:
					if amount != results[thing] {
						continue partTwo
					}
				}
			}
			fmt.Println(number)
		}
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
