package main

import (
	"fmt"
	"github.com/imlinder/adventofcode-go/util"
	"regexp"
	"strconv"
)

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Printf("Part 1: %v \n", part1(lines))
}

func part1(lines []string) int {
	sum := 0
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		re := regexp.MustCompile(`\d`)
		matches := re.FindAllString(l, -1)
		num, _ := strconv.Atoi(matches[0] + matches[len(matches)-1])
		sum += num
	}

	return sum
}
