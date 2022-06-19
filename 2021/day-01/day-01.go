package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
)

func main() {
  f, err := os.Open("day-01-input.txt")

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

  fmt.Printf("Increments: %v", increments)

}
