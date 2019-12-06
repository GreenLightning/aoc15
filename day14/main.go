package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Reindeer struct {
	Name        string
	Speed       int
	FlyingTime  int
	RestingTime int

	CurrentlyFlying bool
	CurrentDistance int
	CurrentPoints   int
	RemainingTime   int
}

func main() {
	lines := readLines("input.txt")

	var reindeer []*Reindeer

	regex := regexp.MustCompile(`^(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds\.$`)

	for _, line := range lines {
		match := regex.FindStringSubmatch(line)
		var r Reindeer
		r.Name = match[1]
		r.Speed = toInt(match[2])
		r.FlyingTime = toInt(match[3])
		r.RestingTime = toInt(match[4])

		r.CurrentlyFlying = true
		r.RemainingTime = r.FlyingTime

		reindeer = append(reindeer, &r)
	}

	for second := 1; second <= 2503; second++ {
		for _, r := range reindeer {
			if r.CurrentlyFlying {
				r.CurrentDistance += r.Speed
			}
			r.RemainingTime -= 1
			if r.RemainingTime <= 0 {
				r.CurrentlyFlying = !r.CurrentlyFlying
				if r.CurrentlyFlying {
					r.RemainingTime = r.FlyingTime
				} else {
					r.RemainingTime = r.RestingTime
				}
			}
		}

		bestDistance := bestDistance(reindeer)
		for _, r := range reindeer {
			if r.CurrentDistance == bestDistance {
				r.CurrentPoints++
			}
		}
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(bestDistance(reindeer))
	}

	{
		fmt.Println("--- Part Two ---")
		var bestPoints int
		for _, r := range reindeer {
			bestPoints = max(bestPoints, r.CurrentPoints)
		}
		fmt.Println(bestPoints)
	}
}

func bestDistance(reindeer []*Reindeer) (bestDistance int) {
	for _, r := range reindeer {
		bestDistance = max(bestDistance, r.CurrentDistance)
	}
	return
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
