package main

import (
	_ "embed"
	"fmt"
	"math/bits"
	"regexp"
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
	display.Print(12, 1, test_input, input, part1)
	display.Print(12, 2, test_input, input, part2)
}

func pow(i int, exp int) int {
	t := i
	for j := 1; j < exp; j++ {
		t *= i
	}
	return t

}

func check(springs string, groups []int) bool {
	s := strings.FieldsFunc(springs, func(r rune) bool { return r == '.' })
	if len(s) != len(groups) {
		return false
	}
	for i, g := range s {
		if len(g) != groups[i] {
			return false
		}
	}
	return true
}

func part1(input string) string {
	var total int
	for _, line := range strings.Split(input, "\n") {
		springs, grs, _ := strings.Cut(line, " ")
		var groups []int
		for _, i := range strings.Split(grs, ",") {
			n, _ := strconv.Atoi(i)
			groups = append(groups, n)
		}
		slots := regexp.MustCompile("\\?").FindAllIndex([]byte(springs), -1)
		possibilities := pow(2, len(slots))
		for i := 0; i <= possibilities; i++ {
			curr := []rune(springs)
			for idx, sl := range slots {
				curr[sl[0]] = '.'
				if i&(1<<idx) != 0 {
					curr[sl[0]] = '#'
				}
			}
			if check(string(curr), groups) {
				total++
			}
		}
	}
	return fmt.Sprint(total)
}

var (
	unknown = regexp.MustCompile("\\?")
	damaged = regexp.MustCompile("#")
)

func part2(input string) string {
	var total int
	for _, line := range strings.Split(input, "\n") {
		springs, grs, _ := strings.Cut(line, " ")
		springs = strings.Join([]string{springs, springs, springs, springs, springs}, "?")
		grs = strings.Join([]string{grs, grs, grs, grs, grs}, ",")
		var groups []int
		var target int
		for _, i := range strings.Split(grs, ",") {
			n, _ := strconv.Atoi(i)
			target += n
			groups = append(groups, n)
		}
		slots := unknown.FindAllIndex([]byte(springs), -1)
		knownDamaged := damaged.FindAllIndex([]byte(springs), -1)
		unknownDamaged := target - len(knownDamaged)
		possibilities := pow(2, len(slots))
		var subtotal int
		for i := 0; i <= possibilities; i++ {
			if bits.OnesCount(uint(i)) != unknownDamaged {
				continue
			}
			curr := []rune(springs)
			for idx, sl := range slots {
				curr[sl[0]] = '.'
				if i&(1<<idx) != 0 {
					curr[sl[0]] = '#'
				}
			}
			if check(string(curr), groups) {
				subtotal++
			}
		}
		fmt.Println(subtotal)
		total += subtotal
	}
	return fmt.Sprint(total * 5)
}
