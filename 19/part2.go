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

type PartRange struct {
	x1, x2 int
	m1, m2 int
	a1, a2 int
	s1, s2 int
	empty  bool
}

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	var workflows [][]string
	workflowIdxs := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		i := strings.Index(line, "{")
		name := line[:i]

		workflowIdxs[name] = len(workflows)
		workflows = append(workflows, strings.Split(line[i+1:len(line)-1], ","))
	}

	pr := PartRange{1, 4000, 1, 4000, 1, 4000, 1, 4000, false}

	res := check("in", pr, workflowIdxs, workflows)

	fmt.Println(res)
}

func check(start string, pr PartRange, workflowIdxs map[string]int, workflows [][]string) int {
	idx := workflowIdxs[start]

	result := 0

	for _, rule := range workflows[idx] {
		r := strings.Split(rule, ":")
		if len(r) == 1 {
			ret := r[0]

			if ret == "A" {
				result += calc(pr)
			} else if ret != "R" && ret != "" {
				result += check(ret, pr, workflowIdxs, workflows)
			}

		} else {
			attr := r[0][0]
			comp := r[0][1]
			compVal, _ := strconv.Atoi(r[0][2:])
			ret := r[1]

			var prCur PartRange
			prCur, pr = intersect(pr, attr, comp, compVal)

			if ret == "A" {
				result += calc(prCur)
			} else if ret != "R" && ret != "" {
				result += check(ret, prCur, workflowIdxs, workflows)
			}
		}
	}

	return result
}

func calc(pr PartRange) int {
	return (pr.x2 - pr.x1 + 1) *
		(pr.m2 - pr.m1 + 1) *
		(pr.a2 - pr.a1 + 1) *
		(pr.s2 - pr.s1 + 1)
}

func intersect(orig PartRange, attr, comp byte, compVal int) (PartRange, PartRange) {
	new := orig

	offset := 0
	if comp == '>' {
		offset = 1
	}

	switch attr {
	case 'x':
		if orig.x2 < compVal {
			new.empty = true
		} else if compVal < orig.x1 {
			orig.empty = true
		} else {
			new.x1 = compVal + offset
			orig.x2 = compVal + offset - 1
		}

	case 'm':
		if orig.m2 < compVal {
			new.empty = true
		} else if compVal < orig.m1 {
			orig.empty = true
		} else {
			new.m1 = compVal + offset
			orig.m2 = compVal + offset - 1
		}
	case 'a':
		if orig.a2 < compVal {
			new.empty = true
		} else if compVal < orig.a1 {
			orig.empty = true
		} else {
			new.a1 = compVal + offset
			orig.a2 = compVal + offset - 1
		}
	case 's':
		if orig.s2 < compVal {
			new.empty = true
		} else if compVal < orig.s1 {
			orig.empty = true
		} else {
			new.s1 = compVal + offset
			orig.s2 = compVal + offset - 1
		}
	}

	if comp == '>' {
		return new, orig
	}
	return orig, new
}
