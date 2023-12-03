package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/joerdav/advent-of-code-2023/display"
)

var (
	//go:embed input.txt
	input string
	//go:embed test.txt
	test_input string
)

func main() {
	display.Print(3, 1, test_input, input, part1)
	display.Print(3, 2, test_input, input, part2)
}

type coord struct {
	x, y int
}

func part1(input string) string {
	var symbols []coord
	nums := make(map[coord]*int)
	lines := strings.Split(input, "\n")
	width := len([]rune(lines[0]))
	for y, line := range lines {
		var currNum []rune
		numStart := 0
		for x, r := range line {
			if unicode.IsDigit(r) {
				if len(currNum) == 0 {
					numStart = x
				}
				currNum = append(currNum, r)
				continue
			}
			if len(currNum) != 0 {
				num, _ := strconv.Atoi(string(currNum))
				for i := numStart; i < x; i++ {
					nums[coord{i, y}] = &num
				}
				currNum = nil
			}
			if r == '.' {
				continue
			}
			symbols = append(symbols, coord{x, y})

		}
		if len(currNum) != 0 {
			num, _ := strconv.Atoi(string(currNum))
			for i := numStart; i < width; i++ {
				nums[coord{i, y}] = &num
			}
			currNum = nil
		}
	}
	seen := map[*int]bool{}
	var total int
	for _, s := range symbols {
		adj := []coord{
			{s.x - 1, s.y},     //u
			{s.x + 1, s.y},     //d
			{s.x, s.y - 1},     //l
			{s.x, s.y + 1},     //r
			{s.x - 1, s.y - 1}, //ul
			{s.x - 1, s.y + 1}, //ur
			{s.x + 1, s.y - 1}, //dl
			{s.x + 1, s.y + 1}, //dr
		}
		for _, co := range adj {
			if n, ok := nums[co]; ok && !seen[n] {
				total += *n
				seen[n] = true
			}
		}
	}
	return fmt.Sprint(total)
}

func part2(input string) string {
	var symbols []coord
	nums := make(map[coord]*int)
	lines := strings.Split(input, "\n")
	width := len([]rune(lines[0]))
	for y, line := range lines {
		var currNum []rune
		numStart := 0
		for x, r := range line {
			if unicode.IsDigit(r) {
				if len(currNum) == 0 {
					numStart = x
				}
				currNum = append(currNum, r)
				continue
			}
			if len(currNum) != 0 {
				num, _ := strconv.Atoi(string(currNum))
				for i := numStart; i < x; i++ {
					nums[coord{i, y}] = &num
				}
				currNum = nil
			}
			if r == '.' {
				continue
			}
			symbols = append(symbols, coord{x, y})

		}
		if len(currNum) != 0 {
			num, _ := strconv.Atoi(string(currNum))
			for i := numStart; i < width; i++ {
				nums[coord{i, y}] = &num
			}
			currNum = nil
		}
	}
	var total int
	for _, s := range symbols {
		seen := map[*int]bool{}
		foundNums := []int{}
		adj := []coord{
			{s.x - 1, s.y},     //u
			{s.x + 1, s.y},     //d
			{s.x, s.y - 1},     //l
			{s.x, s.y + 1},     //r
			{s.x - 1, s.y - 1}, //ul
			{s.x - 1, s.y + 1}, //ur
			{s.x + 1, s.y - 1}, //dl
			{s.x + 1, s.y + 1}, //dr
		}
		for _, co := range adj {
			if n, ok := nums[co]; ok && !seen[n] {
				seen[n] = true
				foundNums = append(foundNums, *n)
			}
		}
		if len(foundNums) == 2 {
			total += foundNums[0] * foundNums[1]
		}
	}
	return fmt.Sprint(total)
}
