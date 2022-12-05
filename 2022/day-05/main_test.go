package main

import (
	"github.com/imlinder/adventofcode-go/util"
	"testing"
)

func TestPart1(t *testing.T) {
	lines := util.ReadFile("testinput.txt")
	want := "CMZ"
	result := part1(lines)

	if result != want {
		t.Fatalf(`part1(lines) = %v, want %v`, result, want)
	}
}

func TestPart2(t *testing.T) {
	lines := util.ReadFile("testinput.txt")
	want := "MCD"
	result := part2(lines)

	if result != want {
		t.Fatalf(`part2(lines) = %v, want %v`, result, want)
	}
}
