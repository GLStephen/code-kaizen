package main

import "fmt"

type problemInput struct {
	histories []measurementHistory
}
type measurementHistory struct {
	measurements []int // remember to sort this
}

func main() {
	input := problemInput{}
	result := solve(input)

	fmt.Printf("No Nain Result, Run Tests: %d\n", result)
}

func solve(input problemInput) int {
	// I think I can just summarize the differences between each measurement and not correlate them to the history but need to verify
	//allDifferences := []measurementHistory{}
	for _, history := range input.histories {
		differences := []measurementHistory{}
		findDifferences(&differences, history)
		fmt.Printf("History: %+v Diffs: %+v\n", history, differences)
	}

	return 0
}

func findDifferences(set *[]measurementHistory, history measurementHistory) {
	differences := measurementHistory{measurements: []int{}} // sort?

	for i := 1; i < len(history.measurements); i++ {
		differences.measurements = append(differences.measurements, history.measurements[i]-history.measurements[i-1])
	}
	*set = append(*set, differences)

	sum := 0
	for _, differences := range differences.measurements {
		sum = sum + differences
	}
	if sum != 0 {
		findDifferences(set, differences)
	}
}
