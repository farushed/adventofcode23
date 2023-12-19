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

type Part struct {
	x, m, a, s int
}

type Rule func(Part) string
type Workflow []Rule

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	var workflows []Workflow
	workflowIdxs := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		i := strings.Index(line, "{")
		name := line[:i]

		var workflow Workflow

		for _, r := range strings.Split(line[i+1:len(line)-1], ",") {
			var f Rule

			r2 := strings.Split(r, ":")
			if len(r2) == 1 {
				ret := r2[0]
				f = func(p Part) string { return ret }
			} else {
				attr := r2[0][0]
				comp := r2[0][1]
				compVal, _ := strconv.Atoi(r2[0][2:])
				ret := r2[1]

				f = func(p Part) string {
					var val int
					switch attr {
					case 'x':
						val = p.x
					case 'm':
						val = p.m
					case 'a':
						val = p.a
					case 's':
						val = p.s
					}
					if comp == '>' {
						if val > compVal {
							return ret
						}
					} else {
						if val < compVal {
							return ret
						}
					}
					return ""
				}
			}

			workflow = append(workflow, f)
		}
		workflowIdxs[name] = len(workflows)
		workflows = append(workflows, workflow)

	}

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		part := Part{}
		fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d}", &part.x, &part.m, &part.a, &part.s)

		total += checkPart(part, workflowIdxs, workflows)
	}

	fmt.Println(total)
}

func checkPart(part Part, workflowIdxs map[string]int, workflows []Workflow) int {
	idx := workflowIdxs["in"]
	for idx < len(workflows) {

		workflow := workflows[idx]
	Wkflw:
		for _, rule := range workflow {
			switch res := rule(part); res {
			case "":
				continue
			case "A":
				return part.x + part.m + part.a + part.s
			case "R":
				return 0
			default:
				idx = workflowIdxs[res]
				break Wkflw
			}
		}

	}

	return 0
}
