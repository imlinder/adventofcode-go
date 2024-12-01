package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	lines := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}

	want := 11
	result := part1(lines)

	if result != want {
		t.Fatalf("part1() = %v, want %v", result, want)
	}
}

func TestPart2(t *testing.T) {
	lines := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}

	want := 31
	result := part2(lines)

	if result != want {
		t.Fatalf("part1() = %v, want %v", result, want)
	}
}
