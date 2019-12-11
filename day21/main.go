package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Character struct {
	HitPoints int
	Damage    int
	Armor     int
}

type Item struct {
	Name   string
	Cost   int
	Damage int
	Armor  int
}

func main() {
	lines := readLines("input.txt")

	var enemy Character

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		trait, value := parts[0], toInt(parts[1])
		switch trait {
		case "Hit Points":
			enemy.HitPoints = value
		case "Damage":
			enemy.Damage = value
		case "Armor":
			enemy.Armor = value
		default:
			panic(fmt.Sprintf("unknown enemy trait: %s", trait))
		}
	}

	weapons := []Item{
		Item{"Dagger", 8, 4, 0},
		Item{"Shortsword", 10, 5, 0},
		Item{"Warhammer", 25, 6, 0},
		Item{"Longsword", 40, 7, 0},
		Item{"Greataxe", 74, 8, 0},
	}

	// Armor is optional, so we include a "No Armor" option.
	armors := []Item{
		Item{"No Armor", 0, 0, 0},
		Item{"Leather", 13, 0, 1},
		Item{"Chainmail", 31, 0, 2},
		Item{"Splintmail", 53, 0, 3},
		Item{"Bandedmail", 75, 0, 4},
		Item{"Platemail", 102, 0, 5},
	}

	// You cannot buy the same item twice, so we have to include two "No Ring"
	// options, one for each hand, in case you want to buy no rings at all.
	rings := []Item{
		Item{"No Ring", 0, 0, 0},
		Item{"No Ring", 0, 0, 0},
		Item{"Damage +1", 25, 1, 0},
		Item{"Damage +2", 50, 2, 0},
		Item{"Damage +3", 100, 3, 0},
		Item{"Defense +1", 20, 0, 1},
		Item{"Defense +2", 40, 0, 2},
		Item{"Defense +3", 80, 0, 3},
	}

	minCost, maxCost := math.MaxInt32, 0
	for _, weapon := range weapons {
		for _, armor := range armors {
			for leftIndex, leftRing := range rings {
				for rightIndex, rightRing := range rings {
					if rightIndex == leftIndex {
						continue
					}

					cost := weapon.Cost + armor.Cost + leftRing.Cost + rightRing.Cost

					var player Character
					player.HitPoints = 100
					player.Damage = weapon.Damage + armor.Damage + leftRing.Damage + rightRing.Damage
					player.Armor = weapon.Armor + armor.Armor + leftRing.Armor + rightRing.Armor

					// Characters always deal the same amount of damage each round,
					// because the damage only depends on the fixed stats.
					playerRoundDamge := max(player.Damage-enemy.Armor, 1)
					enemyRoundDamage := max(enemy.Damage-player.Armor, 1)

					// Therefore we can calculate how many rounds each
					// character needs to kill the other one.
					playerRounds := (enemy.HitPoints + playerRoundDamge - 1) / playerRoundDamge
					enemyRounds := (player.HitPoints + enemyRoundDamage - 1) / enemyRoundDamage

					// Since the player goes first, they also win if both
					// opponents need the same amount of rounds.
					if playerRounds <= enemyRounds {
						minCost = min(minCost, cost)
					} else {
						maxCost = max(maxCost, cost)
					}
				}
			}
		}
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(minCost)
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(maxCost)
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

func min(x, y int) int {
	if y < x {
		return y
	}
	return x
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
