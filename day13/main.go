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
	display.Print(13, 1, test_input, input, part1)
	display.Print(13, 2, test_input, input, part2)
}

type cell bool

func isMirror(row []cell, idx int) bool {
	i := 1
	j := 0
	for idx+i < len(row) && idx-j >= 0 {
		if row[idx+i] != row[idx-j] {
			return false
		}
		j++
		i++
	}
	return true
}

func getMirror(grid [][]cell) int {
lineLoop:
	for i := 0; i < len(grid[0])-1; i++ {
		for _, row := range grid {
			if !isMirror(row, i) {
				continue lineLoop
			}
		}
		return i + 1
	}
	return 0
}
func isMirror2(row []cell, idx,ridx , rx, ry int) bool {
	i := idx + 1
	j := idx
	for i < len(row) && j >= 0 {
		ib := row[i]
		if i == rx && ridx==ry {
			ib = !ib
		}
		jb := row[j]
		if j == rx && ridx==ry {
			jb = !jb
		}
		if ib != jb {
			return false
		}
		j--
		i++
	}
	return true
}

func getMirror2(grid [][]cell) int {
	og := getMirror(grid)
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
		lineLoop:
			for i := 0; i < len(grid[0])-1; i++ {
				if i+1 == og {
					continue
				}
				for r, row := range grid {
					if !isMirror2(row, i, r, x, y) {
						continue lineLoop
					}
				}
				return i + 1
			}
		}
	}
	return 0
}

func part1(input string) string {
	puzzleStrings := strings.Split(input, "\n\n")
	var total int
	for _, p := range puzzleStrings {
		lines := strings.Split(strings.TrimSpace(p), "\n")
		grid := [][]cell{}
		gridH := make([][]cell, len(lines[0]))
		for _, l := range lines {
			var row []cell
			for i, ch := range l {
				row = append(row, ch == '#')
				gridH[i] = append(gridH[i], ch == '#')
			}
			grid = append(grid, row)
		}
		if c := getMirror(grid); c != 0 {
			total += c
			continue
		}
		total += getMirror(gridH) * 100
	}
	return fmt.Sprint(total)
}

func part2(input string) string {
	puzzleStrings := strings.Split(input, "\n\n")
	var total int
	for _, p := range puzzleStrings {
		lines := strings.Split(strings.TrimSpace(p), "\n")
		grid := [][]cell{}
		gridH := make([][]cell, len(lines[0]))
		for _, l := range lines {
			var row []cell
			for i, ch := range l {
				row = append(row, ch == '#')
				gridH[i] = append(gridH[i], ch == '#')
			}
			grid = append(grid, row)
		}
		if c := getMirror2(grid); c != 0 {
			total += c
			continue
		}
		total += getMirror2(gridH) * 100
	}
	return fmt.Sprint(total)
}
