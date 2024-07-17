package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/imlinder/adventofcode-go/util"
)

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Printf("Part 1: %v \n", part1(lines))
	fmt.Printf("Part 2: %v \n", part2(lines))
}

func part1(lines []string) int {
	var sum int

	for _, l := range lines {
		var points int
		wins := getNumOfWins(l)
		if wins > 0 {
			points = 1
			for i := 1; i < wins; i++ {
				points = points + points
			}
		}
		sum += points
	}

	return sum
}

func part2(lines []string) int {
	var sum int
	// Slice to store all extra cards in
	cardCount := make([]int, len(lines))
	for i, l := range lines {
		// Number of copies won for for this card
		extraCards := cardCount[i]

		// We now now these will add to the result
		sum += 1 + extraCards

		// Add the number of copies won for the cards following
		wins := getNumOfWins(l)
		for x := i + 1; x <= i+wins; x++ {
			if len(cardCount) >= x {
				cardCount[x] += 1 + extraCards
			}
		}
	}
	return sum
}

/*
* Split a scratchcard to three parts:
* The card title, the winning numbers and the numbers you have.
 */
func splitCard(line string) []string {
	// Remove double whitespace
	line = strings.ReplaceAll(line, "  ", " ")
	// Split by : and |
	split := regexp.MustCompile("[\\:\\|]+").Split(line, -1)
	return split
}

func getNumOfWins(card string) int {
	c := splitCard(card)
	winningNums := strings.Split(c[1], " ")
	playedNums := strings.Split(c[2], " ")

	var points int

	for _, n := range playedNums {
		for _, w := range winningNums {
			if n == w && len(n) > 0 {
				points++
			}
		}
	}

	return points

}
