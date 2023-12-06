package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/joerdav/advent-of-code-2023/display"
)

var (
	//go:embed input.txt
	input string
	//go:embed test.txt
	test_input string
)

func main() {
	display.Print(5, 1, test_input, input, part1)
	display.Print(5, 2, test_input, input, part2)
}

type idMapItem struct {
	source, dest, rang int
}

func (m idMap) convert(f int) int {
	for _, mi := range m.maps {
		if mi.source <= f && f < mi.source+mi.rang {
			dif := f - mi.source
			return mi.dest + dif
		}
	}
	return f
}

type idMap struct {
	to   string
	maps []idMapItem
}

func part1(input string) string {
	idMaps := make(map[string]idMap)
	lines := strings.Split(input, "\n")
	for i := 1; i < len(lines); i++ {
		l := strings.TrimSpace(lines[i])
		if l == "" {
			continue
		}
		from, to, _ := strings.Cut(strings.Fields(l)[0], "-to-")
		mi := idMap{to: to}
		for {
			i++
			if i >= len(lines) || strings.TrimSpace(lines[i]) == "" {
				break
			}
			m := strings.Fields(lines[i])
			sour, _ := strconv.Atoi(m[1])
			dest, _ := strconv.Atoi(m[0])
			size, _ := strconv.Atoi(m[2])
			mi.maps = append(mi.maps, idMapItem{sour, dest, size})
		}
		slices.SortFunc(mi.maps, func(l, r idMapItem) int { return l.source - r.source })
		idMaps[from] = mi
	}
	result := math.MaxInt
	seedsString := strings.Split(lines[0], ": ")[1]
	seedStrings := strings.Fields(strings.TrimSpace(seedsString))
	for _, seed := range seedStrings {
		v, _ := strconv.Atoi(seed)
		curr := "seed"
		for {
			m, ok := idMaps[curr]
			if !ok {
				break
			}
			curr = m.to
			v = m.convert(v)
		}
		result = min(result, v)
	}
	return fmt.Sprint(result)
}

func part2(input string) string {
	idMaps := make(map[string]idMap)
	lines := strings.Split(input, "\n")
	for i := 1; i < len(lines); i++ {
		l := strings.TrimSpace(lines[i])
		if l == "" {
			continue
		}
		from, to, _ := strings.Cut(strings.Fields(l)[0], "-to-")
		mi := idMap{to: to}
		for {
			i++
			if i >= len(lines) || strings.TrimSpace(lines[i]) == "" {
				break
			}
			m := strings.Fields(lines[i])
			sour, _ := strconv.Atoi(m[1])
			dest, _ := strconv.Atoi(m[0])
			size, _ := strconv.Atoi(m[2])
			mi.maps = append(mi.maps, idMapItem{sour, dest, size})
		}
		slices.SortFunc(mi.maps, func(l, r idMapItem) int { return l.source - r.source })
		idMaps[from] = mi
	}
	seedsString := strings.Split(lines[0], ": ")[1]
	seedStrings := strings.Fields(strings.TrimSpace(seedsString))
	result := math.MaxInt
	var resultMutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < len(seedStrings); i += 2 {
		v, _ := strconv.Atoi(seedStrings[i])
		r, _ := strconv.Atoi(seedStrings[i+1])
		wg.Add(1)
		go func() {
			defer wg.Done()
			subResult := math.MaxInt
			for x := v; x < v+r; x++ {
				val := x
				curr := "seed"
				for {
					m, ok := idMaps[curr]
					if !ok {
						break
					}
					curr = m.to
					val = m.convert(val)
				}
				if result < val {
					return
				}
				subResult = min(subResult, val)
			}
			resultMutex.Lock()
			defer resultMutex.Unlock()
			result = min(result, subResult)
		}()
	}
	wg.Wait()
	return fmt.Sprint(result)
}
