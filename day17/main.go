package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	capacities := readNumbers("input.txt")

	// We store combinations as bitfields inside of uints, so we can have 31
	// containers max (the last bit is reserved for the past-the-end value).
	if len(capacities) >= 32 {
		panic("input too large")
	}

	end := uint(1) << uint(len(capacities))

	valid := 0
	minValid, minContainers := 0, math.MaxInt32
	for combination := uint(0); combination < end; combination++ {
		volume, containers := 0, 0
		for index, capacity := range capacities {
			if combination&(uint(1)<<uint(index)) != 0 {
				volume += capacity
				containers++
			}
		}
		if volume == 150 {
			valid++
			if containers < minContainers {
				minValid, minContainers = 1, containers
			} else if containers == minContainers {
				minValid++
			}
		}
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(valid)
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(minValid)
	}
}

func readNumbers(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, toInt(scanner.Text()))
	}
	return numbers
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
