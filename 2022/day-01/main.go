package main

import (
	"fmt"
	"github.com/imlinder/adventofcode-go/util"
	"log"
	"sort"
	"strconv"
)

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Printf("Part 1: %v \n", part1(lines))
	fmt.Printf("Part 2: %v \n", part2(lines))
}

func part1(lines []string) int {
	highest := 0
	current := 0

	for _, l := range lines {
		// If empty line, we are done with this elf, check if it's the highest
		// value
		if len(l) == 0 {
			if current > highest {
				highest = current
			}
			current = 0
		} else {
			cal, err := strconv.Atoi(l)
			if err != nil {
				log.Fatal(err)
			}
			current += cal
		}
	}

	// Also compare last elf just to be safe
	// Not needed when input ends with blank line
	if current > highest {
		highest = current
	}

	return highest
}

func part2(lines []string) int {
	total := 0
	var caloriesPerElf []int
	current := 0

	for _, l := range lines {
		// If empty line, we are done with this elf, check if it's the highest
		// value
		if len(l) == 0 {
			caloriesPerElf = append(caloriesPerElf, current)
			current = 0
		} else {
			cal, err := strconv.Atoi(l)
			if err != nil {
				log.Fatal(err)
			}
			current += cal
		}
	}

	// Also add last elf if it's not already done
	// Not needed when input ends with blank line
	if current > 0 {
		caloriesPerElf = append(caloriesPerElf, current)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(caloriesPerElf)))

	for i := 0; i < 3; i++ {
		total += caloriesPerElf[i]
	}

	return total
}
