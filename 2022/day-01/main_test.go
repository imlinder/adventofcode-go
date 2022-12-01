package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	lines := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}
	want := 24000
	result := part1(lines)

	if result != want {
		t.Fatalf(`part1(lines) = %v, want %v`, result, want)
	}
}

func TestPart2(t *testing.T) {
	lines := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}
	want := 45000
	result := part2(lines)

	if result != want {
		t.Fatalf(`part2(lines) = %v, want %v`, result, want)
	}
}
