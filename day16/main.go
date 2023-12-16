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
	display.Print(16, 1, test_input, input, part1)
	display.Print(16, 2, test_input, input, part2)
}

type vec struct {
	x, y int
}

func (v vec) add(n vec) vec {
	return vec{v.x + n.x, v.y + n.y}
}

type instruction struct {
	pos vec
	dir vec
}

var (
	u = vec{0, -1}
	d = vec{0, 1}
	l = vec{-1, 0}
	r = vec{1, 0}
)

func move(
	components map[vec]rune,
	energized map[vec]bool,
	processed map[instruction]bool,
	i instruction,
) {
	cc, ok := components[i.pos]
	if !ok {
		return
	}
	if processed[i] {
		return
	}
	processed[i] = true
	energized[i.pos] = true
	switch cc {
	case '\\':
		switch i.dir {
		case r:
			i.dir = d
		case l:
			i.dir = u
		case u:
			i.dir = l
		case d:
			i.dir = r
		}
	case '/':
		switch i.dir {
		case r:
			i.dir = u
		case l:
			i.dir = d
		case u:
			i.dir = r
		case d:
			i.dir = l
		}
	case '|':
		switch i.dir {
		case r, l:
			move(components, energized, processed, instruction{i.pos.add(u), u})
			move(components, energized, processed, instruction{i.pos.add(d), d})
			return
		}
	case '-':
		switch i.dir {
		case u, d:
			move(components, energized, processed, instruction{i.pos.add(l), l})
			move(components, energized, processed, instruction{i.pos.add(r), r})
			return
		}
	}
	i.pos = i.pos.add(i.dir)
	move(components, energized, processed, i)
}

func part1(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	components := map[vec]rune{}
	for y, line := range lines {
		for x, ch := range line {
			components[vec{x, y}] = ch
		}
	}
	return fmt.Sprint(calculate(components, instruction{dir: r, pos: vec{}}))
}

func calculate(components map[vec]rune, start instruction) int {
	energized := map[vec]bool{}
	move(components, energized, map[instruction]bool{}, start)
	return len(energized)
}

func part2(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	components := map[vec]rune{}
	for y, line := range lines {
		for x, ch := range line {
			components[vec{x, y}] = ch
		}
	}
	var m int
	// top
	for i := 0; i < len(lines[0]); i++ {
		m = max(m, calculate(components, instruction{dir: d, pos: vec{i, 0}}))
	}
	// bottom
	for i := 0; i < len(lines[0]); i++ {
		m = max(m, calculate(components, instruction{dir: u, pos: vec{i, len(lines) - 1}}))
	}
	// left
	for i := 0; i < len(lines); i++ {
		m = max(m, calculate(components, instruction{dir: r, pos: vec{0, i}}))
	}
	// right
	for i := 0; i < len(lines); i++ {
		m = max(m, calculate(components, instruction{dir: l, pos: vec{len(lines[0]), i}}))
	}
	return fmt.Sprint(m)
}
