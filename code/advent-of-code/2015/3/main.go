package main

import "fmt"

func main() {
	fmt.Println("Hello from 2015 Day 3")
}

func solve(input []string) (int, error) {
	// get the first entry of the input string array
	// then loop through each character of the string
	// start at coordinate 0,0
	// if the character is < then move left
	// if the character is > then move right
	// if the character is ^ then move up
	// if the character is ^ then move up
	// store the current coordinate in a map
	// once complete count the entries in the map and return the count

	// get the first entry of the input string array
	directions := input[0]

	// start at coordinate 0,0
	x := 0
	y := 0

	// store the current coordinate in a map
	coordinates := make(map[string]int)

	// loop through each character of the string
	for index, dir := range directions {
		switch dir {
		case '<':
			x--
		case '>':
			x++
		case '^':
			y++
		case 'v':
			y--
		}

		// store coordinate
		fmt.Printf("%d: %s\n", index, string(dir))
		if _, ok := coordinates[fmt.Sprintf("%d,%d", x, y)]; !ok {
			coordinates[fmt.Sprintf("%d,%d", x, y)] = 1
			fmt.Printf("%d,%d NOT %d\n", x, y, len(coordinates))
		} else {
			fmt.Printf("%d,%d FOUND\n", x, y)
		}
	}

	// once complete count the entries in the map and return the count
	return len(coordinates) + 1, nil
}

func solveBonus(input []string) (int, error) {
	// get the first entry of the input string array
	// then loop through each character of the string
	// start at coordinate 0,0
	// if the character is < then move left
	// if the character is > then move right
	// if the character is ^ then move up
	// if the character is ^ then move up
	// store the current coordinate in a map
	// once complete count the entries in the map and return the count

	// get the first entry of the input string array
	directions := input[0]

	// start at coordinate 0,0
	xSanta := 0
	ySanta := 0

	// start at coordinate 0,0
	xRobot := 0
	yRobot := 0

	// store the current coordinate in a map
	coordinates := make(map[string]int)

	coordinates[fmt.Sprintf("%d,%d", 0, 0)] = 1

	// loop through each character of the string
	for index, dir := range directions {
		if index > 5 {
			// break
		}
		if (index % 2) == 0 {
			switch dir {
			case '<':
				xSanta--
			case '>':
				xSanta++
			case '^':
				ySanta++
			case 'v':
				ySanta--
			}

			// store coordinate
			fmt.Printf("%d: %s\n", index, string(dir))
			if _, ok := coordinates[fmt.Sprintf("%d,%d", xSanta, ySanta)]; !ok {
				coordinates[fmt.Sprintf("%d,%d", xSanta, ySanta)] = 1
				fmt.Printf("%d,%d Santa NOT %d\n", xSanta, ySanta, len(coordinates))
			} else {
				fmt.Printf("%d,%d Santa FOUND\n", xSanta, ySanta)
			}
		} else {
			switch dir {
			case '<':
				xRobot--
			case '>':
				xRobot++
			case '^':
				yRobot++
			case 'v':
				yRobot--
			}

			// store coordinate
			fmt.Printf("%d: %s\n", index, string(dir))
			if _, ok := coordinates[fmt.Sprintf("%d,%d", xRobot, yRobot)]; !ok {
				coordinates[fmt.Sprintf("%d,%d", xRobot, yRobot)] = 1
				fmt.Printf("%d,%d Robot NOT %d\n", xRobot, yRobot, len(coordinates))
			} else {
				fmt.Printf("%d,%d Robot FOUND\n", xRobot, yRobot)
			}
		}
	}

	// once complete count the entries in the map and return the count
	return len(coordinates), nil
}
