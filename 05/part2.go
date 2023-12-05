//go:build part2

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
		// split up ranges, then add them back to the list. so keep looking through the list of ranges till they're all processed
		for len(vals) > 0 {
			valLen := len(vals)
			for i := 0; i < valLen/2; i++ {
				seedStart := vals[i*2]
				seedEnd := seedStart + vals[i*2+1] - 1

				added := false
				for _, ml := range map_ {

					overlap := false
					if seedStart < ml.source && seedEnd >= ml.source {
						vals = append(vals, seedStart, ml.source-seedStart) // split up range
						seedStart = ml.source
						overlap = true
					}
					if seedEnd > ml.source+ml.len-1 && seedStart <= ml.source+ml.len-1 {
						vals = append(vals, ml.source+ml.len, seedEnd-(ml.source+ml.len-1)) // split up range
						seedEnd = ml.source + ml.len - 1
						overlap = true
					}
					if overlap || (seedStart >= ml.source && seedEnd < ml.source+ml.len) {
						newVals = append(newVals, seedStart-ml.source+ml.dest, seedEnd-seedStart+1) // remap vals to next set
						added = true
					}
				}
				if !added { // we've looked through all map ranges, cut off our overlaps, but seedStart-seedEnd was never found in the map
					newVals = append(newVals, seedStart, seedEnd-seedStart+1)
				}
			}
			vals = vals[valLen:]
		}

		vals = newVals
	}

	// only care about the lowest of each range
	lowest := vals[0]
	for i := 1; i < len(vals)/2; i++ {
		if vals[i*2] < lowest {
			lowest = vals[i*2]
		}
	}

	fmt.Println(lowest)
}
