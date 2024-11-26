package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	input := "yzbqklnj"

	count, err := solve(input)

	assert.Nil(t, err)
	assert.Equal(t, 0, count)
}
