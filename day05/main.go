package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := readLines("input.txt")

	{
		nice := 0
		for _, line := range lines {
			vowels := 0
			for _, char := range line {
				if strings.ContainsRune("aeiou", char) {
					vowels++
				}
			}
			if vowels < 3 {
				continue
			}

			twice := 0
			for i := 0; i+1 < len(line); i++ {
				if line[i] == line[i+1] {
					twice++
				}
			}
			if twice == 0 {
				continue
			}

			if strings.Contains(line, "ab") || strings.Contains(line, "cd") || strings.Contains(line, "pq") || strings.Contains(line, "xy") {
				continue
			}

			nice++
		}

		fmt.Println("--- Part One ---")
		fmt.Println(nice)
	}

	{
		nice := 0
		for _, line := range lines {
			twice := 0
			for i := 0; i+1 < len(line); i++ {
				for j := i + 2; j+1 < len(line); j++ {
					if line[j] == line[i] && line[j+1] == line[i+1] {
						twice++
					}
				}
			}
			if twice == 0 {
				continue
			}

			repeat := 0
			for i := 0; i+2 < len(line); i++ {
				if line[i] == line[i+2] {
					repeat++
				}
			}
			if repeat == 0 {
				continue
			}

			nice++

		}

		fmt.Println("--- Part Two ---")
		fmt.Println(nice)
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
