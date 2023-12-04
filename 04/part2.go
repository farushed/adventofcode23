// go:build part2

package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	cards := parse()
	extraCards := make(map[int]int)

	total := 0

	for i, card := range cards {
		matching := 0
		for num := range card.have {
			if _, ok := card.winning[num]; ok {
				matching++
			}
		}

		for j := 1; j <= matching; j++ {
			extraCards[i+j] += 1 + extraCards[i]

		}

		total += 1 + extraCards[i]
	}

	fmt.Println(total)
}
