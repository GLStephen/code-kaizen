package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	println("Running Solution for Day 1")

	column1, column2, err := extractColumns("input1.txt")

	if err != nil {
		panic(err)
	}

	slices.Sort(column1)
	slices.Sort(column2)

	testColumn1, testColumn2, err := extractColumns("input-test.txt")

	if err != nil {
		panic(err)
	}

	runSolution1(column1, column2) // 2113135 is correct

	runSolution2("Test", testColumn1, testColumn2) // 31 is correct

	runSolution2("Full", column1, column2) // 31 is correct

}

func runSolution2(label string, column1 []int, column2 []int) {
	println("Column 1:")
	score := calculateSimilarityScore(column1, column2)
	fmt.Printf("%s Score: %d\n", label, score)
}

func calculateSimilarityScore(leftList, rightList []int) int {
	totalScore := 0

	// For each number in the left list
	for _, leftNum := range leftList {
		// Count occurrences in right list
		occurrences := 0
		for _, rightNum := range rightList {
			if leftNum == rightNum {
				occurrences++
			}
		}
		// Add to total score (number * occurrences)
		totalScore += leftNum * occurrences
	}

	return totalScore
}

// 2113135 is correct
func runSolution1(column1 []int, column2 []int) {
	println("Column 1:")
	for i, value := range column1 {
		fmt.Printf("Index %d: %v\n", i, value)
	}

	println("Column 2:")
	for i, value := range column2 {
		fmt.Printf("Index %d: %v\n", i, value)
	}

	fmt.Printf("Lengths Column 1: %d Column 2: %d\n", len(column1), len(column2))

	sumEntryDiffs(column1, column2)
}

func sumEntryDiffs(column1 []int, column2 []int) {
	sum := 0
	for i := 0; i < len(column1); i++ {
		sum += int(math.Abs(float64(column2[i] - column1[i])))
		if (column2[i] - column1[i]) < 0 {
			fmt.Printf("Entry %d: %d\n", i, column1[i]-column2[i])
		}
	}
	fmt.Printf("Sum of Entry Diffs: %d\n", sum)
}

func extractColumns(filepath string) ([]int, []int, error) {
	column1 := make([]int, 0)
	column2 := make([]int, 0)

	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 2 {
			val1, err := strconv.Atoi(fields[0])
			if err != nil {
				return nil, nil, err
			}
			val2, err := strconv.Atoi(fields[1])
			if err != nil {
				return nil, nil, err
			}
			column1 = append(column1, val1)
			column2 = append(column2, val2)
		}
	}

	return column1, column2, nil
}
