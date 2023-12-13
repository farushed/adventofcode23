package main

import (
	"bufio"
	"os"
)

func parse() [][]string {
	scanner := bufio.NewScanner(os.Stdin)

	var patterns [][]string
	var pattern []string

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			patterns = append(patterns, pattern)
			pattern = make([]string, 0)
		} else {
			pattern = append(pattern, line)
		}
	}
	patterns = append(patterns, pattern)

	return patterns
}
