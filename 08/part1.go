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

	lefts := make(map[string]string)
	rights := make(map[string]string)

	scanner.Scan()
	instructions := scanner.Text()

	scanner.Scan() // skip newline
	for scanner.Scan() {
		var node, left, right string
		fmt.Sscanf(scanner.Text(), "%3s = (%3s, %3s)", &node, &left, &right)
		lefts[node] = left
		rights[node] = right
	}

	steps := 0
	cur := "AAA"
Outer:
	for {
		for _, dir := range instructions {
			if dir == 'L' {
				cur = lefts[cur]
			} else {
				cur = rights[cur]
			}
			steps++
			if cur == "ZZZ" {
				break Outer
			}
		}
	}

	fmt.Println(steps)
}
