package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
  "strings"
  "strconv"
)

func main() {
  lines := parseFile("input.txt")
  fmt.Printf("Part 1: %v \n", part1(lines))
}

func part1(lines []string) int {
  // Horizontal position and depth 
  hor, dep := 0, 0

  for _, l := range lines {

    // Split line into [instruction, distance]
    instr := strings.Split(l, " ")

    // Convert distance to int
    dist, err := strconv.Atoi(instr[1])

    if err != nil {
      log.Fatal(err)
    }

    switch instr[0] {
      case "forward":
        hor += dist
      case "down":
        dep += dist
      case "up":
        dep -= dist
    }
  }

  return hor * dep
}

func parseFile(file string) []string {

  f, err := os.Open(file)

  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()

  var lines []string

  scanner := bufio.NewScanner(f)

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return lines
}
