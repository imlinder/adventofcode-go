package main

import (
	"fmt"
	"strings"

	"github.com/imlinder/adventofcode-go/util"
)

func main() {
	input := util.ReadFile("input.txt")
	part1 := findMarker(input[0], 4)
	part2 := findMarker(input[0], 14)
	fmt.Printf("Part 1: %v \n", part1)
	fmt.Printf("Part 2: %v \n", part2)
}

func findMarker(input string, markerLength int) int {
	for i := 0; i < len(input)-markerLength-1; i++ {
		if !hasDuplicate(input[i : i+markerLength]) {
			return i + markerLength
		}
	}
	return 0
}

func hasDuplicate(input string) bool {
	for _, c := range input {
		if strings.Count(input, string(c)) > 1 {
			return true
		}
	}
	return false
}
