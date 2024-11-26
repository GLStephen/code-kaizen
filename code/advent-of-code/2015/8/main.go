package main

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello from 2015 Day 8")
}

func solve(input []string) (int, int, error) {
	totalCount := 0
	totalReencoded := 0

	for _, line := range input {
		// clean up the quoted strings
		cleanLine := line
		cleanLine = strings.TrimPrefix(cleanLine, `"`)
		cleanLine = strings.Trim(cleanLine, `"`)
		cleanLine = strings.ReplaceAll(cleanLine, `\\`, `\`)
		cleanLine = strings.ReplaceAll(cleanLine, `\"`, `"`)

		pattern := regexp.MustCompile(`\\x([0-9a-fA-F]{2})`)
		match := pattern.FindAllStringSubmatch(cleanLine, -1)

		if match != nil {
			for _, m := range match {
				fmt.Printf("Found match ->%s<- ->%s<- ->%s<-\n", m[0], m[1], hexToText(m[1]))
				cleanLine = strings.ReplaceAll(cleanLine, m[0], "-")
			}
		}

		fmt.Printf("String ->%s<- to ->%s<-\n", line, cleanLine)

		totalReencoded += (strings.Count(line, `\`) + strings.Count(line, `"`) + len(line) + 2) - len(line)

		totalCount += len(line) - len(cleanLine)
	}

	return totalCount, totalReencoded, nil
}

func hexToText(hexStr string) string {

	// Decode hex string to bytes
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		panic(err)
	}

	// Convert bytes to rune slice
	var runes []rune
	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		runes = append(runes, r)
		bytes = bytes[size:]
	}

	// Convert rune slice to string
	return string(runes)
}
