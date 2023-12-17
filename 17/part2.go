//go:build part2

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
		func(straightCount int) bool { return straightCount < 10 }, // condition to go straight
		func(straightCount int) bool { return straightCount >= 4 }, // condition to turn or stop
	)

	fmt.Println(leastCost)
}
