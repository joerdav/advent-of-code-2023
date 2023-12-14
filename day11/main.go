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
	display.Print(11, 1, test_input, input, part1)
	display.Print(11, 2, test_input, input, part2)
}

type vec struct {
	x, y int
}

func part1(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	extraCols := make([]int, len(lines[0]))
	extraRows := make([]int, len(lines))
	galaxies := []vec{}

	for i, row := range lines {
		var found bool
		for j, ch := range row {
			if ch == '#' {
				found = true
				galaxies = append(galaxies, vec{j, i})
			}
		}
		if !found {
			extraRows[i] = 1
		}
	}
	for i := range lines[0] {
		var found bool
		for _, row := range lines {
			if row[i] == '#' {
				found = true
				break
			}
		}
		if !found {
			extraCols[i] = 1
		}
	}
	var total int
	for i := range galaxies {
		if i == len(galaxies)-1 {
			continue
		}
		for _, j := range galaxies[i+1:] {

			// add distance
			t := Abs(j.x - galaxies[i].x) + Abs(j.y - galaxies[i].y)
			// add extra
			s, e := min(galaxies[i].x, j.x), max(galaxies[i].x, j.x)
			for _, e := range extraCols[s:e] {
				t +=	Abs(e)
			}
			s, e = min(galaxies[i].y, j.y), max(galaxies[i].y, j.y)
			for _, e := range extraRows[s:e] {
				t +=	Abs(e)
			}
			total += t
		}
	}
	return fmt.Sprint(total)
}

func Abs(x int) int {
	if x<0 {
		return-x
	}
	return x
}

func part2(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	extraCols := make([]int, len(lines[0]))
	extraRows := make([]int, len(lines))
	galaxies := []vec{}

	for i, row := range lines {
		var found bool
		for j, ch := range row {
			if ch == '#' {
				found = true
				galaxies = append(galaxies, vec{j, i})
			}
		}
		if !found {
			extraRows[i] = 999999
		}
	}
	for i := range lines[0] {
		var found bool
		for _, row := range lines {
			if row[i] == '#' {
				found = true
				break
			}
		}
		if !found {
			extraCols[i] = 999999
		}
	}
	var total int
	for i := range galaxies {
		if i == len(galaxies)-1 {
			continue
		}
		for _, j := range galaxies[i+1:] {

			// add distance
			t := Abs(j.x - galaxies[i].x) + Abs(j.y - galaxies[i].y)
			// add extra
			s, e := min(galaxies[i].x, j.x), max(galaxies[i].x, j.x)
			for _, e := range extraCols[s:e] {
				t +=	Abs(e)
			}
			s, e = min(galaxies[i].y, j.y), max(galaxies[i].y, j.y)
			for _, e := range extraRows[s:e] {
				t +=	Abs(e)
			}
			total += t
		}
	}
	return fmt.Sprint(total)
}
