package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := toInt(readFile("input.txt"))

	{
		fmt.Println("--- Part One ---")
		for house := 1; ; house++ {
			score, elf := 0, 1
			for ; elf*elf < house; elf++ {
				if house%elf == 0 {
					score += elf + house/elf
				}
			}
			if elf*elf == house {
				score += elf
			}
			score *= 10
			if score >= input {
				fmt.Println(house)
				break
			}
		}
	}

	{
		fmt.Println("--- Part Two ---")
		for house := 1; ; house++ {
			score, elf := 0, 1
			for ; elf*elf < house; elf++ {
				if house%elf == 0 {
					factor := house / elf
					if factor <= 50 {
						score += elf
					}
					if elf <= 50 {
						score += factor
					}
				}
			}
			if elf*elf == house && elf <= 50 {
				score += elf
			}
			score *= 11
			if score >= input {
				fmt.Println(house)
				break
			}
		}
	}
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
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
