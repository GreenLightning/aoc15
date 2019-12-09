package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readLines("input.txt")

	{
		fmt.Println("--- Part One ---")
		fmt.Println(animate(lines, false))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(animate(lines, true))
	}
}

func animate(lines []string, on bool) int {
	var storage [2][1 + 100 + 1][1 + 100 + 1]byte
	grid, next := &storage[0], &storage[1]

	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				grid[1+y][1+x] = 1
			}
		}
		if on {
			grid[1][1] = 1
			grid[1][100] = 1
			grid[100][1] = 1
			grid[100][100] = 1
		}
	}

	for step := 0; step < 100; step++ {
		for y := 1; y <= 100; y++ {
			for x := 1; x <= 100; x++ {
				top := grid[y-1][x-1] + grid[y-1][x] + grid[y-1][x+1]
				mid := grid[y][x-1] + grid[y][x+1]
				bot := grid[y+1][x-1] + grid[y+1][x] + grid[y+1][x+1]
				neighbors := top + mid + bot
				next[y][x] = 0
				if (grid[y][x] == 0 && neighbors == 3) || (grid[y][x] == 1 && neighbors >= 2 && neighbors <= 3) {
					next[y][x] = 1
				}
			}
		}
		if on {
			next[1][1] = 1
			next[1][100] = 1
			next[100][1] = 1
			next[100][100] = 1
		}
		grid, next = next, grid
	}

	total := 0
	for y := 1; y <= 100; y++ {
		for x := 1; x <= 100; x++ {
			total += int(grid[y][x])
		}
	}
	return total
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
