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

		total += getPrevVal(seq)
	}

	fmt.Println(total)
}

func getPrevVal(seq []int) int {
	// with a being the first val, b the 1st first difference, c the 1st second difference, etc
	// prevVal = a - (b - (c - (d - (e ...))))
	// = a - b + c - d + e  ...

	accum := 0
	sign := 1
	seqLen := len(seq)
	allZero := false
	for !allZero {
		accum += sign * seq[0]
		sign *= -1

		allZero = true
		for i := 0; i < seqLen-1; i++ {
			diff := seq[i+1] - seq[i]
			if diff != 0 {
				allZero = false
			}
			seq[i] = diff
		}
		seqLen--
	}

	return accum
}
