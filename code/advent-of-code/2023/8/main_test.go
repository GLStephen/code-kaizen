package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestKnownSolve(t *testing.T) {
	input, err := loadInput("known_solve_1.txt")
	if err != nil {
		t.Error(err)
	}
	expected := 2
	result := solve(input)
	if result != expected {
		t.Errorf("%s = %d; want %d", "known_solve_1.txt", result, expected)
	}
}

func TestKnownSolve2(t *testing.T) {
	input, err := loadInput("known_solve_2.txt")
	if err != nil {
		t.Error(err)
	}
	expected := 6
	result := solve(input)
	if result != expected {
		t.Errorf("%s = %d; want %d", "known_solve_2.txt", result, expected)
	}
}

// func TestUnknownSolve(t *testing.T) {
// 	input, err := loadInput("input.txt")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	expected := 0
// 	result := solve(input)
// 	if result != expected {
// 		t.Errorf("%s = %d; want %d", "input.txt", result, expected)
// 	}
// }

func loadInput(loadLocation string) (inputValues, error) {
	file, err := os.Open(loadLocation)
	if err != nil {
		return inputValues{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	directions := scanner.Text()

	var directionsArray []string
	for _, byte := range directions {
		directionsArray = append(directionsArray, string(byte))
	}

	var nodes = make(map[string]inputNode)
	var start = "AAA" // it's always AAA
	for scanner.Scan() {
		// skip blanks
		line := scanner.Text()
		if line != "" {
			//nodes = append(nodes, strings.TrimSpace(scanner.Text()))
			linePieces := strings.Split(line, "=")
			lineNodes := strings.Trim(linePieces[1], " ")
			lineNodes = strings.Trim(lineNodes, "(")
			lineNodes = strings.Trim(lineNodes, ")")
			nodeElements := strings.Split(lineNodes, ", ")

			inputNode := inputNode{left: nodeElements[0], right: nodeElements[1]}
			nodes[strings.Trim(linePieces[0], " ")] = inputNode
			// START IS ALWAYS AAA
			// if start == "" {
			// 	start = strings.Trim(linePieces[0], " ")
			// }
			//nodes = append(nodes, inputNode)
		}
	}

	// verify their shape
	// for _, direction := range directions {
	// 	fmt.Printf("-%+v-\n", direction)
	// }

	// for _, node := range nodes {
	// 	fmt.Printf("-%+v-\n", node)
	// }

	loadedValues := inputValues{start: start, directions: directionsArray, nodes: nodes}

	// fmt.Printf("Loaded \n%+v\n", loadedValues)

	return loadedValues, nil
}

// func TestKnownParallel(t *testing.T) {
// 	input, err := loadParellelInput("known_parallel.txt")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	expected := 6
// 	result := solveParallelAlternative(input)
// 	if result != expected {
// 		t.Errorf("%s = %d; want %d", "input.txt", result, expected)
// 	}
// }

func TestUnknownParallel(t *testing.T) {
	input, err := loadParellelInput("input.txt")
	if err != nil {
		t.Error(err)
	}
	expected := 0
	result := solveParallelAlternative(input)
	if result != expected {
		t.Errorf("%s = %d; want %d", "input.txt", result, expected)
	}
}

func loadParellelInput(loadLocation string) (inputValuesParallel, error) {
	file, err := os.Open(loadLocation)
	if err != nil {
		return inputValuesParallel{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	directions := scanner.Text()

	var directionsArray []string
	for _, byte := range directions {
		directionsArray = append(directionsArray, string(byte))
	}

	var nodes = make(map[string]inputNode)
	var start []string // not always A, we're going to have to find them!
	for scanner.Scan() {
		// skip blanks
		line := scanner.Text()
		if line != "" {
			//nodes = append(nodes, strings.TrimSpace(scanner.Text()))
			linePieces := strings.Split(line, "=")
			lineNodes := strings.Trim(linePieces[1], " ")
			lineNodes = strings.Trim(lineNodes, "(")
			lineNodes = strings.Trim(lineNodes, ")")
			nodeElements := strings.Split(lineNodes, ", ")

			inputNode := inputNode{left: nodeElements[0], right: nodeElements[1]}
			nodes[strings.Trim(linePieces[0], " ")] = inputNode

			trimmedItemKey := strings.Trim(linePieces[0], " ")
			if string(trimmedItemKey[len(trimmedItemKey)-1:][0]) == "A" {
				start = append(start, trimmedItemKey)
			}
			// START IS ALWAYS AAA
			// if start == "" {
			// 	start = strings.Trim(linePieces[0], " ")
			// }
			//nodes = append(nodes, inputNode)
		}
	}

	// verify their shape
	// for _, direction := range directions {
	// 	fmt.Printf("-%+v-\n", direction)
	// }

	// for _, node := range nodes {
	// 	fmt.Printf("-%+v-\n", node)
	// }

	loadedValues := inputValuesParallel{start: start, directions: directionsArray, nodes: nodes}

	// fmt.Printf("Loaded \n%+v\n", loadedValues)

	return loadedValues, nil
}
