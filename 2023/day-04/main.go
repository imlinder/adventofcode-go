package main

import (
	"fmt"
	"github.com/imlinder/adventofcode-go/util"
	"regexp"
	"strings"
)

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Printf("Part 1: %v \n", part1(lines))
	fmt.Printf("Part 2: %v \n", part2(lines))
}

func part1(lines []string) int {
	var sum int

	for _, l := range lines {
		if len(l) < 1 {
			continue
		}
		card := splitCard(l)
		winningNums := strings.Split(card[1], " ")
		playedNums := strings.Split(card[2], " ")

		var points int

		for _, n := range playedNums {
			for _, w := range winningNums {
				if n == w && len(n) > 0 {
					if points == 0 {
						points = 1
					} else {
						points = points * 2
					}
				}
			}
		}

		sum += points
	}

	return sum
}

func part2(lines []string) int {
	return 0
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
