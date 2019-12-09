package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	lines := readLines("input.txt")

	input := lines[len(lines)-1]
	lines = lines[:len(lines)-1]

	rules := make(map[string][]string)

	regex := regexp.MustCompile(`^(\w+) => (\w+)$`)
	for _, line := range lines {
		if line != "" {
			match := regex.FindStringSubmatch(line)
			rules[match[1]] = append(rules[match[1]], match[2])
		}
	}

	{
		fmt.Println("--- Part One ---")

		results := make(map[string]bool)

		for from, replacements := range rules {
			base := 0
			for {
				index := strings.Index(input[base:], from)
				if index == -1 {
					break
				}
				index += base

				for _, replacement := range replacements {
					result := input[:index] + replacement + input[index+len(from):]
					results[result] = true
				}

				base = index + len(from)
			}
		}

		fmt.Println(len(results))
	}

	{
		fmt.Println("--- Part Two ---")

		count := func(input string, pattern string) int {
			return len(regexp.MustCompile(pattern).FindAllString(input, -1))
		}

		// This calculation requires that the following assumptions hold for the replacement rules.
		//
		// An element name is a single uppercase letter followed by zero or more lowercase letters.
		// The target molecule contains only element names (i.e. no numbers and no electrons; each
		// atom is described by its element name).
		//
		// In general each rule replaces one thing (atom or electron) with two atoms, so we start by
		// counting the number of atoms, which, according to the assumptions above, is equal to the
		// number of uppercase letters.
		//
		// A rule can insert additional Ar or Rn atoms. Since these atoms can be inserted without
		// performing an additional replacement, we subtract them from the count.
		//
		// A rule can further insert any additional atom as long as a Y atom is inserted as well.
		// Therefore we subtract the number of Y atoms twice (once for the Y itself and once for the
		// additional atom).
		//
		// For this to work, there must be no rules that can replace an Ar, Rn or Y atom, so that
		// they all appear in the target molecule.
		//
		// Finally, we start with one electron, for which no replacement is necessary, so we
		// subtract one from the count.
		fmt.Println(count(input, "[A-Z]") - count(input, "Ar|Rn") - 2*count(input, "Y") - 1)
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
