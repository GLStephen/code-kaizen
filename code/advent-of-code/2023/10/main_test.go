package main

import (
	"bufio"
	"fmt"
	"testing"

	"github.com/glstephen/aoc-2023/2023/pkg"
)

// func TestKnownSimpleSolve(t *testing.T) {
// 	input := loadInput("known1.txt")

// 	result := solve(input)

// 	expected := float64(4)

// 	if result != expected {
// 		t.Errorf("known1.txt = %f; want %f", result, expected)
// 	}

// }

// func TestKnownSimpleSolv2(t *testing.T) {
// 	input := loadInput("known2.txt")

// 	result := solve(input)

// 	expected := float64(8)

// 	if result != expected {
// 		t.Errorf("known2.txt = %f; want %f", result, expected)
// 	}

// }

// func TestUnknownSolve(t *testing.T) {
// 	input := loadInput("input.txt")

// 	result := solve(input)

// 	expected := float64(0)

// 	if result != expected {
// 		t.Errorf("known2.txt = %f; want %f", result, expected)
// 	}

// }

func TestPart2Known(t *testing.T) {
	input := loadInput("step2known1.txt")

	result := solveContained(input)

	expected := float64(8)

	if result != expected {
		t.Errorf("known2.txt = %f; want %f", result, expected)
	}

}

func loadInput(location string) problemInput {
	file, err := pkg.OpenFile(location)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	problemInput := problemInput{
		coordinates: make(map[int]map[int]connections),
	}

	yCoord := 0
	xCoord := 0
	for scanner.Scan() {
		xCoord = 0
		line := scanner.Text()
		problemInput.coordinates[yCoord] = make(map[int]connections)
		for _, entry := range line {
			problemInput.coordinates[yCoord][xCoord] = getLocationConnections(yCoord, xCoord, entry)
			fmt.Printf("%+s", string(entry))
			xCoord++
		}
		fmt.Println("")
		yCoord++
	}

	totalLines := yCoord
	totalColumns := xCoord

	startFoundX := -1
	startFoundY := -1

	// find the starting coordinates
	foundSouth := false
	for yCoord, coordinate := range problemInput.coordinates {
		for xCoord, connections := range coordinate {
			if connections.marker == "S" {
				foundSouth = true
				//check north
				if yCoord-1 >= 0 && problemInput.coordinates[yCoord-1][xCoord].S != nil {
					fmt.Printf("Connecting Start to North\n")
					connections.N = getNorth(xCoord, yCoord)
				}
				if yCoord+1 <= totalLines && problemInput.coordinates[yCoord+1][xCoord].N != nil {
					fmt.Printf("Connecting Start to South\n")
					connections.S = getSouth(xCoord, yCoord)
				}
				if xCoord-1 >= 0 && problemInput.coordinates[yCoord][xCoord-1].E != nil {
					fmt.Printf("Connecting Start to West\n")
					connections.W = getWest(xCoord, yCoord)
				}
				if xCoord+1 <= totalColumns && problemInput.coordinates[yCoord][xCoord+1].W != nil {
					fmt.Printf("Connecting Start to East\n")
					connections.E = getEast(xCoord, yCoord)
				}
				startFoundY = yCoord
				startFoundX = xCoord
				problemInput.coordinates[yCoord][xCoord] = connections
				fmt.Printf("marker: %s row: %d column: %d\n", connections.marker, yCoord, xCoord)
			}
			if foundSouth {
				break
			}
		}
		if foundSouth {
			break
		}
	}

	problemInput.start = coordinate{x: startFoundX, y: startFoundY}

	// prune all connections out of bounds - where X or Y is smaller than 0 or greater then the number of lines
	// if xCoord < 0 or yCoord < 0 or xCoord > totalColumns or yCoord > totalLines then prune

	fmt.Printf("totalLines: %d totalColumns: %d problemInput: \n", yCoord, xCoord)
	//spew.Dump(problemInput)

	return problemInput
}

func getLocationConnections(yCoord int, xCoord int, entry rune) connections {
	conn := connections{}
	entryValue := string(entry)
	conn.marker = entryValue
	switch entryValue {
	case "|":
		// NS
		conn.N = getNorth(xCoord, yCoord)
		conn.S = getSouth(xCoord, yCoord)
	case "-":
		// EW
		conn.E = getEast(xCoord, yCoord)
		conn.W = getWest(xCoord, yCoord)
	case "L":
		// NE
		conn.N = getNorth(xCoord, yCoord)
		conn.E = getEast(xCoord, yCoord)
	case "J":
		// NW
		conn.N = getNorth(xCoord, yCoord)
		conn.W = getWest(xCoord, yCoord)
	case "7":
		// SW
		conn.S = getSouth(xCoord, yCoord)
		conn.W = getWest(xCoord, yCoord)
	case "F":
		// SE
		conn.S = getSouth(xCoord, yCoord)
		conn.E = getEast(xCoord, yCoord)
	case "S":
		// SE
		// we will figure this out later
	case ".":
		// no connections
	default:
		panic("invalid entry")
	}

	return conn
}

func getNorth(x int, y int) *coordinate {
	return &coordinate{x, y - 1}
}

func getSouth(x int, y int) *coordinate {
	return &coordinate{x, y + 1}
}

func getEast(x int, y int) *coordinate {
	return &coordinate{x + 1, y}
}

func getWest(x int, y int) *coordinate {
	return &coordinate{x - 1, y}
}
