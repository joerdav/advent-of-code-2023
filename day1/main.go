package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"

	"github.com/joerdav/advent-of-code-2023/display"
)

var (
	//go:embed input.txt
	input string
	//go:embed test.txt
	test_input string
	//go:embed test2.txt
	test_input2 string
)

func main() {
	display.Print(1, 1, test_input, input, part1)
	display.Print(1, 2, test_input2, input, part2)
}

func part1(input string) string {
	var result int
	for _, line := range strings.Split(input, "\n") {
		for _, r := range line {
			if unicode.IsDigit(r) {
				result += int((r - '0') * 10)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			r := []rune(line)[i]
			if unicode.IsDigit(r) {
				result += int((r - '0'))
				break
			}
		}

	}
	return fmt.Sprint(result)
}

func part2(input string) string {
	var result int
	for _, line := range strings.Split(input, "\n") {
		for i := 0; i < len(line); i++ {
			r := []rune(line)[i]
			if unicode.IsDigit(r) {
				result += int((r - '0') * 10)
				break
			}
			if d := endsWithDigit(line[:i+1]); d != 0 {
				result += d*10
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			r := []rune(line)[i]
			if unicode.IsDigit(r) {
				result += int((r - '0'))
				break
			}
			if d := startsWithDigit(line[i:]); d != 0 {
				result += d
				break
			}
		}

	}
	return fmt.Sprint(result)
}

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func endsWithDigit(s string) int {
	for i, d := range digits {
		if strings.HasSuffix(s, d) {
			return i + 1
		}
	}
	return 0
}
func startsWithDigit(s string) int {
	for i, d := range digits {
		if strings.HasPrefix(s, d) {
			return i + 1
		}
	}
	return 0
}
