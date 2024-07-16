package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	want := 4361
	result := part1(lines)

	if result != want {
		t.Fatalf("part1() = %v, want %v", result, want)
	}
}

func TestPart2(t *testing.T) {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	want := 467835
	result := part2(lines)

	if result != want {
		t.Fatalf("part2() = %v, want %v", result, want)
	}
}
