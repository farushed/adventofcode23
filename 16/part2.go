//go:build part2

package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	grid := parse()

	max := 0

	for i := range grid {
		e := countEnergised(Beam{right, Coord{i, 0}}, grid)
		if e > max {
			max = e
		}
		e = countEnergised(Beam{left, Coord{i, len(grid[i]) - 1}}, grid)
		if e > max {
			max = e
		}
	}

	for i := range grid[0] {
		e := countEnergised(Beam{down, Coord{0, i}}, grid)
		if e > max {
			max = e
		}
		e = countEnergised(Beam{up, Coord{len(grid) - 1, i}}, grid)
		if e > max {
			max = e
		}
	}

	fmt.Println(max)
}
