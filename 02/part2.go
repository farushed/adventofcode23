//go:build part2

package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	games := parse()
	total := 0

	for _, reveals := range games {
		m := minimumRequiredBag(reveals)
		total += m.Red * m.Green * m.Blue
	}

	fmt.Println(total)
}

func minimumRequiredBag(reveals []cubes) cubes {
	minBag := cubes{0, 0, 0}
	for _, reveal := range reveals {
		if reveal.Red > minBag.Red {
			minBag.Red = reveal.Red
		}
		if reveal.Green > minBag.Green {
			minBag.Green = reveal.Green
		}
		if reveal.Blue > minBag.Blue {
			minBag.Blue = reveal.Blue
		}
	}
	return minBag
}
