package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/joerdav/advent-of-code-2023/display"
)

var (
	//go:embed input.txt
	input string
	//go:embed test.txt
	test_input string
)

func main() {
	display.Print(14, 1, test_input, input, part1)
	display.Print(14, 2, test_input, input, part2)
}

func part1(input string) string {
	var total int
	lines := strings.Split(strings.TrimSpace(input), "\n")
	lastRocks := make([]int, len(lines[0]))
	maxLoad := len(lines)
	for r, l := range lines {
		for c, ch := range l {
			if ch == '#' {
				lastRocks[c] = r + 1
				continue
			}
			if ch == 'O' {
				total += (maxLoad - lastRocks[c])
				lastRocks[c]++
			}
		}
	}
	return fmt.Sprint(total)
}

type cell int

const (
	empty cell = iota
	square
	rounded
)

func part2(input string) string {
	var total int
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := [][]cell{}
	for _, l := range lines {
		row := make([]cell, len(l))
		for c, ch := range l {
			if ch == '#' {
				row[c] = square
				continue
			}
			if ch == 'O' {
				row[c] = rounded
				continue
			}
		}
		grid = append(grid, row)
	}
	return fmt.Sprint(total)
}
