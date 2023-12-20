//go:build part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Action struct {
	signal bool
	module string
}

const (
	noType = iota
	flipFlop
	conjunction
)

var (
	connections    = make(map[string][]string)
	types          = make(map[string]int)
	flipFlopStates = make(map[string]bool)
	conStates      = make(map[string]map[string]bool)
)

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		ln := strings.Split(line, " -> ")
		name := ln[0]
		if name[0] == '%' {
			name = name[1:]
			types[name] = flipFlop
		} else if name[0] == '&' {
			name = name[1:]
			types[name] = conjunction
		}
		connections[name] = strings.Split(ln[1], ", ")
	}

	for m := range connections {
		if types[m] == conjunction {
			conStates[m] = make(map[string]bool)
		}
	}
	for m, c := range connections {
		for _, m2 := range c {
			if types[m2] == conjunction {
				conStates[m2][m] = false
			}
		}
	}

	cycles := make(map[string]int)
	var preRxName string
	for name, conns := range connections {
		if len(conns) == 1 && conns[0] == "rx" {
			preRxName = name
			break
		}
	}
	for name, conns := range connections {
		if len(conns) == 1 && conns[0] == preRxName {
			cycles[name] = 0
		}
	}
	numCycles := 0

Outer:
	for i := 0; i < 100000; i++ {
		actions := []Action{
			{false, "broadcaster"},
		}

		for len(actions) > 0 {
			l := len(actions)
			for j := 0; j < l; j++ {
				a := actions[j]

				if v, ok := cycles[a.module]; ok && a.signal && v == 0 {
					cycles[a.module] = i + 1
					numCycles++
					if numCycles == len(cycles) {
						break Outer
					}
				}

				res := send(a)
				actions = append(actions, res...)
			}
			actions = actions[l:]
		}
	}

	fmt.Println(cycles)

	lcm := 1
	for _, x := range cycles {
		lcm = (lcm * x) / gcd(lcm, x)
	}

	fmt.Println(lcm)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func send(action Action) []Action {
	var result []Action

	if action.signal {
		for _, new := range connections[action.module] {
			// flipflops ignored
			if types[new] == conjunction {
				conStates[new][action.module] = true
				allTrue := true
				for _, v := range conStates[new] {
					if !v {
						allTrue = false
						break
					}
				}

				// send low if all inputs high
				result = append(result, Action{!allTrue, new})
			}
		}
	} else {

		for _, new := range connections[action.module] {
			// fmt.Println(action, new)
			switch types[new] {
			case flipFlop:
				flipFlopStates[new] = !flipFlopStates[new]
				result = append(result, Action{flipFlopStates[new], new}) // send high if was off, now on
			case conjunction:
				conStates[new][action.module] = false
				// since at least one input is now low, send high
				result = append(result, Action{true, new})
			}
		}
	}

	return result
}
