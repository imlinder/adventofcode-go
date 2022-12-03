package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	lines := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
	want := 157
	result := part1(lines)

	if result != want {
		t.Fatalf(`part1(lines) = %v, want %v`, result, want)
	}
}
