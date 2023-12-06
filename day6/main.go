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
	display.Print(6, 1, test_input, input, part1)
	display.Print(6, 2, test_input, input, part2)
}

func beatsRecord(held, time, record int) bool {
	remainingTime := time - held
	distance := remainingTime * held
	return distance > record
}

func part1(input string) string {
	lines := strings.Split(input, "\n")
	timeStrings := strings.Fields(lines[0][9:])
	distanceStrings := strings.Fields(lines[1][9:])
	var total int
	for i := 0; i < len(timeStrings); i++ {
		var ways int
		t, _ := strconv.Atoi(timeStrings[i])
		d, _ := strconv.Atoi(distanceStrings[i])
		for j := 0; j <= t; j++ {
			if beatsRecord(j, t, d) {
				ways++
			}
		}
		if total == 0 {
			total = ways
			continue
		}
		total *= ways
	}
	return fmt.Sprint(total)
}

func part2(input string) string {
	lines := strings.Split(input, "\n")
	timeString := strings.ReplaceAll(lines[0][9:], " ", "")
	distanceString := strings.ReplaceAll(lines[1][9:], " ", "")
	var ways int
	t, _ := strconv.Atoi(timeString)
	d, _ := strconv.Atoi(distanceString)
	for j := 0; j <= t; j++ {
		if beatsRecord(j, t, d) {
			ways++
		}
	}
	return fmt.Sprint(ways)
}
