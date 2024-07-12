package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	lines := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	want := 142
	result := part1(lines)

	if result != want {
		t.Fatalf("part1() = %v, want %v", result, want)
	}
}

func TestPart2(t *testing.T) {
	lines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	want := 281
	result := part2(lines)

	if result != want {
		t.Fatalf("part2() = %v, want %v", result, want)
	}
}
