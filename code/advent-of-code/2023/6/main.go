package main

import "fmt"

type inputValues struct {
	records []recordTiming
}

type recordTiming struct {
	time     int
	distance int
	wins     int
}

func main() {
	fmt.Printf("No Real Main")
}

func solve(input inputValues) int {
	// implement the simplest brute force way to solve the problem
	// minHold := 0

	// Your toy boat has a starting speed of zero millimeters per millisecond.
	// startingSpeed := 0

	// For each whole millisecond you spend at the beginning of the race holding down the button, the boat's speed increases by one millimeter per millisecond
	// speedPerMillisecondHeld := 0

	// note: travel is 1 millimeter per second per hold

	// the max hold is time of the race

	// determine the number of ways you can beat the record in each race
	// for each hold calculate distance traveled and total time
	for index, record := range input.records {
		wins := 0
		fmt.Printf("record: %v\n", record)
		fmt.Printf("==============\nRace %d\n==============\n", record.time)
		for held := 0; held < record.time; held++ {
			remainingTime := record.time - held
			distanceTraveled := (held * remainingTime)
			//time := held

			if distanceTraveled > record.distance {
				wins++
			}
			//fmt.Printf("rate: %d remainingTime: %d traveled: %d target time: %d target distance: %d wins: %d\n", held, remainingTime, distanceTraveled, record.time, record.distance, wins)
		}
		input.records[index].wins = wins
	}

	fmt.Printf("records: %v\n", input.records)

	solution := 1
	for _, record := range input.records {
		solution *= record.wins
	}

	return solution
}
