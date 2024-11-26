package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGiven(t *testing.T) {
	result, err := solve("1")

	assert.Nil(t, err)
	assert.Equal(t, 312211, result)
}
