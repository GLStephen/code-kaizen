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
	solved, err := solve(input)

	assert.Nil(t, err)
	assert.Equal(t, 232, solved)
}

func TestSolveBonus(t *testing.T) {
	input, err := loadInput()
	if err != nil {
		assert.Fail(t, "Error loading input", err)
	}
	solved, currentFloor, err := solveBonus(input)

	assert.Nil(t, err)
	assert.Equal(t, -1, currentFloor)
	assert.Equal(t, 12, solved)
}

func loadInput() ([]byte, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return make([]byte, 0), err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return []byte(scanner.Text()), nil

}
