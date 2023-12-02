//go:build part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dns\n", time.Since(start).Nanoseconds()) }(time.Now())

	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	scanner := bufio.NewScanner(os.Stdin)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		var first, cur int

		pattern := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|\\d)")

		idx := pattern.FindStringIndex(line)
		for idx != nil {
			number := line[idx[0]:idx[1]]
			if len(number) == 1 && number[0] >= '0' && number[0] <= '9' {
				cur = int(number[0] - '0')
			} else {
				cur = numberMap[number]
			}
			// fmt.Printf("%s %d\t", number, cur)

			if first == 0 {
				first = cur
			}

			line = line[idx[0]+1:]
			idx = pattern.FindStringIndex(line)
		}
		// fmt.Println(first*10 + cur)

		total += first*10 + cur
	}

	fmt.Println(total)
}
