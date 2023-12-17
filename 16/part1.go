//go:build part1

package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	grid := parse()

	result := countEnergised(Beam{right, Coord{0, 0}}, grid)

	fmt.Println(result)
}
