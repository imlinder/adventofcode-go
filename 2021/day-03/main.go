package main

import (
  "bufio"
  "fmt"
  "log"
  "math"
  "os"
  "strconv"
)

func main() {
  lines := parseFile("input.txt")
  fmt.Printf("Part 1: %v \n", part1(lines))
  fmt.Printf("Part 2: %v \n", part2(lines))
}

func part1(lines []string) int {
  gBit, eBit := "", ""
  bitLen := len(lines[0])
  bitCount := make([]int, bitLen)

  for _, l := range lines {
    for n, c := range l {
      switch c {
        case 48:
          bitCount[n]--
        case 49:
          bitCount[n]++
      }
    }
  }

  for _, b := range bitCount {
    if math.Signbit(float64(b)) {
      gBit = gBit + "1"
      eBit = eBit + "0"
    } else {
      gBit = gBit + "0"
      eBit = eBit + "1"
    }
  }

  gamma, err := strconv.ParseInt(gBit, 2, 64)

  if err != nil {
    log.Fatal(err)
  }

  epsilon, err := strconv.ParseInt(eBit, 2, 64)

  if err != nil {
    log.Fatal(err)
  }

  return int(gamma) * int(epsilon)
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
