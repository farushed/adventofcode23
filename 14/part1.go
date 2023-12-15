//go:build part1

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	load := 0

	for col := range lines[0] {
		lastFull := -1

		for row := range lines {
			switch lines[row][col] {
			case '#':
				lastFull = row
			case 'O':
				lastFull = lastFull + 1
				load += (len(lines) - lastFull)
			}
		}
	}

	fmt.Println(load)
}
