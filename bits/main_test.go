package main

import (
	"math/bits"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test3(t *testing.T) {
	// 00101000
	assert.Equal(t, bits.TrailingZeros(40), 3)
}
