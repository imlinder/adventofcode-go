package main

import (
	"fmt"
	"github.com/imlinder/adventofcode-go/util"
	// "slices"
	// "strconv"
	// "strings"
)

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Printf("Part 1: %v \n", part1(lines))
	fmt.Printf("Part 2: %v \n", part2(lines))
}

func part1(lines []string) int {

	sum := 0

	for linePos, line := range lines {
		for charPos, char := range line {

			if char != 'X' {
				continue
			}

			// Horizontal L->R
			if charPos+3 < len(line) {
				if line[charPos+1] == 'M' && line[charPos+2] == 'A' && line[charPos+3] == 'S' {
					sum++
				}
			}

			// Horizontal R->L
			if charPos-3 >= 0 {
				if line[charPos-1] == 'M' && line[charPos-2] == 'A' && line[charPos-3] == 'S' {
					sum++
				}
			}

			// Vertical T->B
			if linePos+3 < len(lines) {
				if lines[linePos+1][charPos] == 'M' && lines[linePos+2][charPos] == 'A' && lines[linePos+3][charPos] == 'S' {
					sum++
				}
			}

			// Vertical B->T
			if linePos-3 >= 0 {
				if lines[linePos-1][charPos] == 'M' && lines[linePos-2][charPos] == 'A' && lines[linePos-3][charPos] == 'S' {
					sum++
				}
			}

			// Diagonal NW -> SE
			if linePos+3 < len(lines) && charPos+3 < len(line) {
				if lines[linePos+1][charPos+1] == 'M' && lines[linePos+2][charPos+2] == 'A' && lines[linePos+3][charPos+3] == 'S' {
					sum++
				}
			}

			// Diagonal NE -> SW
			if linePos+3 < len(lines) && charPos-3 >= 0 {
				if lines[linePos+1][charPos-1] == 'M' && lines[linePos+2][charPos-2] == 'A' && lines[linePos+3][charPos-3] == 'S' {
					sum++
				}
			}

			// Diagonal SW -> NE
			if linePos-3 >= 0 && charPos+3 < len(line) {
				if lines[linePos-1][charPos+1] == 'M' && lines[linePos-2][charPos+2] == 'A' && lines[linePos-3][charPos+3] == 'S' {
					sum++
				}
			}

			// Diagonal SE -> NW
			if linePos-3 >= 0 && charPos-3 >= 0 {
				if lines[linePos-1][charPos-1] == 'M' && lines[linePos-2][charPos-2] == 'A' && lines[linePos-3][charPos-3] == 'S' {
					sum++
				}
			}
		}
	}

	return sum
}

func part2(lines []string) int {

	sum := 0

	for linePos, line := range lines {
		for charPos, char := range line {

			if char != 'A' {
				continue
			}

			if linePos-1 < 0 || linePos+1 >= len(lines) {
				continue
			}

			if charPos-1 < 0 || charPos+1 >= len(line) {
				continue
			}

			nw := lines[linePos-1][charPos-1]
			ne := lines[linePos-1][charPos+1]
			sw := lines[linePos+1][charPos-1]
			se := lines[linePos+1][charPos+1]

			if ((nw == 'M' && se == 'S') || (nw == 'S' && se == 'M')) && ((ne == 'M' && sw == 'S') || (ne == 'S' && sw == 'M')) {
				sum++
			}

		}
	}

	return sum
}
