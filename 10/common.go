package main

import (
	"bufio"
	"os"
)

func parse() ([2]int, map[[2]int][][2]int, int, int) {
	scanner := bufio.NewScanner(os.Stdin)

	connections := make(map[[2]int][][2]int) // map coords to list of coords

	var start [2]int

	var line string
	row := 0
	for scanner.Scan() {
		line = scanner.Text()
		for col, p := range line {
			if p == '.' {
				continue
			}

			cur := [2]int{row, col}

			var out [][2]int
			switch p {
			case 'F':
				out = [][2]int{{row + 1, col}, {row, col + 1}}
			case '7':
				out = [][2]int{{row + 1, col}, {row, col - 1}}
			case 'L':
				out = [][2]int{{row - 1, col}, {row, col + 1}}
			case 'J':
				out = [][2]int{{row - 1, col}, {row, col - 1}}
			case '|':
				out = [][2]int{{row - 1, col}, {row + 1, col}}
			case '-':
				out = [][2]int{{row, col - 1}, {row, col + 1}}

			case 'S':
				start = cur
			}

			connections[cur] = out
		}

		row++
	}

	for _, c := range [][2]int{{start[0] - 1, start[1]}, {start[0] + 1, start[1]}, {start[0], start[1] - 1}, {start[0], start[1] + 1}} {
		if len(connections[c]) == 2 && (start == connections[c][0] || start == connections[c][1]) {
			connections[start] = append(connections[start], c)
		}
	}

	return start, connections, row, len(line)
}
