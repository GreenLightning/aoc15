package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Ingredient struct {
	capacity, durability, flavor, texture, calories int
}

func main() {
	lines := readLines("input.txt")

	regex := regexp.MustCompile(`^(\w+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)$`)

	var ingredients []Ingredient
	for _, line := range lines {
		match := regex.FindStringSubmatch(line)
		ingredients = append(ingredients, Ingredient{
			capacity:   toInt(match[2]),
			durability: toInt(match[3]),
			flavor:     toInt(match[4]),
			texture:    toInt(match[5]),
			calories:   toInt(match[6]),
		})
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(findBestScore(ingredients, nil, 100, 0))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(findBestScore(ingredients, nil, 100, 500))
	}
}

// The amounts slice must be shorter than the ingredients slice.
// Each entry in the amounts slice fixes the amount for the ingredient with the same index.
// The remaining amount will be distributed to the remaining ingredients in all possible ways
// and the best score will be returned.
func findBestScore(ingredients []Ingredient, amounts []int, remaining int, targetCalories int) int {
	// If there is only one ingredient left,
	// calculate the score for the assigned amounts.
	if len(amounts) == len(ingredients)-1 {

		// The last ingredient must take the remaining amount.
		amounts = append(amounts, remaining)

		var capacity, durability, flavor, texture, calories int

		for i := range ingredients {
			capacity += amounts[i] * ingredients[i].capacity
			durability += amounts[i] * ingredients[i].durability
			flavor += amounts[i] * ingredients[i].flavor
			texture += amounts[i] * ingredients[i].texture
			calories += amounts[i] * ingredients[i].calories
		}

		if capacity <= 0 || durability <= 0 || flavor <= 0 || texture <= 0 {
			return 0
		}

		if targetCalories != 0 && calories != targetCalories {
			return 0
		}

		return capacity * durability * flavor * texture
	}

	// Otherwise, we try all possible values for the next amount and keep track
	// of the best score.
	best := 0
	for value := 0; value <= remaining; value++ {
		score := findBestScore(ingredients, append(amounts, value), remaining-value, targetCalories)
		best = max(best, score)
	}
	return best
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
