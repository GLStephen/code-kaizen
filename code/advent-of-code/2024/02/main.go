package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	knownData, err := extractData("./input-known.txt")
	if err != nil {
		panic(err)
	}

	showData(knownData)

	runSolution(knownData)

	data, err := extractData("./input1.txt")
	if err != nil {
		panic(err)
	}
	runSolution(data)

	runSolutionBonus(knownData)

	runSolutionBonus(data)

	// bonusData, err := extractData("./input2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// runSolution(bonusData)
}

func runSolution(data [][]int) {
	safeCount := 0

	for _, report := range data {
		if isReportSafe(report) {
			safeCount++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)
}

func isReportSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	// Check first difference to determine if we're looking for increasing or decreasing
	isIncreasing := levels[1] > levels[0]

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		// Check if difference is between 1 and 3
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}

		// Check if direction matches initial direction
		if isIncreasing && diff <= 0 {
			return false
		}
		if !isIncreasing && diff >= 0 {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func showData(data [][]int) {
	println("Data:")
	for i, row := range data {
		println(fmt.Sprintf("Row %d: %v", i, row))
	}
}

func extractData(filepath string) ([][]int, error) {
	data := make([][]int, 0)

	file, err := os.Open(filepath)
	if err != nil {
		return data, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into a list of numbers
		numbers := strings.Split(line, " ")
		// Convert the numbers to integers
		numbersInt := make([]int, len(numbers))
		for i, number := range numbers {
			numbersInt[i], _ = strconv.Atoi(number)
		}
		// Add the list of numbers to the data
		data = append(data, numbersInt)
	}

	return data, nil
}

func runSolutionBonus(data [][]int) {
	safeCount := 0

	for _, report := range data {
		if isReportSafeWithDampener(report) {
			safeCount++
		}
	}

	fmt.Printf("Number of safe reports with Problem Dampener: %d\n", safeCount)
}

func isReportSafeWithDampener(levels []int) bool {
	// If it's already safe, no need to try removing levels
	if isReportSafe(levels) {
		return true
	}

	// Try removing each level one at a time
	for i := 0; i < len(levels); i++ {
		// Create new slice without current level
		dampened := make([]int, 0, len(levels)-1)
		dampened = append(dampened, levels[:i]...)
		dampened = append(dampened, levels[i+1:]...)

		// Check if removing this level makes the report safe
		if isReportSafe(dampened) {
			return true
		}
	}

	return false
}
