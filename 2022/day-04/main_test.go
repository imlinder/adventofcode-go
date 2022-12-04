package main

import (
	"testing"
)

func Test(t *testing.T) {
	lines := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
	wantPart1 := 2
	wantPart2 := 4
	resultPart1, resultPart2 := run(lines)

	if resultPart1 != wantPart1 {
		t.Fatalf(`part1(lines) = %v, want %v`, resultPart1, wantPart2)
	}

	if resultPart2 != wantPart2 {
		t.Fatalf(`part1(lines) = %v, want %v`, resultPart1, wantPart2)
	}
}
