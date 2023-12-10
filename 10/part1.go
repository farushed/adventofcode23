//go:build part1

package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	start, connections, _, _ := parse()

	x := connections[start]
	side1, side2 := x[0], x[1]
	prev1, prev2 := start, start

	steps := 1
	for side1 != side2 {
		next1 := connections[side1][0]
		if next1 == prev1 {
			next1 = connections[side1][1]
		}
		prev1, side1 = side1, next1

		next2 := connections[side2][0]
		if next2 == prev2 {
			next2 = connections[side2][1]
		}
		prev2, side2 = side2, next2

		if side1 == prev2 {
			break
		}
		steps++
	}

	fmt.Println(steps)
}
