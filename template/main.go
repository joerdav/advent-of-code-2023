package main

import (
	_ "embed"

	"github.com/joerdav/advent-of-code-2023/display"
)

var (
	//go:embed input.txt
	input string
	//go:embed test.txt
	test_input string
)

func main() {
	display.Print(1, 1, test_input, input, part1)
	display.Print(1, 2, test_input, input, part2)
}

func part1(input string) string {
	return "not implemented"
}

func part2(input string) string {
	return "not implemented"
}
