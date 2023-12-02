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
	display.Print(2, 1, test_input, input, part1)
	display.Print(2, 2, test_input, input, part2)
}

func gamePossible1(game string, limits map[string]int) int {
		gameName, cubeSets, _ := strings.Cut(game, ":")
		gameNum, _ := strconv.Atoi(strings.Fields(gameName)[1])
		for _, set := range strings.Split(cubeSets, ";") {
			for _, col := range strings.Split(set, ",") {
				colCount, colName, _ := strings.Cut(strings.TrimSpace(col), " ")
				colCounti, _ := strconv.Atoi(colCount)
				if colCounti > limits[colName] {
					return 0
				}
			}
		}
		return gameNum
}
func part1(input string) string {
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	var count int
	for _, game := range strings.Split(input, "\n") {
		if strings.TrimSpace(game) == "" {
			continue
		}
		count += gamePossible1(game, limits)
	}
	return fmt.Sprint(count)
}

func gamePowers(game string) int {
	mins := make(map[string]int)
		_, cubeSets, _ := strings.Cut(game, ":")
		for _, set := range strings.Split(cubeSets, ";") {
			for _, col := range strings.Split(set, ",") {
				colCount, colName, _ := strings.Cut(strings.TrimSpace(col), " ")
				colCounti, _ := strconv.Atoi(colCount)
				if colCounti > mins[colName] {
					mins[colName] = colCounti
				}
			}
		}
		return mins["red"] * mins["green"] * mins["blue"]
}

func part2(input string) string {
	var count int
	for _, game := range strings.Split(input, "\n") {
		if strings.TrimSpace(game) == "" {
			continue
		}
		count += gamePowers(game)
	}
	return fmt.Sprint(count)
}
