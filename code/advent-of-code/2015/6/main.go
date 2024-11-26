package main

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	command int // 1 = on, 0 = off, -1 = toggle
	startX  int
	startY  int
	endX    int
	endY    int
}

func main() {
	fmt.Println("Hello from 2015 Day 6")
}

func solve(input []string) (int, error) {
	pixels := initPixelArray()

	for _, instructionLine := range input {
		instruction := getInstruction(instructionLine)
		fmt.Printf("Instruction: %+v %s\n", instruction, instructionLine)
		for x := instruction.startX; x <= instruction.endX; x++ {
			for y := instruction.startY; y <= instruction.endY; y++ {
				pixels[x][y] = applyCommand(instruction.command, pixels[x][y])
				//fmt.Printf("%d ", pixels[x][y])
			}
			//fmt.Println()
		}
	}

	totalOn := 0

	for x := 0; x <= 999; x++ {
		for y := 0; y <= 999; y++ {
			if pixels[x][y] == 1 {
				totalOn++
			}
		}
	}

	return totalOn, nil
}

func applyCommand(command int, value int) int {
	if command == 1 {
		return 1
	} else if command == 0 {
		return 0
	} else if command == -1 {
		if value == 1 {
			return 0
		} else {
			return 1
		}
	}
	return 0
}

func initPixelArray() [][]int {
	pixels := make([][]int, 1000)
	for i := range pixels {
		pixels[i] = make([]int, 1000)
	}

	for x := 0; x <= 999; x++ {
		for y := 0; y <= 999; y++ {
			if pixels[x][y] == 1 {
				pixels[x][y] = 0
			}
		}
	}

	return pixels
}

func getInstruction(input string) instruction {
	instruction := instruction{}
	startIndex := 2
	endIndex := 4
	if strings.Contains(input, "on") {
		instruction.command = 1
	} else if strings.Contains(input, "off") {
		instruction.command = 0
	} else if strings.Contains(input, "toggle") {
		instruction.command = -1

		startIndex = 1
		endIndex = 3
	}

	spaceSplit := strings.Split(input, " ")
	start := strings.Split(spaceSplit[startIndex], ",")
	end := strings.Split(spaceSplit[endIndex], ",")

	instruction.startX, _ = strconv.Atoi(start[0])
	instruction.startY, _ = strconv.Atoi(start[1])

	instruction.endX, _ = strconv.Atoi(end[0])
	instruction.endY, _ = strconv.Atoi(end[1])

	return instruction
}

func solveBonus(input []string) (int, error) {
	pixels := initPixelArray()

	for _, instructionLine := range input {
		//fmt.Printf("Instruction: %s ", instructionLine)
		instruction := getInstruction(instructionLine)
		//fmt.Printf("Instruction: %+v %s\n", instruction, instructionLine)
		for x := instruction.startX; x <= instruction.endX; x++ {
			for y := instruction.startY; y <= instruction.endY; y++ {
				pixels[x][y] = pixels[x][y] + applyBonusCommand(instruction.command, pixels[x][y])
				//fmt.Printf(" %d\n", pixels[x][y])
			}
			//fmt.Println()
		}
	}

	totalOn := 0

	for x := 0; x <= 999; x++ {
		for y := 0; y <= 999; y++ {
			totalOn = totalOn + pixels[x][y]
		}
	}

	return totalOn, nil
}

func applyBonusCommand(command int, value int) int {
	//fmt.Printf("Command: %d", command)
	if command == 1 { // on
		return 1
	} else if command == 0 { // off
		if value > 0 {
			return -1
		} else {
			return 0
		}
	} else { // toggle
		return 2
	}
}
