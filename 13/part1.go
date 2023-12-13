//go:build part1

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
			if reflectsHorizontal(p, i) {
				numRows += i
			}
		}
		for i := 1; i < len(p[0]); i++ {
			if reflectsVertical(p, i) {
				numCols += i
			}
		}
	}

	fmt.Println(numCols + 100*numRows)
}

func reflectsHorizontal(pattern []string, beforeRow int) bool {
	for i := 0; i < beforeRow; i++ {
		if 2*beforeRow-i-1 < len(pattern) && pattern[i] != pattern[2*beforeRow-i-1] {
			return false
		}
	}
	return true
}

func reflectsVertical(pattern []string, beforeCol int) bool {
	for j := 0; j < len(pattern); j++ {
		for i := 0; i < beforeCol; i++ {
			if 2*beforeCol-i-1 < len(pattern[j]) && pattern[j][i] != pattern[j][2*beforeCol-i-1] {
				return false
			}
		}
	}
	return true
}
