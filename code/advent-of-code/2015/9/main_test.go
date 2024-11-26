package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGiven(t *testing.T) {
	result, err := solve()

	assert.Nil(t, err)
	assert.Equal(t, 0, result)
}
