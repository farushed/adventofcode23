//go:build part1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	sum := 0

	for _, step := range strings.Split(line, ",") {
		sum += hash(step)
	}

	fmt.Println(sum)
}

func hash(str string) int {
	val := 0
	for _, c := range str {
		val = (val + int(c)) * 17 % 256
	}
	return val
}
