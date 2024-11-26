package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstructions(t *testing.T) {
	instruction := getInstruction("turn on 0,0 through 999,999")
	assert.Equal(t, 1, instruction.command)
	assert.Equal(t, 0, instruction.startX)
	assert.Equal(t, 0, instruction.startY)
	assert.Equal(t, 999, instruction.endX)
	assert.Equal(t, 999, instruction.endY)

	instruction = getInstruction("toggle 6,3 through 997,0")
	assert.Equal(t, -1, instruction.command)
	assert.Equal(t, 6, instruction.startX)
	assert.Equal(t, 3, instruction.startY)
	assert.Equal(t, 997, instruction.endX)
	assert.Equal(t, 0, instruction.endY)

	instruction = getInstruction("turn off 0,0 through 999,0")
	assert.Equal(t, 0, instruction.command)
	assert.Equal(t, 0, instruction.startX)
	assert.Equal(t, 0, instruction.startY)
	assert.Equal(t, 999, instruction.endX)
	assert.Equal(t, 0, instruction.endY)
}

func TestGiven(t *testing.T) {
	count, _ := solve([]string{"turn on 0,0 through 999,999"})
	assert.Equal(t, 1000000, count)

	count, _ = solve([]string{"toggle 0,0 through 999,0"})
	assert.Equal(t, 1000, count)

	count, _ = solve([]string{"turn on 0,0 through 999,999", "toggle 0,0 through 999,0"})
	assert.Equal(t, 999000, count)

	count, _ = solve([]string{"turn on 0,0 through 0,9", "toggle 0,0 through 0,4", "turn off 0,0 through 0,1"})
	assert.Equal(t, 5, count)

	count, _ = solve([]string{"turn on 0,0 through 1,9", "toggle 0,0 through 0,4", "turn off 0,0 through 0,1"})
	assert.Equal(t, 15, count)

	count, _ = solve([]string{"turn on 0,0 through 1,9", "turn off 0,0 through 1,9"})
	assert.Equal(t, 0, count)
}

func TestSolve(t *testing.T) {
	input, _ := loadInput()
	count, err := solve(input)

	assert.Nil(t, err)
	assert.Equal(t, 400410, count)
}

func TestBonusGiven(t *testing.T) {
	count, _ := solveBonus([]string{"turn on 0,0 through 0,0"})
	assert.Equal(t, 1, count)

	count, _ = solveBonus([]string{"toggle 0,0 through 999,999"})
	assert.Equal(t, 2000000, count)
}

func TestSolveBonus(t *testing.T) {
	input, _ := loadInput()
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
