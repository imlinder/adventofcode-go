package main
import (
  "testing"
)

func TestPart1(t *testing.T) {
  lines := []string {
    "00100",
    "11110",
    "10110",
    "10111",
    "10101",
    "01111",
    "00111",
    "11100",
    "10000",
    "11001",
    "00010",
    "01010",
  }

  want := 198

  result := part1(lines)

  if result != want {
    t.Fatalf(`part1(lines) = %v, want %#v`, result, want)
  }
}

