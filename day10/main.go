package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := readFile("input.txt")

	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(len(input))
	}

	for i := 0; i < 10; i++ {
		input = lookAndSay(input)
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(len(input))
	}
}

func lookAndSay(text string) string {
	var result strings.Builder
	for i := 0; i < len(text); {
		digit, length := text[i], 0
		for ; i < len(text) && text[i] == digit; i++ {
			length++
		}
		result.WriteString(strconv.FormatInt(int64(length), 10))
		result.WriteByte(digit)
	}
	return result.String()
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
