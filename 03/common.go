package main

import (
	"bufio"
	"os"
)

func parse() []string {
	scanner := bufio.NewScanner(os.Stdin)

	var grid []string

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	return grid
}
