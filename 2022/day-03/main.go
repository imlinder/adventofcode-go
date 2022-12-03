package main

import (
	"fmt"
	"github.com/imlinder/adventofcode-go/util"
	"strings"
	"unicode"
)

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Printf("Part 1: %v \n", part1(lines))
	fmt.Printf("Part 2: %v \n", part2(lines))
}

func part1(lines []string) int {
	total := 0
	for _, l := range lines {
		firstHalf := l[:len(l)/2]
		secondHalf := l[len(l)/2:]
		firstSlice := strings.Split(firstHalf, "")
		secondSlice := strings.Split(secondHalf, "")

		item := firstIntersection(firstSlice, secondSlice)
		if len(item) > 0 {
			prio := itemPriority([]rune(item)[0])
			total += prio
		}
	}

	return total
}

func part2(lines []string) int {
	total := 0
	for i := 0; i < len(lines)/3; i++ {
		group := lines[i*3 : i*3+3]
		badge := findBadge(group)
		prio := itemPriority([]rune(badge)[0])
		total += prio
	}
	return total
}

func firstIntersection(s1 []string, s2 []string) string {
	// O(n^2)
	for _, r := range s1 {
		for _, r2 := range s2 {
			if r == r2 {
				return r
			}
		}
	}
	return ""
}

func itemPriority(letter rune) int {
	if !unicode.IsLetter(letter) {
		return 0
	}
	num := int(letter)
	num -= 38
	if unicode.IsLower(letter) {
		num -= 58
	}
	return num
}

func findBadge(lines []string) string {
	for _, c := range lines[0] {
		inAll := true
		for _, l := range lines[1:] {
			if !strings.Contains(l, string(c)) {
				inAll = false
				break
			}
		}

		if inAll {
			return string(c)
		}
	}

	return ""
}
