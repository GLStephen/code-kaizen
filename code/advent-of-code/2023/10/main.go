package main

import (
	"fmt"
	"math"
)

type problemInput struct {
	start       coordinate
	coordinates map[int]map[int]connections
}

type connections struct {
	marker string
	N      *coordinate
	S      *coordinate
	E      *coordinate
	W      *coordinate
}

type coordinate struct {
	x int
	y int
}

func main() {
	input := problemInput{}
	solve(input)
}

func solve(input problemInput) float64 {
	count := 1

	firstY := input.start.y
	firstX := input.start.x

	firstStep := input.coordinates[input.start.y][input.start.x]
	from := ""
	// we can head in any direction
	if firstStep.N != nil {
		from = "S"
		firstY = input.start.y - 1
	} else if firstStep.S != nil {
		from = "N"
		firstY = input.start.y + 1
	} else if firstStep.W != nil {
		from = "E"
		firstX = input.start.x - 1
	} else if firstStep.E != nil {
		from = "W"
		firstX = input.start.x + 1
	}

	//fmt.Printf("On %d,%d Moving from %s on step %d \n", firstY, firstX, from, &count)
	findStart(&count, from, firstY, firstX, input)

	// handle came from and going to...

	return math.Floor(float64(count) / 2)
}

func solveContained(input problemInput) float64 {
	count := 1

	firstY := input.start.y
	firstX := input.start.x

	firstStep := input.coordinates[input.start.y][input.start.x]
	from := ""
	// we can head in any direction
	if firstStep.N != nil {
		from = "S"
		firstY = input.start.y - 1
	} else if firstStep.S != nil {
		from = "N"
		firstY = input.start.y + 1
	} else if firstStep.W != nil {
		from = "E"
		firstX = input.start.x - 1
	} else if firstStep.E != nil {
		from = "W"
		firstX = input.start.x + 1
	}

	//fmt.Printf("On %d,%d Moving from %s on step %d \n", firstY, firstX, from, &count)
	findStart(&count, from, firstY, firstX, input)

	// handle came from and going to...

	return math.Floor(float64(count) / 2)
}

func findStart(count *int, from string, currentY int, currentX int, input problemInput) {
	stop := input.coordinates[currentY][currentX]
	fmt.Printf("+On %d,%d Moving from %s on step %d to %#v \n", currentY, currentX, from, *count, stop)
	nextFrom := ""
	if stop.marker == "S" {
		return
	} else {
		*count++
		if from != "N" && stop.N != nil {
			nextFrom = "S"
			currentY = currentY - 1
			fmt.Printf("Moving N\n")
		} else if from != "S" && stop.S != nil {
			nextFrom = "N"
			currentY = currentY + 1
			fmt.Printf("Moving S\n")
		} else if from != "W" && stop.W != nil {
			nextFrom = "E"
			currentX = currentX - 1
			fmt.Printf("Moving W\n")
		} else if from != "E" && stop.E != nil {
			nextFrom = "W"
			currentX = currentX + 1
			fmt.Printf("Moving E\n")
		}

		fmt.Printf("-Moving to %d,%d from %s on step %d \n", currentY, currentX, nextFrom, *count)
		findStart(count, nextFrom, currentY, currentX, input)
	}
}
