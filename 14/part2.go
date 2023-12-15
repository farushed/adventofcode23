//go:build part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	scanner := bufio.NewScanner(os.Stdin)

	var lines [][]rune
	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}

	memo := make(map[string]int)
	var states []string
	var state string

	for i := 0; i < 1000000000; i++ {
		var s strings.Builder
		for _, line := range lines {
			s.WriteString(string(line))
			s.WriteRune('\n')
		}
		state = s.String()

		if seenAt, ok := memo[state]; ok {
			idx := (1000000000-seenAt)%(i-seenAt) + seenAt
			state = states[idx]
			break
		} else {
			memo[state] = i
			states = append(states, state)
			doCycle(lines)
		}
	}

	load := 0

	for i, line := range strings.Split(state, "\n") {
		for _, c := range line {
			if c == 'O' {
				load += len(lines) - i
			}
		}
	}

	fmt.Println(load)
}

func doCycle(lines [][]rune) {
	// roll north
	for col := range lines[0] {
		lastFull := -1
		for row := range lines {
			switch lines[row][col] {
			case '#':
				lastFull = row
			case 'O':
				lastFull = lastFull + 1
				lines[row][col] = '.'
				lines[lastFull][col] = 'O'
			}
		}
	}

	// roll west
	for _, line := range lines {
		lastFull := -1
		for col, c := range line {
			switch c {
			case '#':
				lastFull = col
			case 'O':
				lastFull = lastFull + 1
				line[col] = '.'
				line[lastFull] = 'O'
			}
		}
	}

	// roll south
	for col := range lines[0] {
		lastFull := len(lines)
		for row := len(lines) - 1; row >= 0; row-- {
			switch lines[row][col] {
			case '#':
				lastFull = row
			case 'O':
				lastFull = lastFull - 1
				lines[row][col] = '.'
				lines[lastFull][col] = 'O'
			}
		}
	}

	// roll east
	for row := len(lines) - 1; row >= 0; row-- {
		line := lines[row]
		lastFull := len(line)
		for col := len(line) - 1; col >= 0; col-- {
			switch line[col] {
			case '#':
				lastFull = col
			case 'O':
				lastFull = lastFull - 1
				line[col] = '.'
				line[lastFull] = 'O'
			}
		}
	}
}
