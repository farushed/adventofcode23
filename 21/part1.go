//go:build part1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var dirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	var grid [][]rune
	var start [2]int

	scanner := bufio.NewScanner(os.Stdin)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))

		if col := strings.IndexRune(line, 'S'); col != -1 {
			start = [2]int{row, col}
		}
		row++
	}

	fmt.Println(start)

	reachable := 0

	queue := [][2]int{start}
	odd := false
	steps := 0
	for len(queue) > 0 && steps <= 64 {
		l := len(queue)
		for i := 0; i < l; i++ {

			cur := queue[i]
			if g := grid[cur[0]][cur[1]]; g != '.' && g != 'S' {
				continue
			}
			if odd {
				grid[cur[0]][cur[1]] = 'o'
			} else {
				grid[cur[0]][cur[1]] = 'e'
				reachable++
			}

			for _, dir := range dirs {
				nr, nc := cur[0]+dir[0], cur[1]+dir[1]

				if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[nr]) ||
					grid[nr][nc] != '.' {
					continue
				}

				queue = append(queue, [2]int{nr, nc})
			}

		}
		queue = queue[l:]
		odd = !odd
		steps++
	}

	fmt.Println(steps)
	for _, l := range grid {
		fmt.Println(string(l))
	}
	fmt.Println(reachable)
}
