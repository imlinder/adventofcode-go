package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/imlinder/adventofcode-go/util"
)

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Printf("Part 1: %v \n", part1(lines))
	fmt.Printf("Part 2: %v \n", part2(lines))
}

func part1(lines []string) int {
	var sum int

	numRe := regexp.MustCompile(`\d+`)
	symbRe := regexp.MustCompile(`[^\d\.]`)

	for i, line := range lines {

		if len(line) < 1 {
			continue
		}

		numbers := numRe.FindAllStringIndex(line, -1)

		for _, n := range numbers {
			var match bool
			number, err := strconv.Atoi(line[n[0]:n[1]])
			if err != nil {
				panic(err)
			}

			start, end := MinZero(n[0]-1), MinInt(len(line), n[1]+1)
			currentLine := line[start:end]
			m := symbRe.MatchString(currentLine)

			if m {
				match = true
			}

			if i > 0 && !match {
				prevLine := lines[i-1][start:end]
				m := symbRe.MatchString(prevLine)
				if m {
					match = true
				}
			}

			if i < len(lines)-1 && !match {
				nextLine := lines[i+1]
				if len(nextLine) > 0 {
					nextLine = nextLine[start:end]
					m := symbRe.MatchString(nextLine)
					if m {
						match = true
					}
				}
			}

			if match {
				sum += number
			}
		}
	}

	return sum
}

func part2(lines []string) int {
	return 0
}

func MinZero(x int) int {
	if x < 0 {
		return 0
	}
	return x
}

func MinInt(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
