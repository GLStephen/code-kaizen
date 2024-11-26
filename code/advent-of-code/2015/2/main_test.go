package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {

	strings, err := loadInput()
	if err != nil {
		assert.Fail(t, "Error loading input", err)
	}
	result, ribbonFeet, err := solve(strings)

	assert.Nil(t, err)
	assert.Equal(t, 1606483, result)
	assert.Equal(t, 0, ribbonFeet)
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
