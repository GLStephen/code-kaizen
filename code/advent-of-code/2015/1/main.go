package main

import "fmt"

func main() {
	fmt.Println("Hello from 2015 Day 1!")
}

func solve(input []byte) (int, error) {
	startingIndex := 0
	for _, entry := range input {
		//fmt.Printf("-->%s<--", string(entry))
		switch string(entry) {
		case "(":
			startingIndex++
		case ")":
			startingIndex--
		}
	}
	return startingIndex, nil
}

func solveBonus(input []byte) (int, int, error) {
	currentFloor := 0
	for index, entry := range input {
		//fmt.Printf("-->%s<--", string(entry))
		switch string(entry) {
		case "(":
			currentFloor++
		case ")":
			currentFloor--
		}
		if currentFloor == -1 {
			return index + 1, currentFloor, nil
		}
	}
	return currentFloor, 0, fmt.Errorf("no floor found")
}
