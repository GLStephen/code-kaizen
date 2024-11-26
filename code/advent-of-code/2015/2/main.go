package main

import (
	"fmt"
	"strconv"
	"strings"
)

type presentSet struct {
	presents []present
}

type present struct {
	x                  int
	y                  int
	z                  int
	side1Area          int
	side2Area          int
	side3Area          int
	smallest           int
	sizeNeeded         int
	ribbonFeetRequired int
	bowFeetRequired    int
	smallestFaceLength int
}

func main() {
	fmt.Println("Day 2, 2015!")
}

func solve(inputs []string) (int, int, error) {

	presentSet := presentSet{}
	fullPaperNeeded := 0
	fullRibbonFeetNeeded := 0

	for _, input := range inputs {
		present := present{}
		fmt.Printf("input: -->%s<--\n", input)
		parts := strings.Split(input, "x")
		fmt.Printf("input: %+v\n", parts)
		present.x, _ = strconv.Atoi(parts[0])
		present.y, _ = strconv.Atoi(parts[1])
		present.z, _ = strconv.Atoi(parts[2])

		side1 := present.x * present.z
		side1FaceLength := (present.x + present.z) * 2
		present.side1Area = side1 * 2
		present.smallest = side1
		present.smallestFaceLength = side1FaceLength

		side2 := present.y * present.z
		side2FaceLength := (present.y + present.z) * 2
		present.side2Area = side2 * 2
		if side2 < present.smallest {
			present.smallest = side2
		}
		if side2FaceLength < present.smallestFaceLength {
			present.smallestFaceLength = side2FaceLength
		}

		side3 := present.x * present.y
		side3FaceLength := (present.x + present.y) * 2
		present.side3Area = side3 * 2
		if side3 < present.smallest {
			present.smallest = side3
		}
		if side3FaceLength < present.smallestFaceLength {
			present.smallestFaceLength = side3FaceLength
		}

		present.sizeNeeded = present.smallest + present.side1Area + present.side2Area + present.side3Area
		fullPaperNeeded = fullPaperNeeded + present.sizeNeeded

		present.ribbonFeetRequired = present.smallestFaceLength
		present.bowFeetRequired = present.x * present.y * present.z
		fullRibbonFeetNeeded = fullRibbonFeetNeeded + present.ribbonFeetRequired + present.bowFeetRequired

		presentSet.presents = append(presentSet.presents, present)
	}

	fmt.Printf("Presents: %+v\n", presentSet)

	return fullPaperNeeded, fullRibbonFeetNeeded, nil
}
