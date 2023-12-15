//go:build part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type lens struct {
	focalLength int
	label       string
	next        *lens
}

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	boxes := [256]lens{}
	for i := range boxes {
		boxes[i] = lens{}
	}

	for _, step := range strings.Split(line, ",") {
		if step[len(step)-1] == '-' {
			stepLabel := step[:len(step)-1]

			for prev := &boxes[hash(stepLabel)]; prev.next != nil; prev = prev.next {
				cur := prev.next
				if cur.label == stepLabel {
					prev.next = cur.next
					break
				}
			}
		} else {
			stepLabel := step[:len(step)-2]
			stepFocalLength := int(step[len(step)-1] - '0')

			changed := false
			prev := &boxes[hash(stepLabel)]
			for ; prev.next != nil; prev = prev.next {
				cur := prev.next
				if cur.label == stepLabel {
					cur.focalLength = stepFocalLength
					changed = true
					break
				}
			}
			if !changed {
				prev.next = &lens{stepFocalLength, stepLabel, nil}
			}
		}
	}

	power := 0

	for i, dummy := range boxes {
		for slot, l := 1, dummy.next; l != nil; slot, l = slot+1, l.next {
			power += (i + 1) * slot * l.focalLength
		}
	}

	fmt.Println(power)
}

func hash(str string) int {
	val := 0
	for _, c := range str {
		val = (val + int(c)) * 17 % 256
	}
	return val
}
