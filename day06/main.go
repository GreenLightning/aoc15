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

	var binaryLights [1000][1000]bool
	var dimmableLights [1000][1000]int

	regex := regexp.MustCompile(`(toggle|turn on|turn off) (\d+),(\d+) through (\d+),(\d+)`)

	for _, line := range lines {
		match := regex.FindStringSubmatch(line)
		instruction := match[1]
		xmin, ymin := toInt(match[2]), toInt(match[3])
		xmax, ymax := toInt(match[4]), toInt(match[5])
		for y := ymin; y <= ymax; y++ {
			for x := xmin; x <= xmax; x++ {
				switch instruction {
				case "toggle":
					binaryLights[y][x] = !binaryLights[y][x]
					dimmableLights[y][x] += 2
				case "turn on":
					binaryLights[y][x] = true
					dimmableLights[y][x] += 1
				case "turn off":
					binaryLights[y][x] = false
					dimmableLights[y][x] = max(dimmableLights[y][x]-1, 0)
				}
			}
		}
	}

	{
		count := 0
		for y := 0; y < 1000; y++ {
			for x := 0; x < 1000; x++ {
				if binaryLights[y][x] {
					count++
				}
			}
		}

		fmt.Println("--- Part One ---")
		fmt.Println(count)
	}

	{
		brightness := 0
		for y := 0; y < 1000; y++ {
			for x := 0; x < 1000; x++ {
				brightness += dimmableLights[y][x]
			}
		}

		fmt.Println("--- Part Two ---")
		fmt.Println(brightness)
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

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
