package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello from 2015 Day 5")
}

func solve(input []string) (int, int, error) {
	nice := 0
	naughty := 0

	for _, line := range input {
		if testThreeVowels(line) && testDoubleLetter(line) && testNoBadStrings(line) {
			nice++
		} else {
			naughty++
		}
	}

	return nice, naughty, nil
}

func testDoubleLetter(line string) bool {
	// make an array of all letters of the alphabet doubled
	doubled := []string{"aa", "bb", "cc", "dd", "ee", "ff", "gg", "hh", "ii", "jj", "kk", "ll", "mm", "nn", "oo", "pp", "qq", "rr", "ss", "tt", "uu", "vv", "ww", "xx", "yy", "zz"}

	for _, doublet := range doubled {
		if index := strings.Index(line, doublet); index >= 0 {
			return true
		}
	}

	return false
}

func testNoBadStrings(line string) bool {
	// make an array of all letters of the alphabet doubled
	badList := []string{"ab", "cd", "pq", "xy"}

	for _, bad := range badList {
		if index := strings.Index(line, bad); index >= 0 {
			return false
		}
	}

	return true
}

func testThreeVowels(test string) bool {
	chars := []rune{'a', 'e', 'i', 'o', 'u'}

	if countChars(test, chars, 3) == 3 {
		return true
	} else {
		return false
	}
}

func countChars(s string, chars []rune, limit int) int {
	counts := 0

	for _, c := range s {
		if containsRune(chars, c) {
			counts++
			if counts >= limit {
				return counts
			}
		}
	}

	return counts
}

func containsRune(s []rune, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

func solveBonus(input []string) (int, int, error) {
	nice := 0
	naughty := 0

	for _, line := range input {
		if hasNonOverlappingPair(line) && hasRepeatWithOneBetween(line) {
			nice++
		} else {
			naughty++
		}
	}

	return nice, naughty, nil
}

func hasNonOverlappingPair(line string) bool {
	paddedString := fmt.Sprintf("~%s+", line)
	fmt.Printf("testing overlap %s\n", line)
	for i := 0; i < len(paddedString)-1; i++ {
		pair := paddedString[i : i+2]
		fmt.Printf("%d %s\n", strings.Count(paddedString, pair), pair)
		if strings.Count(paddedString, pair) >= 2 {
			return true
		}
	}
	return false
}

func hasRepeatWithOneBetween(line string) bool {
	paddedString := fmt.Sprintf("~%s+", line)
	fmt.Printf("testing repeat %s\n", line)
	for i := 0; i < len(paddedString)-2; i++ {
		fmt.Printf("%s %s\n", string(paddedString[i]), string(paddedString[i+2]))
		if paddedString[i] == paddedString[i+2] {
			return true
		}
	}

	return false
}
