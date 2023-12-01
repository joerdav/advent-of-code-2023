package main

import (
	_ "embed"
	"fmt"
	"time"
)

var (
	//go:embed input.txt
	input string
	//go:embed test.txt
	test_input string
)

func main() {
	fmt.Println("1.1")
	start := time.Now()
	real1 := part1(input)
	duration := time.Since(start)
	fmt.Printf("  real: %s (%dms)\n", real1, duration.Milliseconds())
	fmt.Printf("  test: %s\n", part1(test_input))
	fmt.Println("1.2")
	start2 := time.Now()
	real2 := part2(input)
	duration2 := time.Since(start2)
	fmt.Printf("  real: %s (%dms)\n", real2, duration2.Milliseconds())
	fmt.Printf("  test: %s\n", part2(test_input))

}

func part1(input string) string {
	return "not implemented"
}

func part2(input string) string {
	return "not implemented"
}
