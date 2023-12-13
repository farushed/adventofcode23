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

		x := strings.Split(line, " ")
		row := x[0]
		var sizes []int
		var tmpSizes []int
		for _, s := range strings.Split(x[1], ",") {
			si, _ := strconv.Atoi(s)
			tmpSizes = append(tmpSizes, si)
		}
		for i := 0; i < 5; i++ {
			sizes = append(sizes, tmpSizes...)
		}

		memo := make(map[[2]int]int)
		ret := rec(row+"?"+row+"?"+row+"?"+row+"?"+row+".", sizes, 0, 0, memo)

		total += ret
	}

	fmt.Println(total)
}

func rec(row string, sizes []int, i, j int, memo map[[2]int]int) int {
	if val, ok := memo[[2]int{i, j}]; ok {
		return val
	}

	if i == len(row) {
		if i == len(row) && j == len(sizes) {
			memo[[2]int{i, j}] = 1
			return 1
		}
		memo[[2]int{i, j}] = 0
		return 0
	}
	if j == len(sizes) {
		if strings.ContainsRune(row[i:], '#') {
			memo[[2]int{i, j}] = 0
			return 0
		}
		memo[[2]int{i, j}] = 1
		return 1
	}

	cur := sizes[j]

	total := 0

	if len(row)-i > cur && !strings.ContainsRune(row[i:i+cur], '.') && row[i+cur] != '#' {
		total += rec(row, sizes, i+cur+1, j+1, memo) // + 1 to account for space
	}

	if row[i] != '#' {
		total += rec(row, sizes, i+1, j, memo)
	}

	memo[[2]int{i, j}] = total
	return total
}
