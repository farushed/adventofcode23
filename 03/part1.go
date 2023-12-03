//go:build part1

package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	grid := parse()

	total := 0

	for i, line := range grid {
		var num int
		var startIdx, endIdx int
		var symbol bool

		for j, c := range line + "." { // add dot so the code processes numbers at the end
			if c >= '0' && c <= '9' {
				endIdx = j
				num = num*10 + int(c-'0')
			} else {
				if num != 0 { // we finished processing a number
					if symbol || c != '.' || // symbol left or right, so we can include
						symbolInRow(grid, i-1, startIdx-1, endIdx+1) || // symbol above
						symbolInRow(grid, i+1, startIdx-1, endIdx+1) { //symbol below
						total += num
					}
				}
				num = 0
				startIdx = j + 1

				symbol = c != '.'
			}
		}
	}

	fmt.Println(total)
}

func symbolInRow(grid []string, row, start, end int) bool {
	if row < 0 || row >= len(grid) {
		return false
	}

	for i := start; i <= end; i++ {
		if i < 0 || i >= len(grid[row]) {
			continue
		}
		if c := grid[row][i]; c != '.' && !(c >= '0' && c <= '9') {
			return true
		}
	}

	return false
}
