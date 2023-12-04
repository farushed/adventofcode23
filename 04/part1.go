// // go:build part1

package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	cards := parse()

	total := 0

	for _, card := range cards {
		val := 0
		for num := range card.have {
			if _, ok := card.winning[num]; ok {
				if val == 0 {
					val = 1
				} else {
					val *= 2
				}
			}
		}
		total += val
	}

	fmt.Println(total)
}
