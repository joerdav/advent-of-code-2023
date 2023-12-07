package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"github.com/joerdav/advent-of-code-2023/display"
)

var (
	//go:embed input.txt
	input string
	//go:embed test.txt
	test_input string
)

func main() {
	display.Print(7, 1, test_input, input, part1)
	display.Print(7, 2, test_input, input, part2)
}

type hand struct {
	cards   []int
	s       string
	cardMap map[rune]int
	bid     int
}

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

const (
	T = iota + 10
	J
	Q
	K
	A
)

func (h hand) handRank() int {
	l := len(h.cardMap)
	if l == 1 {
		return FiveOfAKind
	}
	for _, c := range h.cardMap {
		if c == 4 {
			return FourOfAKind
		}
	}
	if l == 2 {
		return FullHouse
	}
	for _, c := range h.cardMap {
		if c == 3 {
			return ThreeOfAKind
		}
	}
	if l == 3 {
		return TwoPair
	}
	if l == 4 {
		return OnePair
	}
	return HighCard
}

func cardval(r rune) int {
	if unicode.IsDigit(r) {
		return int(r - '0')
	}
	switch r {
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	}
	return -1
}

func part1(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var hands []hand
	for _, l := range lines {
		var h hand
		c, b, _ := strings.Cut(l, " ")
		h.cardMap = map[rune]int{}
		h.bid, _ = strconv.Atoi(b)
		h.s = c
		for _, ca := range c {
			h.cards = append(h.cards, cardval(ca))
			h.cardMap[ca]++
		}
		hands = append(hands, h)
	}
	slices.SortFunc(hands, func(l, r hand) int {
		lr, rr := l.handRank(), r.handRank()
		if lr != rr {
			return lr - rr
		}
		for i := range l.cards {
			if l.cards[i] != r.cards[i] {
				return l.cards[i] - r.cards[i]
			}
		}
		panic("equal hands")
	})
	var total int
	for i, hand := range hands {
		r := i + 1
		total += hand.bid * r
	}
	return fmt.Sprint(total)
}

func cardval2(r rune) int {
	if unicode.IsDigit(r) {
		return int(r - '0')
	}
	switch r {
	case 'T':
		return 10
	case 'J':
		return 1
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	}
	return -1
}

func (h hand) handRank2() int {
	jokers := h.cardMap['J']
	l := len(h.cardMap)
	if l == 1 || (l == 2 && jokers > 0) {
		return FiveOfAKind
	}
	for v, c := range h.cardMap {
		if c == 4 {
			return FourOfAKind
		}
		if v == 'J' {
			continue
		}
		if c+jokers == 4 {
			return FourOfAKind
		}
	}
	if l == 2 || (l == 3 && jokers > 0) {
		return FullHouse
	}
	for _, c := range h.cardMap {
		if c+jokers == 3 {
			return ThreeOfAKind
		}
	}
	if l == 3 || (l == 4 && jokers > 0) {
		return TwoPair
	}
	if l == 4 || jokers > 0 {
		return OnePair
	}
	return HighCard
}

func part2(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var hands []hand
	for _, l := range lines {
		var h hand
		c, b, _ := strings.Cut(l, " ")
		h.cardMap = map[rune]int{}
		h.bid, _ = strconv.Atoi(b)
		h.s = c
		for _, ca := range c {
			h.cards = append(h.cards, cardval2(ca))
			h.cardMap[ca]++
		}
		hands = append(hands, h)
	}
	slices.SortFunc(hands, func(l, r hand) int {
		lr, rr := l.handRank2(), r.handRank2()
		if lr != rr {
			return lr - rr
		}
		for i := range l.cards {
			if l.cards[i] != r.cards[i] {
				return l.cards[i] - r.cards[i]
			}
		}
		panic("equal hands")
	})
	var total int
	for i, hand := range hands {
		r := i + 1
		total += hand.bid * r
	}
	return fmt.Sprint(total)
}
