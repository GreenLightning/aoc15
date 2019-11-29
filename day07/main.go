package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("input.txt")

	// We do not capture the notion of "no signal".
	// Instead the wires will provide garbage signals,
	// but this is irrelevant for the puzzle.
	wires := make(map[string]uint16)

	{
		fmt.Println("--- Part One ---")
		simulate(lines, wires)
		fmt.Println(wires["a"])
	}

	{
		fmt.Println("--- Part Two ---")

		// Set wire b to wire a and remove the instruction for wire b,
		// because we want to override its signal.
		wires["b"] = wires["a"]
		for index, line := range lines {
			if strings.HasSuffix(line, "-> b") {
				lines[index] = lines[len(lines)-1]
				lines = lines[:len(lines)-1]
				break
			}
		}

		simulate(lines, wires)
		fmt.Println(wires["a"])
	}
}

var (
	single = regexp.MustCompile(`^(NOT )?(\w+) -> (\w+)$`)
	double = regexp.MustCompile(`^(\w+) (AND|OR|LSHIFT|RSHIFT) (\w+) -> (\w+)$`)
)

func simulate(lines []string, wires map[string]uint16) {
	get := func(input string) uint16 {
		if signal, err := strconv.ParseUint(input, 10, 16); err == nil {
			return uint16(signal)
		}
		return wires[input]
	}

	for {
		changed := false

		for _, line := range lines {
			var signal uint16
			var target string

			if match := single.FindStringSubmatch(line); match != nil {
				invert := len(match[1]) != 0
				signal = get(match[2])
				target = match[3]

				if invert {
					signal = ^signal
				}

			} else if match := double.FindStringSubmatch(line); match != nil {
				left, op, right := get(match[1]), match[2], get(match[3])
				target = match[4]

				switch op {
				case "AND":
					signal = left & right
				case "OR":
					signal = left | right
				case "LSHIFT":
					signal = left << right
				case "RSHIFT":
					signal = left >> right
				}

			} else {
				panic(line)
			}

			if wires[target] != signal {
				wires[target] = signal
				changed = true
			}
		}

		if !changed {
			break
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
