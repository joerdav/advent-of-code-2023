package main

import (
	_ "embed"
	"fmt"
	"strings"
	"sync"

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
	display.Print(8, 1, test_input, input, part1)
	display.Print(8, 2, test_input2, input, part2)
}

type node struct {
	n, l, r string
}

func (n node) String() string {
	return n.n + " = (" + n.l + ", " + n.r + ")"
}

func part1(input string) string {
	lines := strings.Split(input, "\n")
	instructions := strings.TrimSpace(lines[0])
	nodes := make(map[string]node)
	for _, l := range lines[1:] {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		name, lr, _ := strings.Cut(l, " = ")
		l, r, _ := strings.Cut(strings.Trim(lr, "()"), ", ")
		nodes[name] = node{name, l, r}
	}
	var steps int
	curr := "AAA"
	for curr != "ZZZ" {
		i := instructions[steps%len(instructions)]
		steps++
		if i == 'R' {
			curr = nodes[curr].r
			continue
		}
		curr = nodes[curr].l
	}
	return fmt.Sprint(steps)
}

func part2(input string) string {
	lines := strings.Split(input, "\n")
	instructions := strings.TrimSpace(lines[0])
	nodes := make(map[string]node)
	var routes []string
	for _, l := range lines[1:] {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		name, lr, _ := strings.Cut(l, " = ")
		if strings.HasSuffix(name, "A") {
			routes = append(routes, name)
		}
		l, r, _ := strings.Cut(strings.Trim(lr, "()"), ", ")
		nodes[name] = node{name, l, r}
	}
	matches := make([]int, len(routes))
	var wg sync.WaitGroup
	for i := range routes {
		wg.Add(1)
		i := i
		var steps int
		go func() {
			defer wg.Done()
			for {
				in := instructions[steps%len(instructions)]
				steps++
				if in == 'R' {
					routes[i] = nodes[routes[i]].r
				} else {
					routes[i] = nodes[routes[i]].l
				}
				if strings.HasSuffix(routes[i], "Z") {
					matches[i] = steps
					return
				}
			}
		}()
	}
	wg.Wait()
	r := 1
	for _, ma := range matches {
		r = r * ma / gcd(r, ma)
	}
	return fmt.Sprint(r)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
