package main

import (
	"bufio"
	"os"
)

func solve(expansion int) int {
	scanner := bufio.NewScanner(os.Stdin)

	prefixCols := []int{0}
	prefixRows := []int{0}

	var galaxies [][2]int

	colHasGalaxy := make(map[int]interface{})

	var line string
	row := 0
	for scanner.Scan() {
		line = scanner.Text()

		rowHasGalaxy := false
		for col, c := range line {
			if c == '#' {
				rowHasGalaxy = true
				colHasGalaxy[col] = struct{}{}
				galaxies = append(galaxies, [2]int{row, col})
			}
		}

		pre := prefixRows[len(prefixRows)-1]
		new := pre + 1
		if !rowHasGalaxy {
			new += expansion
		}
		prefixRows = append(prefixRows, new)

		row++
	}

	for i := 0; i < len(line); i++ {
		pre := prefixCols[len(prefixCols)-1]
		new := pre + 1
		if _, ok := colHasGalaxy[i]; !ok {
			new += expansion
		}
		prefixCols = append(prefixCols, new)
	}

	total := 0
	for _, g1 := range galaxies {
		for _, g2 := range galaxies {
			y := prefixRows[g1[0]] - prefixRows[g2[0]]
			if y < 0 {
				y = -y
			}
			x := prefixCols[g1[1]] - prefixCols[g2[1]]
			if x < 0 {
				x = -x
			}
			total += x + y
		}
	}
	total >>= 1 // account for double counting

	return total
}
