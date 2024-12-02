package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	lines := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	want := 2
	result := part1(lines)

	if result != want {
		t.Fatalf("part1() = %v, want %v", result, want)
	}
}

func TestPart2(t *testing.T) {
	lines := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	want := 4
	result := part2(lines)

	if result != want {
		t.Fatalf("part2() = %v, want %v", result, want)
	}
}
