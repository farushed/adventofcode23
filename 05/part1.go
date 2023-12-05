//go:build part1

package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	vals, maps := parse()

	for _, map_ := range maps {
		var newVals []int
		for _, val := range vals {

			added := false
			for _, mapLine := range map_ {
				if val >= mapLine.source && val < mapLine.source+mapLine.len {
					newVals = append(newVals, val-mapLine.source+mapLine.dest)
					added = true
				}
			}
			if !added {
				newVals = append(newVals, val)
			}
		}
		vals = newVals
	}

	lowest := vals[0]
	for _, val := range vals[1:] {
		if val < lowest {
			lowest = val
		}
	}

	fmt.Println(lowest)
}
