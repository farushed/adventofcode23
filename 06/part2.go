//go:build part2

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

	t, _ := strconv.Atoi(strings.Join(times, ""))
	d, _ := strconv.Atoi(strings.Join(distances, ""))

	lower := math.Floor(
		(-float64(t)+math.Sqrt(float64(t*t-4*d)))/-2,
	) + 1
	upper := math.Ceil(
		(-float64(t)-math.Sqrt(float64(t*t-4*d)))/-2,
	) - 1

	fmt.Println(t, d, upper, lower)

	fmt.Println(int(upper - lower + 1))
}
