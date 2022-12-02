package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	lines := []string{
		"A Y",
		"B X",
		"C Z",
	}
	want := 15
	result := part1(lines)

	if result != want {
		t.Fatalf(`part1(lines) = %v, want %v`, result, want)
	}
}

func TestPart2(t *testing.T) {
	lines := []string{
		"A Y",
		"B X",
		"C Z",
	}
	want := 12
	result := part2(lines)

	if result != want {
		t.Fatalf(`part2(lines) = %v, want %v`, result, want)
	}
}
