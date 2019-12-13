package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type State struct {
	BossHitPoints int
	BossDamage    int

	PlayerHitPoints int
	PlayerMana      int

	ShieldTimer   int
	PoisonTimer   int
	RechargeTimer int

	TotalManaSpent int
}

type Spell struct {
	Name  string
	Cost  int
	Apply func(state *State)
}

func main() {
	lines := readLines("input.txt")

	var start State

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		trait, value := parts[0], toInt(parts[1])
		switch trait {
		case "Hit Points":
			start.BossHitPoints = value
		case "Damage":
			start.BossDamage = value
		default:
			panic(fmt.Sprintf("unknown boss trait: %s", trait))
		}
	}

	start.PlayerHitPoints = 50
	start.PlayerMana = 500

	spells := []Spell{
		Spell{Name: "Magic Missile", Cost: 53, Apply: func(state *State) { state.BossHitPoints -= 4 }},
		Spell{Name: "Drain", Cost: 73, Apply: func(state *State) { state.BossHitPoints -= 2; state.PlayerHitPoints += 2 }},
		Spell{Name: "Shield", Cost: 113, Apply: func(state *State) { state.ShieldTimer = 6 }},
		Spell{Name: "Poison", Cost: 173, Apply: func(state *State) { state.PoisonTimer = 6 }},
		Spell{Name: "Recharge", Cost: 229, Apply: func(state *State) { state.RechargeTimer = 5 }},
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(findLeastAmountOfManaRequired(start, spells, false))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(findLeastAmountOfManaRequired(start, spells, true))
	}
}

func findLeastAmountOfManaRequired(start State, spells []Spell, hard bool) int {
	var queue PriorityQueue
	queue.Push(start)

	for !queue.Empty() {
		state := queue.Pop()
		for _, spell := range spells {
			result, newState := simulate(state, spell, hard)
			switch result {
			case ResultNone:
				queue.Push(newState)
			case ResultWon:
				return newState.TotalManaSpent
			case ResultLost:
				// discard state
			}
		}
	}

	return -1
}

const (
	ResultNone = 0
	ResultWon  = 1
	ResultLost = 2
)

func simulate(state State, spell Spell, hard bool) (int, State) {
	// Player turn.

	if hard {
		state.PlayerHitPoints--
		if state.PlayerHitPoints <= 0 {
			return ResultLost, state
		}
	}

	applyEffects(&state)

	if state.BossHitPoints <= 0 {
		return ResultWon, state
	}

	// If you cannot afford to cast any spell, you lose.
	if spell.Cost > state.PlayerMana {
		return ResultLost, state
	}

	// You cannot cast a spell that would start an effect which is already active.
	if spell.Name == "Shield" && state.ShieldTimer > 0 {
		return ResultLost, state
	}
	if spell.Name == "Poison" && state.PoisonTimer > 0 {
		return ResultLost, state
	}
	if spell.Name == "Recharge" && state.RechargeTimer > 0 {
		return ResultLost, state
	}

	state.PlayerMana -= spell.Cost
	state.TotalManaSpent += spell.Cost
	spell.Apply(&state)

	if state.BossHitPoints <= 0 {
		return ResultWon, state
	}

	// Boss turn.

	applyEffects(&state)

	if state.BossHitPoints <= 0 {
		return ResultWon, state
	}

	playerArmor := 0
	if state.ShieldTimer > 0 {
		playerArmor = 7
	}

	damage := max(state.BossDamage-playerArmor, 1)
	state.PlayerHitPoints -= damage

	if state.PlayerHitPoints <= 0 {
		return ResultLost, state
	}

	return ResultNone, state
}

func applyEffects(state *State) {
	if state.ShieldTimer > 0 {
		state.ShieldTimer--
	}

	if state.PoisonTimer > 0 {
		state.PoisonTimer--
		state.BossHitPoints -= 3
	}

	if state.RechargeTimer > 0 {
		state.RechargeTimer--
		state.PlayerMana += 101
	}
}

type PriorityStorage []State

func (s PriorityStorage) Len() int {
	return len(s)
}

func (s PriorityStorage) Less(i, j int) bool {
	return s[i].TotalManaSpent < s[j].TotalManaSpent
}

func (s PriorityStorage) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *PriorityStorage) Push(x interface{}) {
	state := x.(State)
	*s = append(*s, state)
}

func (s *PriorityStorage) Pop() interface{} {
	len := len(*s)
	state := (*s)[len-1]
	*s = (*s)[:len-1]
	return state
}

type PriorityQueue struct {
	storage PriorityStorage
}

func (q *PriorityQueue) Empty() bool {
	return len(q.storage) == 0
}

func (q *PriorityQueue) Push(state State) {
	heap.Push(&q.storage, state)
}

func (q *PriorityQueue) Pop() State {
	return heap.Pop(&q.storage).(State)
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
