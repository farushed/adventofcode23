package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ScratchCard struct {
	winning, have map[int]interface{} // acts as a set
}

func parse() []ScratchCard {
	scanner := bufio.NewScanner(os.Stdin)

	var cards []ScratchCard

	for scanner.Scan() {
		line := scanner.Text()

		ll := strings.FieldsFunc(line, func(c rune) bool { return c == ':' || c == '|' })
		var cardNum int
		_, _ = fmt.Sscanf(ll[0], "Card %d", &cardNum)

		card := ScratchCard{}

		card.winning = make(map[int]interface{})
		for _, n := range strings.Fields(ll[1]) {
			w, _ := strconv.Atoi(n)
			card.winning[w] = struct{}{}
		}

		card.have = make(map[int]interface{})
		for _, n := range strings.Fields(ll[2]) {
			h, _ := strconv.Atoi(n)
			card.have[h] = struct{}{}
		}

		cards = append(cards, card)
	}

	return cards
}
