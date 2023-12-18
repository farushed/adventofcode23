//go:build part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var dirs = map[byte][2]int{
	'0': {1, 0},  // R
	'1': {0, -1}, // D
	'2': {-1, 0}, // L
	'3': {0, 1},  // U
}

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	var x, y int = 0, 0
	var perim, area int = 0, 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		steps, _ := strconv.ParseInt(line[strings.IndexRune(line, '#')+1:len(line)-2], 16, 0)
		dir := line[len(line)-2]
		dx, dy := int(steps)*dirs[dir][0], int(steps)*dirs[dir][1]

		area += dx*y - dy*x // shoelace formula
		perim += abs(dx) + abs(dy)

		x, y = x+dx, y+dy
	}

	area = abs(area / 2)

	// area calculated by shoelace formula considers coords as integer vertices
	// we can imagine this as points centered in our grid squares ... the area is missing a border of 1/2 a tile
	// shoelace area = interior points + perim/2 - 1 ... as per pick's theorem
	// interior points = shoelace area - perim/2 + 1 ... rearrange, this is th enumber of interior grid squares
	// answer = interior points + perim = shoelace area + perim/2 + 1

	fmt.Println(area + perim/2 + 1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
