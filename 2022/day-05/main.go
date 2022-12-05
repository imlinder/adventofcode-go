package main

import (
	"fmt"
	"github.com/imlinder/adventofcode-go/util"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	lines := util.ReadFile("input.txt")
	part1 := part1(lines)
	fmt.Printf("Part 1: %v \n", part1)
	part2 := part2(lines)
	fmt.Printf("Part 2: %v \n", part2)
}

func part1(lines []string) string {
	stacks := getStacks(&lines)
	stacks = rearrangeStacks(stacks, &lines, false)
	result := ""
	for _, s := range stacks {
		result = result + string(s[0])
	}
	return result
}

func part2(lines []string) string {
	stacks := getStacks(&lines)
	stacks = rearrangeStacks(stacks, &lines, true)
	result := ""
	for _, s := range stacks {
		result = result + string(s[0])
	}
	return result
}

// Get the total number of stacks that exists for the input.
func getStackCount(lines *[]string) int {
	for _, l := range *lines {
		if len(l) < 1 {
			continue
		}
		trimmed := strings.TrimSpace(l)
		if trimmed[0:1] == "1" {
			i, err := strconv.Atoi(trimmed[len(trimmed)-1:])
			if err != nil {
				log.Fatal(err)
			}
			return i
		}
	}
	return 0
}

// Create a slice of strings to represent the stacks
func getStacks(lines *[]string) []string {
	stackCount := getStackCount(lines)
	stacks := make([]string, stackCount)
	for _, l := range *lines {
		if len(l) < 1 {
			continue
		}
		if !strings.Contains(l, "[") {
			break
		}
		for i, c := range l {
			if unicode.IsLetter(c) {
				stackIndex := (i - 1) / 4
				stacks[stackIndex] = stacks[stackIndex] + string(c)
			}
		}
	}

	return stacks
}

func rearrangeStacks(stacks []string, lines *[]string, moveMultiple bool) []string {
	for _, l := range *lines {
		if len(l) < 1 || l[:4] != "move" {
			continue
		}

		instr := strings.Split(l, " ")
		num, from, to := strToInt(instr[1]), strToInt(instr[3])-1, strToInt(instr[5])-1
		// Remove from stack
		if !moveMultiple {
			stacks[to] = reverseString(stacks[from][:num]) + stacks[to]
		} else {
			stacks[to] = stacks[from][:num] + stacks[to]
		}
		stacks[from] = stacks[from][num:]
	}

	return stacks
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func reverseString(str string) string {
	result := ""
	for _, v := range str {
		result = string(v) + result
	}
	return result
}
