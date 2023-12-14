package main

import (
	_ "embed"
	"fmt"
	"regexp"
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
	//display.Print(10, 1, test_input, input, part1)
	display.Print(10, 2, test_input, input, part2)
}

type vec struct {
	x, y int
}

func (v vec) add(i vec) vec {
	return vec{v.x + i.x, v.y + i.y}
}

type pipe struct {
	a, b vec
	inA  bool
}

func (p pipe) travel(i vec) (o vec, ok bool) {
	if i == p.a {
		return p.b, true
	}
	if i == p.b {
		return p.a, true
	}
	return vec{}, false
}

var (
	north = vec{0, -1}
	south = vec{0, 1}
	west  = vec{-1, 0}
	east  = vec{1, 0}
	pipes = map[rune]pipe{
		'L': {a: north, b: east},
		'J': {a: north, b: west},
		'F': {a: south, b: east},
		'7': {a: south, b: west},
		'|': {a: south, b: north},
		'-': {a: east, b: west},
	}
	opposites = map[vec]vec{
		north: south,
		west:  east,
		south: north,
		east:  west,
	}
)

func part1(input string) string {
	lines := strings.Split(input, "\n")
	grid := make([][]pipe, len(lines))
	var animal vec
	for y, line := range lines {
		grid[y] = make([]pipe, len([]rune(line)))
		for x, ch := range line {
			if p, ok := pipes[ch]; ok {
				grid[y][x] = p
			}
			if ch == 'S' {
				animal = vec{x, y}
			}
		}
	}
	curr := animal
	var count int
	var entered vec
	// find first pipe
	for _, d := range []vec{north, east, south, west} {
		p := grid[curr.y+d.y][curr.x+d.x]
		if _, ok := p.travel(opposites[d]); ok {
			count++
			curr = curr.add(d)
			entered = d
			break
		}
	}
	// travel pipes
	for curr != animal {
		count++
		p := grid[curr.y][curr.x]
		if o, ok := p.travel(opposites[entered]); ok {
			curr = curr.add(o)
			entered = o
		}
	}
	return fmt.Sprint(count / 2)
}

func part2(input string) string {
	lines := strings.Split(input, "\n")
	grid := make([][]pipe, len(lines))
	var animal vec
	for y, line := range lines {
		grid[y] = make([]pipe, len([]rune(line)))
		for x, ch := range line {
			if p, ok := pipes[ch]; ok {
				grid[y][x] = p
			}
			if ch == 'S' {
				animal = vec{x, y}
			}
		}
	}
	curr := animal
	var entered vec
	isMainPipe := map[vec]bool{}
	// find first pipe
	for _, d := range []vec{east, south, west, north} {
		p := grid[curr.y+d.y][curr.x+d.x]
		if _, ok := p.travel(opposites[d]); ok {
			curr = curr.add(d)
			isMainPipe[curr] = true
			entered = d
			break
		}
	}
	// travel pipes
	for curr != animal {
		p := grid[curr.y][curr.x]
		if o, ok := p.travel(opposites[entered]); ok {
			curr = curr.add(o)
			isMainPipe[curr] = true
			entered = o
		}
	}
	for y, line := range lines {
		l := []rune(line)
		for x := range line {
			if isMainPipe[vec{x, y}] {
				continue
			}
			l[x] = '.'
		}
		lines[y] = string(l)
	}

	var insides int
	for y, line := range lines {
		fmt.Println()
		for x := range line {
			if isMainPipe[vec{x, y}] {
				fmt.Print(string([]byte{lines[y][x]}))
				continue
			}
			if len(verticals.FindAll([]byte(line[x+1:]), -1))%2 == 0 {
				fmt.Print(string([]byte{lines[y][x]}))
				continue
			}
			fmt.Print("I")
			insides++
		}
	}
	return fmt.Sprint(insides)
}

var verticals = regexp.MustCompile("(\\||F-*J|L-*7)")
