package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dns\n", time.Since(start).Nanoseconds()) }(time.Now())

	scanner := bufio.NewScanner(os.Stdin)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		var first int
		var cur int

		for _, c := range line {
			if c >= '0' && c <= '9' {
				cur = int(c - '0')
				if first == 0 { // numbers can't be 0, so use this as flag
					first = cur
				}
			}
		}

		total += first*10 + cur
	}

	fmt.Println(total)
}
