package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
	"unicode"
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
	fmt.Println("1.1")
	start := time.Now()
	real1 := part1(input)
	duration := time.Since(start)
	fmt.Printf("  real: %s (%v)\n", real1, duration)
	fmt.Printf("  test: %s\n", part1(test_input))
	fmt.Println("1.2")
	start2 := time.Now()
	real2 := part2(input)
	duration2 := time.Since(start2)
	fmt.Printf("  real: %s (%v)\n", real2, duration2)
	fmt.Printf("  test: %s\n", part2(test_input2))

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

var digits = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func endsWithDigit(s string) int {
	for i, d := range digits {
		if strings.HasSuffix(s, d) {
			return i+1
		}
	}
	return 0
}
func startsWithDigit(s string) int {
	for i, d := range digits {
		if strings.HasPrefix(s, d) {
			return i+1
		}
	}
	return 0
}
