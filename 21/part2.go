//go:build part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var dirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	grid := make(map[[2]int]bool) // true for visited, false for wall

	scanner := bufio.NewScanner(os.Stdin)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for col, x := range line {
			if x == '#' {
				grid[[2]int{row, col}] = false
			}
		}
		row++
	}

	// solution relies on properties about the input itself, including the size, and configuration
	// ie the empty row/col containing start at 65,65, and the diamond of empty space at the edge

	// the result is that the reachable amount at 65, 65+131, 65+131*2, ... follows a formula
	// these are when the search gets to the edge of the current layout, then to the edge of the next, and so on
	// since the question's steps are odd, can look only at reachable after odd amount of steps, so will use 65, 65+131*2, 65+131*4
	// another property that allows for this is that the question asks for 26501365 = 65+131*202300 steps

	rows, cols := 131, 131
	start := [2]int{65, 65}

	p0 := calculateReachable(grid, start, rows, cols, 65)
	p1 := calculateReachable(grid, start, rows, cols, 65+131*2)
	p2 := calculateReachable(grid, start, rows, cols, 65+131*4)

	// solve for a, b, c in p(x) = ax^2 + bx + c using p(0), p(1), p(2)
	c := p0
	a := (p2 - 2*p1 + c) / 2
	b := p1 - a - c

	x := (26501365 / 131) / 2 // divide by 2 since p(x) represents 65+131*2x
	ans := a*x*x + b*x + c

	fmt.Println(ans)
}

func calculateReachable(grid map[[2]int]bool, start [2]int, rows, cols int, iters int) int {
	reachableOdd := 0

	grid_ := make(map[[2]int]bool)
	for k, v := range grid {
		grid_[k] = v
	}

	queue := [][2]int{start}
	odd := false
	steps := 0
	for len(queue) > 0 && steps <= iters {
		l := len(queue)
		for i := 0; i < l; i++ {

			cur := queue[i]
			if _, exists := grid_[cur]; exists {
				continue
			}

			grid_[cur] = true
			if odd {
				reachableOdd++
			}

			for _, dir := range dirs {
				nei := [2]int{cur[0] + dir[0], cur[1] + dir[1]}

				if _, exists := grid_[nei]; exists {
					continue
				}

				if notWall, exists := grid_[[2]int{(nei[0]%rows + rows) % rows, (nei[1]%cols + cols) % cols}]; exists && !notWall {
					continue
				}

				queue = append(queue, nei)
			}

		}
		queue = queue[l:]
		odd = !odd
		steps++
	}

	return reachableOdd
}
