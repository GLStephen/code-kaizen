package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	input, err := loadInput()
	if err != nil {
		assert.Fail(t, "Error loading input", err)
	}

	count, err := solve(input)

	assert.Nil(t, err)
	assert.Equal(t, 2592, count)
}

func TestSolveBonus(t *testing.T) {
	input, err := loadInput()
	if err != nil {
		assert.Fail(t, "Error loading input", err)
	}

	count, err := solveBonus(input)

	assert.Nil(t, err)
	assert.Equal(t, 0, count)
}

func loadInput() ([]string, error) {
	results := make([]string, 0)
	file, err := os.Open("./input.txt")
	if err != nil {
		return results, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		results = append(results, line)
		//fmt.Printf("%d: %s\n", count, line)
		count++
	}

	return results, nil
}
