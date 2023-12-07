//go:build part1

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	high_card = iota
	one_pair
	two_pair
	three_kind
	full_house
	four_kind
	five_kind
)

var (
	cardStrength map[rune]int
	cards        = "23456789TJQKA"
)

func init() {
	cardStrength = make(map[rune]int)
	for i, c := range cards {
		cardStrength[c] = i
	}
}

type hand struct {
	cards string
	type_ int
	bid   int
}

func main() {
	defer func(start time.Time) { fmt.Printf("Time taken %dus\n", time.Since(start).Microseconds()) }(time.Now())

	scanner := bufio.NewScanner(os.Stdin)

	var hands []hand

	for scanner.Scan() {
		x := strings.Split(scanner.Text(), " ")
		b, _ := strconv.Atoi(x[1])

		hands = append(hands, hand{
			cards: x[0],
			bid:   b,
			type_: getHandType(x[0]),
		})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].type_ != hands[j].type_ {
			return hands[i].type_ < hands[j].type_
		}
		for idx, ci := range hands[i].cards {
			cj := []rune(hands[j].cards)[idx]
			if ci != cj {
				return cardStrength[ci] < cardStrength[cj]
			}
		}
		return false
	})
	// fmt.Println(hands)

	total := 0
	for i, h := range hands {
		total += (i + 1) * h.bid
	}

	fmt.Println(total)

}

func getHandType(cards string) int {
	freq := make(map[rune]int)
	for _, c := range cards {
		freq[c]++
	}

	switch len(freq) {
	case 5:
		return high_card
	case 4:
		return one_pair
	case 3:
		for _, f := range freq {
			if f == 3 {
				return three_kind
			}
		}
		return two_pair
	case 2:
		for _, f := range freq {
			if f == 4 {
				return four_kind
			}
		}
		return full_house
	case 1:
		return five_kind
	}

	return -1
}
