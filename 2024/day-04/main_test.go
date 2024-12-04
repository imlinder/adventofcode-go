package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	lines := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	want := 18
	result := part1(lines)

	if result != want {
		t.Fatalf("part1() = %v, want %v", result, want)
	}
}

func TestPart2(t *testing.T) {
	lines := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	want := 9
	result := part2(lines)

	if result != want {
		t.Fatalf("part2() = %v, want %v", result, want)
	}
}
