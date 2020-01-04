package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	HLF = iota
	TPL
	INC
	JMP
	JIE
	JIO
)

const (
	RegisterA = 0
	RegisterB = 1
)

type Instruction struct {
	Kind     int
	Register int
	Offset   int
}

func main() {
	lines := readLines("input.txt")

	var instructions []Instruction

	simpleRegex := regexp.MustCompile(`^(hlf|tpl|inc) (a|b)$`)
	jmpRegex := regexp.MustCompile(`^jmp ([+-]\d+)$`)
	jmpRegisterRegex := regexp.MustCompile(`^(jie|jio) (a|b), ([+-]\d+)$`)

	for _, line := range lines {
		var inst Instruction

		if match := simpleRegex.FindStringSubmatch(line); match != nil {
			switch match[1] {
			case "hlf":
				inst.Kind = HLF
			case "tpl":
				inst.Kind = TPL
			case "inc":
				inst.Kind = INC
			default:
				panic(match[1])
			}

			switch match[2] {
			case "a":
				inst.Register = RegisterA
			case "b":
				inst.Register = RegisterB
			default:
				panic(match[2])
			}

		} else if match := jmpRegex.FindStringSubmatch(line); match != nil {
			inst.Kind = JMP
			inst.Offset = toInt(match[1])

		} else if match := jmpRegisterRegex.FindStringSubmatch(line); match != nil {
			switch match[1] {
			case "jie":
				inst.Kind = JIE
			case "jio":
				inst.Kind = JIO
			default:
				panic(match[1])
			}

			switch match[2] {
			case "a":
				inst.Register = RegisterA
			case "b":
				inst.Register = RegisterB
			default:
				panic(match[2])
			}

			inst.Offset = toInt(match[3])

		} else {
			panic(line)
		}

		instructions = append(instructions, inst)
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(simulate(instructions, 0))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(simulate(instructions, 1))
	}
}

func simulate(instructions []Instruction, init int) int {
	var registers [2]int
	registers[RegisterA] = init

	ip := 0
	for ip >= 0 && ip < len(instructions) {
		inst := instructions[ip]
		switch inst.Kind {
		case HLF:
			registers[inst.Register] /= 2
			ip++

		case TPL:
			registers[inst.Register] *= 3
			ip++

		case INC:
			registers[inst.Register]++
			ip++

		case JMP:
			ip += inst.Offset

		case JIE:
			if registers[inst.Register]%2 == 0 {
				ip += inst.Offset
			} else {
				ip++
			}

		case JIO:
			if registers[inst.Register] == 1 {
				ip += inst.Offset
			} else {
				ip++
			}

		default:
			panic(inst.Kind)
		}
	}

	return registers[RegisterB]
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
