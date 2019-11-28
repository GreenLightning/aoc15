package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Position struct {
	x, y int
}

func (p *Position) Move(dir rune) {
	switch dir {
	case '^':
		p.y++
	case 'v':
		p.y--
	case '>':
		p.x++
	case '<':
		p.x--
	}
}

func main() {
	input := readFile("input.txt")

	{
		pos := Position{0, 0}

		visited := make(map[Position]bool)
		visited[pos] = true

		for _, dir := range input {
			pos.Move(dir)
			visited[pos] = true
		}

		fmt.Println("--- Part One ---")
		fmt.Println(len(visited))
	}

	{
		santa, robo := Position{0, 0}, Position{0, 0}

		visited := make(map[Position]bool)
		visited[Position{0, 0}] = true

		for i, dir := range input {
			if i%2 == 0 {
				santa.Move(dir)
				visited[santa] = true
			} else {
				robo.Move(dir)
				visited[robo] = true
			}
		}

		fmt.Println("--- Part Two ---")
		fmt.Println(len(visited))
	}
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
