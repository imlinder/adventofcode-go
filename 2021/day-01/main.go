package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
)

func main() {
  fmt.Printf("Part 1: %v \n", part1())
  fmt.Printf("Part 2: %v \n", part2())
}

func part1() int {
  f, err := os.Open("input.txt")

  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()

  scanner := bufio.NewScanner(f)

  var increments int
  var prevDepth int

  for scanner.Scan() {
    depth, err := strconv.Atoi(scanner.Text())

    if err != nil {
      log.Fatal(err)
    }

    if prevDepth != 0 && depth > prevDepth {
      increments++
    }

    prevDepth = depth

  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return increments
}

func part2() int {
  // Number of data points to use as sum for sliding window.
  const WinSize = 3

  f, err := os.Open("input.txt")

  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()

  var increments int
  var prevDepth int
  var lines []int

  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    depth, err := strconv.Atoi(scanner.Text())

    if err != nil {
      log.Fatal(err)
    }

    lines = append(lines, depth)
  }

  for i := 0; i <= len(lines) - WinSize; i++ {
    var winDepth int
    for j := 0; j < WinSize; j++ {
      winDepth = winDepth + lines[i + j]
    }

    if prevDepth != 0 && winDepth > prevDepth {
      increments++
    }

    prevDepth = winDepth

  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return increments
}
