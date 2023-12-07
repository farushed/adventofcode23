//go:build part1

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	times := strings.Fields(scanner.Text())[1:]
	scanner.Scan()
	distances := strings.Fields(scanner.Text())[1:]

	product := 1

	for i, time := range times {
		t, _ := strconv.Atoi(time)
		d, _ := strconv.Atoi(distances[i])

		// (t-c)*c > d
		// tc - c*2 > d
		// tc - c*2 - d > 0
		// n-shaped quadratic, >0 between roots
		// roots = (-t +- sqrt(t^2 - 4d)) / -2
		// then take inner results
		// instead of ceil then floor, use floor+1 and ceil-1 since if the root is an int, we want the next/prev int

		lower := math.Floor(
			(-float64(t)+math.Sqrt(float64(t*t-4*d)))/-2,
		) + 1
		upper := math.Ceil(
			(-float64(t)-math.Sqrt(float64(t*t-4*d)))/-2,
		) - 1

		fmt.Println(t, d, upper, lower)
		product *= int(upper - lower + 1)
	}

	fmt.Println(product)
}
