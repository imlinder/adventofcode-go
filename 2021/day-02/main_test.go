package main
import (
  "testing"
)

func TestPart1(t *testing.T) {
  lines := []string {
    "forward 5",
    "down 5",
    "forward 8",
    "up 3",
    "down 8",
    "forward 2",
  }

  want := 150

  result := part1(lines)

  if result != want {
    t.Fatalf(`part1(lines) = %v, want %#v`, result, want)
  }
}

func TestPart2 (t *testing.T) {
  lines := []string {
    "forward 5",
    "down 5",
    "forward 8",
    "up 3",
    "down 8",
    "forward 2",
  }

  want := 900

  result := part2(lines)

  if result != want {
    t.Fatalf(`part2(lines) = %v, want %#v`, result, want)
  }
}
