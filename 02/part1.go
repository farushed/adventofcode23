//go:build part1

package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	bagContents := cubes{12, 13, 14}

	games := parse()
	total := 0

	for gameId, reveals := range games {
		if gamePossible(bagContents, reveals) {
			total += gameId
		}
	}

	fmt.Println(total)
}

func gamePossible(bagContents cubes, reveals []cubes) bool {
	for _, reveal := range reveals {
		if reveal.Red > bagContents.Red ||
			reveal.Green > bagContents.Green ||
			reveal.Blue > bagContents.Blue {
			return false
		}
	}
	return true
}
