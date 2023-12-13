//go:build part1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	scanner := bufio.NewScanner(os.Stdin)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		x := strings.Split(line, " ")
		row := x[0]
		var sizes []int
		for _, s := range strings.Split(x[1], ",") {
			si, _ := strconv.Atoi(s)
			sizes = append(sizes, si)
		}

		total += rec(row+".", sizes) // add . to avoid different edge condition
	}

	fmt.Println(total)
}

func rec(row string, sizes []int) int {
	if len(row) == 0 {
		if len(row) == 0 && len(sizes) == 0 {
			return 1 // consumed everything perfectly
		}
		return 0 // leftover sizes
	}
	if len(sizes) == 0 {
		if strings.ContainsRune(row, '#') {
			return 0 // no sizes left but row not empty
		}
		return 1 // no sizes, row can be empty
	}

	cur := sizes[0]

	total := 0

	// try consume a set of # and/or ?
	if len(row) > cur && !strings.ContainsRune(row[:cur], '.') && row[cur] != '#' {
		total += rec(row[cur+1:], sizes[1:]) // + 1 to account for space
	}

	// skip ., try skip ?
	if row[0] != '#' {
		total += rec(row[1:], sizes)
	}

	return total
}
