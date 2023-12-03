//go:build part2

package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	grid := parse()

	gears := make(map[[2]int][]int)

	for i, line := range grid {
		var num int
		var startIdx, endIdx int

		for j, c := range line + "." {
			if c >= '0' && c <= '9' {
				endIdx = j
				num = num*10 + int(c-'0')
			} else {
				if num != 0 { // we finished processing a number
					storeNumIfGearInRange(grid, num, i-1, startIdx-1, endIdx+1, gears)
					storeNumIfGearInRange(grid, num, i+1, startIdx-1, endIdx+1, gears)
					storeNumIfGearAtPos(grid, num, i, startIdx-1, gears)
					storeNumIfGearAtPos(grid, num, i, endIdx+1, gears)
				}
				num = 0
				startIdx = j + 1
			}
		}
	}

	total := 0
	for _, vals := range gears {
		if len(vals) == 2 {
			total += vals[0] * vals[1]
		}
	}

	fmt.Println(total)
}

func storeNumIfGearInRange(grid []string, num, row, start, end int, gears map[[2]int][]int) {
	if row < 0 || row >= len(grid) {
		return
	}

	for i := start; i <= end; i++ {
		storeNumIfGearAtPos(grid, num, row, i, gears)
	}
}

func storeNumIfGearAtPos(grid []string, num, row, col int, gears map[[2]int][]int) {
	if col < 0 || col >= len(grid[row]) {
		return
	}

	if grid[row][col] == '*' {
		gears[[2]int{row, col}] = append(gears[[2]int{row, col}], num)
	}
}
