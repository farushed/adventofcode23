package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cubes struct {
	Red, Green, Blue int
}

func parse() map[int][]cubes {
	scanner := bufio.NewScanner(os.Stdin)

	games := make(map[int][]cubes)

	for scanner.Scan() {
		line := scanner.Text()
		gameId, reveals := parseLine(line)
		games[gameId] = reveals
	}

	return games
}

func parseLine(line string) (int, []cubes) {
	x := strings.Split(line, ":")

	var gameId int
	_, _ = fmt.Sscanf(x[0], "Game %d", &gameId)

	var reveals []cubes

	for _, set := range strings.Split(x[1], ";") {
		reveal := cubes{}
		for _, numCol := range strings.Split(set, ",") {
			var num int
			var col string
			_, _ = fmt.Sscanf(numCol, "%d %s", &num, &col)

			switch col {
			case "red":
				reveal.Red = num
			case "green":
				reveal.Green = num
			case "blue":
				reveal.Blue = num
			}
		}
		reveals = append(reveals, reveal)
	}

	return gameId, reveals
}
