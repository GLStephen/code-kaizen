package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	fmt.Println("Hello from 2015 Day 4")
}

func solve(input string) (int, error) {
	found := false
	iteration := 0
	for !found {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input, iteration))))
		firstSegment := hash[0:6]
		if firstSegment == "000000" {
			found = true
		} else {
			iteration++
		}
	}
	return iteration, nil
}
