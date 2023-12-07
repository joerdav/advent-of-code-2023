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
	lines := strings.Split(strings.TrimSpace(input), "\n")
	gameCut := strings.IndexRune(lines[0], ':')
	cardCut := strings.IndexRune(lines[0], '|')
	for _, line := range lines {
		winningNums, cardNums := line[gameCut+1:cardCut], line[cardCut+1:]
		for i := 0; i+2 < len(winningNums); i += 3 {
			winningMap[winningNums[i:i+3]] = true
		}
		var score int
		for i := 0; i+2 < len(cardNums); i += 3 {
			if !winningMap[cardNums[i:i+3]] {
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
	foundScore            int
}

var winningMap = make(map[string]bool, 10)

func calcScore(win, card string) int {
	for i := 0; i+2 < len(win); i += 3 {
		winningMap[win[i:i+3]] = true
	}
	var score int
	for i := 0; i+2 <= len(card); i += 3 {
		if winningMap[card[i:i+3]] {
			score++
		}
	}
	clear(winningMap)
	return score
}
func calcCards(cards []card, card int) int {
	c := cards[card]
	if c.foundScore >= 0 {
		return c.foundScore
	}
	// 1 for original
	score := 1
	count := calcScore(c.winningNums, c.cardNums)
	for i := card + 1; i <= card+count; i++ {
		score += calcCards(cards, i)
	}
	c.foundScore = score
	cards[card] = c
	return score
}

func part2(input string) string {
	var total int
	lines := strings.Split(strings.TrimSpace(input), "\n")
	cards := make([]card, len(lines))
	gameCut := strings.IndexRune(lines[0], ':')
	cardCut := strings.IndexRune(lines[0], '|')
	for i, line := range lines {
		winningNums, cardNums := line[gameCut+1:cardCut], line[cardCut+1:]
		cards[i] = card{winningNums, cardNums, -1}
	}
	for i := range cards {
		total += calcCards(cards, i)
	}
	return fmt.Sprint(total)
}
