package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/imlinder/adventofcode-go/util"
	"strconv"
	// "strings"
)

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Printf("Part 1: %v \n", part1(lines))
	fmt.Printf("Part 2: %v \n", part2(lines))
}

func part1(lines []string) int {
	var max = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var sum int

	for _, l := range lines {
		if len(l) < 1 {
			continue
		}
		digitRe := regexp.MustCompile(`\d+`)
		var spl []string = strings.Split(l, ":")
		gameId, err := strconv.Atoi(digitRe.FindString(spl[0]))

		if err != nil {
			panic(err)
		}

		games := strings.Split(spl[1], ";")

		var pass = true
		for _, game := range games {
			hands := strings.Split(game, ",")
			for _, h := range hands {
				spl := strings.Split(strings.TrimSpace(h), " ")
				samount, color := spl[0], spl[1]
				amount, err := strconv.Atoi(samount)
				if err != nil {
					panic(err)
				}
				if amount > max[color] {
					pass = false
				}
			}
		}
		if pass {
			sum += gameId
		}

	}
	return sum
}

func part2(lines []string) int {
	return 0
}
