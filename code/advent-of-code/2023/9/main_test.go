package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/glstephen/aoc-2023/2023/pkg"
)

func TestKnown1(t *testing.T) {
	input := loadInput("known1.txt")

	result := solve(input)

	expected := 114

	if result != expected {
		t.Errorf("known1.txt = %d; want %d", result, expected)
	}
}

/*
This function
*/
func loadInput(location string) problemInput {
	file, err := pkg.OpenFile(location)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	problemInput := problemInput{}

	for scanner.Scan() {
		line := scanner.Text()
		entries := strings.Split(line, " ")
		intEntries := make([]int, len(entries))
		for i := range entries {
			intVal, err := strconv.Atoi(strings.TrimSpace(entries[i]))
			if err != nil {
				panic(err)
			}
			intEntries[i] = intVal
		}
		problemInput.histories = append(problemInput.histories, measurementHistory{measurements: intEntries})
	}

	fmt.Printf("problemInput: %+v\n", problemInput)

	return problemInput
}
