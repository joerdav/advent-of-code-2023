package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"

	"github.com/joerdav/advent-of-code-2023/display"
	"github.com/joerdav/advent-of-code-2023/iter"
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
	lines := iter.StringSplit(input, "\n")
	result := iter.Acc(lines, 0,
		func(line string, a int) int {
			if d, ok := iter.First(iter.New([]rune(line)), unicode.IsDigit); ok {
				a += int((d - '0') * 10)
			}
			if d, ok := iter.Last(iter.New([]rune(line)), unicode.IsDigit); ok {
				a += int(d - '0')
			}
			return a
		})
	return fmt.Sprint(result)
}

func part2(input string) string {
	lines := iter.StringSplit(input, "\n")
	result := iter.Acc(lines, 0, func(line string, a int) int {
		for i := 0; i < len(line); i++ {
			r := []rune(line)[i]
			if unicode.IsDigit(r) {
				a += int((r - '0') * 10)
				break
			}
			if d := endsWithDigit(line[:i+1]); d != 0 {
				a += d * 10
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			r := []rune(line)[i]
			if unicode.IsDigit(r) {
				a += int((r - '0'))
				break
			}
			if d := startsWithDigit(line[i:]); d != 0 {
				a += d
				break
			}
		}
		return a
	})
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
