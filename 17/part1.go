//go:build part1

package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	weights := parse()

	leastCost := dijkstra(
		weights,
		func(straightCount int) bool { return straightCount < 3 }, // condition to go straight
		func(straightCount int) bool { return true },              // condition to turn or stop
	)

	fmt.Println(leastCost)
}
