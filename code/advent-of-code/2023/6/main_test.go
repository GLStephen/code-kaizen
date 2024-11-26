package main

import "testing"

func TestKnownSolve(t *testing.T) {
	input := inputValues{
		records: []recordTiming{
			{
				time:     7,
				distance: 9,
			},
			{
				time:     15,
				distance: 40,
			},
			{
				time:     30,
				distance: 200,
			},
		},
	} // Provide the necessary input values for testing
	expected := 288 // Provide the expected output value

	result := solve(input)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestGoalSolve(t *testing.T) {
	input := inputValues{
		records: []recordTiming{
			{
				time:     54,
				distance: 446,
			},
			{
				time:     81,
				distance: 1292,
			},
			{
				time:     70,
				distance: 1035,
			},
			{
				time:     88,
				distance: 1007,
			},
		},
	} // Provide the necessary input values for testing
	expected := 2065338 // Provide the expected output value

	result := solve(input)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestKnownBonusSolve(t *testing.T) {
	input := inputValues{
		records: []recordTiming{
			{
				time:     71530,
				distance: 940200,
			},
		},
	} // Provide the necessary input values for testing
	expected := 71503 // Provide the expected output value

	result := solve(input)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// There's really only one race - ignore the spaces between the numbers on each line.
func TestGoalBonusSolve(t *testing.T) {
	input := inputValues{
		records: []recordTiming{
			{
				time:     54817088,
				distance: 446129210351007,
			},
		},
	} // Provide the necessary input values for testing
	expected := 34934171 // Provide the expected output value

	result := solve(input)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
