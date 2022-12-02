package main

import (
	"fmt"
	"github.com/imlinder/adventofcode-go/util"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Printf("Part 1: %v \n", part1(lines))
	fmt.Printf("Part 2: %v \n", part2(lines))
}

// Rock     A X
// Paper    B Y
// Scissors C Z

func part1(lines []string) int {
	score := 0
	for _, l := range lines {
		if len(l) < 1 {
			continue
		}
		g := strings.Split(l, " ")
		myShape := decodeShape(g[1])
		score += shapeScore(myShape)
		score += roundScore(myShape, g[0])
	}

	return score
}

func part2(lines []string) int {
	score := 0
	for _, l := range lines {
		if len(l) < 1 {
			continue
		}
		g := strings.Split(l, " ")
		myShape := decodeShape2(g[0], g[1])
		score += shapeScore(myShape)
		score += roundScore(myShape, g[0])
	}

	return score
}

func roundScore(s1 string, s2 string) int {
	// Tie
	if s1 == s2 {
		return 3
	}
	// Rock - Paper = loss
	if s1 == "A" && s2 == "B" {
		return 0
	}
	// Rock - Scissors = win
	if s1 == "A" && s2 == "C" {
		return 6
	}
	// Paper - Rock = win
	if s1 == "B" && s2 == "A" {
		return 6
	}
	// Paper - Scissors = loss
	if s1 == "B" && s2 == "C" {
		return 0
	}
	// Scissors - Rock = loss
	if s1 == "C" && s2 == "A" {
		return 0
	}
	// Scissors - Paper = win
	if s1 == "C" && s2 == "B" {
		return 6
	}
	return 0
}

func shapeScore(shape string) int {
	switch shape {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	default:
		return 0
	}
}

func decodeShape(shape string) string {
	switch shape {
	case "X":
		return "A"
	case "Y":
		return "B"
	case "Z":
		return "C"
	default:
		return shape
	}
}

func decodeShape2(shape string, outcome string) string {
	returnShape := shape
	switch outcome {
	// Needs to lose
	case "X":
		if shape == "A" {
			returnShape = "C"
		}
		if shape == "B" {
			returnShape = "A"
		}
		if shape == "C" {
			returnShape = "B"
		}
	// Needs to draw
	case "Y":
		returnShape = shape
	// Needs to win
	case "Z":
		if shape == "A" {
			returnShape = "B"
		}
		if shape == "B" {
			returnShape = "C"
		}
		if shape == "C" {
			returnShape = "A"
		}
	default:
		returnShape = shape
	}
	return returnShape
}
