package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type MapLine struct {
	dest, source, len int
}

type Map []MapLine

func parse() (seeds []int, maps []Map) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	seedLine := scanner.Text()
	for _, seed := range strings.Fields(seedLine)[1:] { // ignore "seeds:"
		n, _ := strconv.Atoi(seed)
		seeds = append(seeds, n)
	}

	map_ := make(Map, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasSuffix(line, "map:") {
			if len(map_) != 0 {
				maps = append(maps, map_)
			}
			map_ = make(Map, 0)
		} else if len(line) != 0 {
			mapLine := MapLine{}
			nums := strings.Fields(line)
			mapLine.dest, _ = strconv.Atoi(nums[0])
			mapLine.source, _ = strconv.Atoi(nums[1])
			mapLine.len, _ = strconv.Atoi(nums[2])
			map_ = append(map_, mapLine)
		}
	}
	maps = append(maps, map_)

	// fmt.Println(seeds, maps)
	return
}
