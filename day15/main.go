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
	display.Print(15, 1, test_input, input, part1)
	display.Print(15, 2, test_input, input, part2)
}

func hash(str string) int {
	var total int
	for _, r := range str {
		total += int(r)
		total *= 17
		total %= 256
	}
	return total
}

func part1(input string) string {
	var total int
	for _, l := range strings.Split(strings.TrimSpace(input), ",") {
		total += hash(l)
	}
	return fmt.Sprint(total)
}

type lense struct {
	label string
	fl    int
}

func addLense(lenses []lense, lense lense) []lense {
	for i := range lenses {
		if lenses[i].label == lense.label {
			lenses[i] = lense
			return lenses
		}
	}
	lenses = removeLense(lenses, lense.label)
	return append(lenses, lense)
}
func removeLense(lenses []lense, label string) []lense {
	var res []lense
	for _, l := range lenses {
		if l.label != label {
			res = append(res, l)
		}
	}
	return res
}

func part2(input string) string {
	var total int
	boxes := map[int][]lense{}
	for _, l := range strings.Split(strings.TrimSpace(input), ",") {
		label, fl, ok := strings.Cut(l, "=")
		label = strings.Trim(label, "-")
		box := hash(label)
		if !ok {
			boxes[box] = removeLense(boxes[box], label)
			continue
		}
		f, _ := strconv.Atoi(fl)
		boxes[box] = addLense(boxes[box], lense{label, f})
	}
	for i, box := range boxes {
		for p, l := range box {
			total += (1 + i) * (1 + p) * l.fl
		}
	}
	return fmt.Sprint(total)
}
