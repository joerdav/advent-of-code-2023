package display

import (
	"fmt"
	"time"
)

func Print(task, part int, test, real string, impl func(string) string) {
	fmt.Printf("%d.%d\n", task, part)
	start := time.Now()
	real1 := impl(real)
	duration := time.Since(start)
	fmt.Printf("  real: %s (%v)\n", real1, duration)
	fmt.Printf("  test: %s\n", impl(test))
}
