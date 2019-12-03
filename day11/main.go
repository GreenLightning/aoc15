package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input := readFile("input.txt")
	password := []byte(input)

	{
		fmt.Println("--- Part One ---")
		findNext(password)
		fmt.Println(string(password))
	}

	{
		fmt.Println("--- Part Two ---")
		findNext(password)
		fmt.Println(string(password))
	}
}

func findNext(password []byte) {
	for {
		incrementOnce(password)
		incrementPastInvalidCharacters(password)

		if !hasIncreasingStraight(password) {
			continue
		}

		first := findPair(password)
		if first == -1 {
			continue
		}

		second := findPair(password[first+2:])
		if second == -1 {
			continue
		}

		return
	}
}

func incrementOnce(password []byte) {
	for i := len(password) - 1; i >= 0; i-- {
		if password[i] != 'z' {
			password[i]++
			break
		}
		password[i] = 'a'
	}
}

func incrementPastInvalidCharacters(password []byte) {
	for i := len(password) - 1; i >= 0; i-- {
		for password[i] == 'i' || password[i] == 'o' || password[i] == 'l' {
			password[i]++
			for j := i + 1; j < len(password); j++ {
				password[j] = 'a'
			}
		}
	}
}

func hasIncreasingStraight(password []byte) bool {
	for i := 0; i+2 < len(password); i++ {
		if password[i]+1 == password[i+1] && password[i]+2 == password[i+2] {
			return true
		}
	}
	return false
}

func findPair(password []byte) int {
	for i := 0; i+1 < len(password); i++ {
		if password[i] == password[i+1] {
			return i
		}
	}
	return -1
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
