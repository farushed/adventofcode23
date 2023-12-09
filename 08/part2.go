//go:build part2

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

	lefts := make(map[string]string)
	rights := make(map[string]string)
	var starts []string

	scanner.Scan()
	instructions := scanner.Text()

	scanner.Scan() // skip newline
	for scanner.Scan() {
		var node, left, right string
		fmt.Sscanf(scanner.Text(), "%3s = (%3s, %3s)", &node, &left, &right)
		lefts[node] = left
		rights[node] = right
		if node[2] == 'A' {
			starts = append(starts, node)
		}
	}

	var stepsList []int
	for _, start := range starts {
		cur := start
		steps := 0
	Outer:
		for {
			for _, dir := range instructions {
				if dir == 'L' {
					cur = lefts[cur]
				} else {
					cur = rights[cur]
				}
				steps++
				if cur[2] == 'Z' {
					break Outer
				}
			}
		}

		stepsList = append(stepsList, steps)
	}

	lcm := stepsList[0]
	for _, x := range stepsList[1:] {
		lcm = (lcm * x) / gcd(lcm, x)
	}

	fmt.Println(lcm)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
