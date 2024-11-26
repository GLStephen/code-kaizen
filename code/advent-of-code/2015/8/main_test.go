package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestGiven(t *testing.T) {
// 	// 	"" is 2 characters of code (the two double quotes), but the string contains zero characters.
// 	// "abc" is 5 characters of code, but 3 characters in the string data.
// 	// "aaa\"aaa" is 10 characters of code, but the string itself contains six "a" characters and a single, escaped quote character, for a total of 7 characters in the string data.
// 	// "\x27" is 6 characters of code, but the string itself contains just one - an apostrophe ('), escaped using hexadecimal notation.

// 	count, _, err := solve([]string{`""`})

// 	assert.Nil(t, err)
// 	assert.Equal(t, 0, count)

// 	count, _, err = solve([]string{`"abc"`})

// 	assert.Nil(t, err)
// 	assert.Equal(t, 3, count)

// 	count, _, err = solve([]string{`"aaa\"aaa"`})

// 	assert.Nil(t, err)
// 	assert.Equal(t, 7, count)

// 	count, _, err = solve([]string{`"\x27`})

// 	assert.Nil(t, err)
// 	assert.Equal(t, 1, count)

// 	count, _, err = solve([]string{`"\x27aaa\x24"`})

// 	assert.Nil(t, err)
// 	assert.Equal(t, 5, count)

// 	count, _, err = solve([]string{`"\\"`})

// 	assert.Nil(t, err)
// 	assert.Equal(t, 1, count)

// 	//"aixfk\xc0iom\x21vueob"
// 	count, _, err = solve([]string{`"aixfk\xc0iom\x21vueob"`})

// 	assert.Nil(t, err)
// 	assert.Equal(t, 15, count)
// }

func TestBonusGiven(t *testing.T) {
	// 	"" encodes to "\"\"", an increase from 2 characters to 6.
	// "abc" encodes to "\"abc\"", an increase from 5 characters to 9.
	// "aaa\"aaa" encodes to "\"aaa\\\"aaa\"", an increase from 10 characters to 16.
	// "\x27" encodes to "\"\\x27\"", an increase from 6 characters to 11.

	_, reencode, err := solve([]string{`""`})

	assert.Nil(t, err)
	assert.Equal(t, 6, reencode)
}

func TestSolve(t *testing.T) {
	input, _ := loadInput()

	count, reencoded, err := solve(input)

	assert.Nil(t, err)
	assert.Equal(t, 1350, count)
	assert.Equal(t, 0, reencoded)
}

func loadInput() ([]string, error) {
	results := make([]string, 0)
	file, err := os.Open("./input.txt")
	if err != nil {
		return results, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		results = append(results, line)
		//fmt.Printf("%d: %s\n", count, line)
		count++
	}

	return results, nil
}
