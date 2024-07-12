package main

import (
	"fmt"
	"github.com/imlinder/adventofcode-go/util"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Printf("Part 1: %v \n", part1(lines))
	fmt.Printf("Part 2: %v \n", part2(lines))
}

func part1(lines []string) int {
	sum := 0
	re := regexp.MustCompile(`\d`)
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		matches := re.FindAllString(l, -1)
		num, _ := strconv.Atoi(matches[0] + matches[len(matches)-1])
		sum += num
	}

	return sum
}

func part2(lines []string) int {
	sum := 0

	replacer := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
	revReplacer := strings.NewReplacer("eno", "1", "owt", "2", "eerht", "3", "ruof", "4", "evif", "5", "xis", "6", "neves", "7", "thgie", "8", "enin", "9")

	re := regexp.MustCompile(`\d`)

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		lr := Reverse(l)

		lforward := replacer.Replace(l)
		lreverse := revReplacer.Replace(lr)

		first := re.FindString(lforward)
		last := re.FindString(lreverse)

		num, _ := strconv.Atoi(first + last)
		sum += num
	}

	return sum
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
