package main

import (
	_ "embed"
	"fmt"
	"strconv"
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
	display.Print(9, 1, test_input, input, part1)
	display.Print(9, 2, test_input, input, part2)
}

func parseLine(line string) []int {
	numstrings := strings.Fields(line)
	var nms []int
	for _, n := range numstrings {
		nu, _ := strconv.Atoi(n)
		nms = append(nms, nu)
	}
	return nms
}

func getDiffs(nums []int, diffs [][]int) [][]int {
	var diff []int
	var hasNoneZero bool
	for i := 0; i < len(nums)-1; i++ {
		n := nums[i+1]-nums[i]
		if n != 0 {
			hasNoneZero = true
		}
		diff = append(diff, n)
	}
	diffs = append(diffs, diff)
	if hasNoneZero {
		return getDiffs(diff, diffs)
	}
	return diffs
}

func part1(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var total int
	for _, l := range lines {
		nums := parseLine(l)
		diffs := getDiffs(nums, [][]int{nums})
		end := 0
		for i := len(diffs)-1; i >=0; i-- {
			end += diffs[i][len(diffs[i])-1]
		}
		total += end
	}
	return fmt.Sprint(total)
}

func part2(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var total int
	for _, l := range lines {
		nums := parseLine(l)
		diffs := getDiffs(nums, [][]int{nums})
		end := 0
		for i := len(diffs)-1; i >=0; i-- {
			end = diffs[i][0] - end
		}
		total += end
	}
	return fmt.Sprint(total)
}
