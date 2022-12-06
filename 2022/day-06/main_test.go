package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	want := 7
	result := findMarker(input, 4)

	if result != want {
		t.Fatalf(`part1(lines) = %v, want %v`, result, want)
	}
}

func TestPart2(t *testing.T) {
	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	want := 19
	result := findMarker(input, 14)

	if result != want {
		t.Fatalf(`part2(lines) = %v, want %v`, result, want)
	}
}
