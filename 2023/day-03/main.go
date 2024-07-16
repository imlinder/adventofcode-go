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

/*
* Return x if it's positive, otherwise 0.
 */
func MinZero(x int) int {
	if x < 0 {
		return 0
	}
	return x
}

/**
* Return the smallest integer
 */
func MinInt(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

/**
* Part 2
 */
func part2(lines []string) int {
	sum := 0
	symbRe := regexp.MustCompile(`\*`)
	for i, line := range lines {
		matches := symbRe.FindAllStringIndex(line, -1)
		for _, match := range matches {
			adjacent := findAdjacentNums(i, match[0], lines)
			if len(adjacent) == 2 {
				sum += stringToInt(adjacent[0]) * stringToInt(adjacent[1])
			}
		}
	}
	return sum
}

/*
* Find adjacent (left/right/diagonal) numbers to index.
 */
func findAdjacentNums(lineIndex int, index int, lines []string) []string {
	currentLine := lines[lineIndex]
	prevLine := ""
	nextLine := ""
	adjacent := []string{}

	if lineIndex > 0 {
		prevLine = lines[lineIndex-1]
	}

	if lineIndex < len(lines) {
		nextLine = lines[lineIndex+1]
	}

	adjacent = append(adjacent, findNumsAtIndex(prevLine, index, 1)...)
	adjacent = append(adjacent, findNumsAtIndex(currentLine, index, 1)...)
	adjacent = append(adjacent, findNumsAtIndex(nextLine, index, 1)...)

	return adjacent
}

/**
* Find numbers in line that is covering index + padding
 */
func findNumsAtIndex(line string, index int, padding int) []string {
	numRe := regexp.MustCompile(`\d+`)
	matches := numRe.FindAllStringIndex(line, -1)
	nums := []string{}
	for _, match := range matches {
		if index >= match[0]-padding && index < match[1]+padding {
			nums = append(nums, line[match[0]:match[1]])
		}
	}

	return nums
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
