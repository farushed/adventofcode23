//go:build part1

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var dirs = map[rune][2]int{
	'R': {1, 0},
	'L': {-1, 0},
	'U': {0, 1},
	'D': {0, -1},
}

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	scanner := bufio.NewScanner(os.Stdin)

	bottomLeft, topRight := [2]int{0, 0}, [2]int{0, 0}
	pos := [2]int{0, 0}
	edges := map[[2]int]interface{}{
		pos: struct{}{},
	}

	for scanner.Scan() {
		line := scanner.Text()

		var dir rune
		var steps int
		fmt.Sscanf(line, "%c %d", &dir, &steps)
		change := dirs[dir]

		for i := 0; i < steps; i++ {
			pos[0] += change[0]
			pos[1] += change[1]
			edges[pos] = struct{}{}
		}

		if pos[0] > topRight[0] {
			topRight[0] = pos[0]
		}
		if pos[1] > topRight[1] {
			topRight[1] = pos[1]
		}
		if pos[0] < bottomLeft[0] {
			bottomLeft[0] = pos[0]
		}
		if pos[1] < bottomLeft[1] {
			bottomLeft[1] = pos[1]
		}
	}

	// same idea as day 10 solution

	interior := 0
	for y := topRight[1]; y >= bottomLeft[1]; y-- {
		inside := false
		for x := bottomLeft[0]; x <= topRight[0]; x++ {
			if _, ok := edges[[2]int{x, y}]; ok {
				if _, ok := edges[[2]int{x, y + 1}]; ok {
					inside = !inside
				}
			} else {
				if inside {
					interior++
				}
			}
		}
	}

	fmt.Println(len(edges) + interior)
}
