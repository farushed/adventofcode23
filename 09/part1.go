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

		var seq []int
		for _, n := range strings.Fields(line) {
			x, _ := strconv.Atoi(n)
			seq = append(seq, x)
		}

		total += getNextVal(seq)
	}

	fmt.Println(total)
}

func getNextVal(seq []int) int {
	accum := 0
	seqLen := len(seq)
	allZero := false
	for !allZero {
		allZero = true
		for i := 0; i < seqLen-1; i++ {
			diff := seq[i+1] - seq[i]
			if diff != 0 {
				allZero = false
			}
			seq[i] = diff
		}
		accum += seq[seqLen-1]
		seqLen--
	}

	return accum
}
