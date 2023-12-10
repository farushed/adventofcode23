//go:build part2

package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	start, connections, rows, cols := parse()

	cur := connections[start][0]
	prev := start

	onLoop := make(map[[2]int]interface{})
	onLoop[start] = struct{}{}

	for cur != start {
		onLoop[cur] = struct{}{}

		next := connections[cur][0]
		if next == prev {
			next = connections[cur][1]
		}

		prev, cur = cur, next
	}

	enclosed := 0

	for r := 0; r < rows; r++ {
		inside := false
		for c := 0; c < cols; c++ {
			cur := [2]int{r, c}
			if _, ok := onLoop[cur]; ok {
				// only consider crossing the polygon boundary when we cross a pipe connected up
				// this accounts for traversing horizontal stretches,
				// and whether the corners of that stretch face the same direction or not
				if connections[cur][0][0] == r-1 || connections[cur][1][0] == r-1 {
					inside = !inside
				}
			} else {
				if inside {
					enclosed++
				}
			}
		}
	}

	fmt.Println(enclosed)
}
