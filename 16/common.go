package main

import (
	"bufio"
	"os"
)

type Coord [2]int

type Beam struct {
	dir Coord
	pos Coord
}

var (
	left  = Coord{0, -1}
	right = Coord{0, 1}
	up    = Coord{-1, 0}
	down  = Coord{1, 0}
)

// for each tile type, for each incoming direction, a list of outgoing directions (if any change)
var beamTranslations = map[rune]map[Coord][]Coord{
	'.': {},
	'/': {
		left:  {down},
		right: {up},
		up:    {right},
		down:  {left},
	},
	'\\': {
		left:  {up},
		right: {down},
		up:    {left},
		down:  {right},
	},
	'|': {
		left:  {up, down},
		right: {up, down},
	},
	'-': {
		down: {left, right},
		up:   {left, right},
	},
}

func parse() [][]rune {
	scanner := bufio.NewScanner(os.Stdin)
	var grid [][]rune

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	return grid
}

func countEnergised(input Beam, grid [][]rune) int {
	energised := make(map[Coord]interface{})
	beams := make(map[Beam]interface{})

	step(input, grid, beams, energised)

	return len(energised)
}

func step(beam Beam, grid [][]rune, beams map[Beam]interface{}, energised map[Coord]interface{}) {
	if beam.pos[0] < 0 || beam.pos[0] >= len(grid) || beam.pos[1] < 0 || beam.pos[1] >= len(grid[0]) {
		return
	}

	if _, ok := beams[beam]; ok {
		return // already stepped a beam from here
	}
	beams[beam] = struct{}{}
	energised[beam.pos] = struct{}{}

	if newDirs, ok := beamTranslations[grid[beam.pos[0]][beam.pos[1]]][beam.dir]; ok {
		for _, newDir := range newDirs {
			step(next(beam, newDir), grid, beams, energised)
		}
	} else {
		step(next(beam, beam.dir), grid, beams, energised)
	}
}

func next(beam Beam, dir Coord) Beam {
	return Beam{dir, [2]int{beam.pos[0] + dir[0], beam.pos[1] + dir[1]}}
}
