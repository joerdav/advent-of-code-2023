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
	display.Print(4, 1, test_input, input, part1)
	display.Print(4, 2, test_input, input, part2)
}

func part1(input string) string {
	var total int
	winningMap := make(map[string]bool)
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		_, card, _ := strings.Cut(line, ":")
		winningNums, cardNums, _ := strings.Cut(strings.TrimSpace(card), "|")
		for _, num := range strings.Fields(strings.TrimSpace(winningNums)) {
			winningMap[num] = true
		}
		var score int
		for _, num := range strings.Fields(strings.TrimSpace(cardNums)) {
			if !winningMap[num] {
				continue
			}
			if score == 0 {
				score = 1
				continue
			}
			score *= 2
		}
		total += score
		clear(winningMap)
	}
	return fmt.Sprint(total)
}

type card struct {
	winningNums, cardNums string
	foundScore            *int
}

func calcScore(win, card string) int {
	winningMap := make(map[string]bool)
	for _, num := range strings.Fields(strings.TrimSpace(win)) {
		winningMap[num] = true
	}
	var score int
	for _, num := range strings.Fields(strings.TrimSpace(card)) {
		if winningMap[num] {
			score++
		}
	}
	return score
}
func calcCards(cards []*card, card int) int {
	// 1 for original
	score := 1
	c := cards[card]
	if c.foundScore == nil {
		score := calcScore(c.winningNums, c.cardNums)
		c.foundScore = &score
	}
	if *c.foundScore == 0 {
		return score
	}
	for i := card+1; i<=card+*c.foundScore; i++ {
		score += calcCards(cards, i)
	}
	return score
}

func part2(input string) string {
	var total int
	lines := strings.Split(strings.TrimSpace(input), "\n")
	cards := make([]*card, len(lines))
	for i, line := range lines {
		_, cardString, _ := strings.Cut(line, ":")
		winningNums, cardNums, _ := strings.Cut(strings.TrimSpace(cardString), "|")
		cards[i] = &card{winningNums, cardNums, nil}
	}
	for i := range cards {
		total += calcCards(cards, i)
	}
	return fmt.Sprint(total)
}
