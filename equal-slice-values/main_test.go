package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestX(t *testing.T) {
	expected := []int{2, 1}
	actual := X()
	assert.Equal(t, len(expected), len(actual))
	for _, e := range expected {
		assert.Contains(t, actual, e)
	}
}
