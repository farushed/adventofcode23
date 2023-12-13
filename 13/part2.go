//go:build part2

package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	patterns := parse()

	numCols := 0
	numRows := 0
	for _, p := range patterns {
		for i := 1; i < len(p); i++ {
			if reflectsHorizontal(p, i) == 1 {
				numRows += i
			}
		}
		for i := 1; i < len(p[0]); i++ {
			if reflectsVertical(p, i) == 1 {
				numCols += i
			}
		}
	}

	fmt.Println(numCols + 100*numRows)
}

func reflectsHorizontal(pattern []string, beforeRow int) int {
	diffs := 0
	for i := 0; i < beforeRow; i++ {
		if 2*beforeRow-i-1 < len(pattern) {
			for j := 0; j < len(pattern[i]); j++ {
				if pattern[i][j] != pattern[2*beforeRow-i-1][j] {
					diffs++
				}
			}
		}
	}
	return diffs
}

func reflectsVertical(pattern []string, beforeCol int) int {
	diffs := 0
	for j := 0; j < len(pattern); j++ {
		for i := 0; i < beforeCol; i++ {
			if 2*beforeCol-i-1 < len(pattern[j]) && pattern[j][i] != pattern[j][2*beforeCol-i-1] {
				diffs++
			}
		}
	}
	return diffs
}
