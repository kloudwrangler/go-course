package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	assert.Equal(t, 5, result, "Addition failed")
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	assert.Equal(t, 6, result, "Multiplication failed")
}

func TestMultWithoutAssert(t *testing.T) {
	got := Multiply(2, 3)
	if got != 6 {
		t.Errorf("Multiply(2,3) = %d; want 6", got)
	}
}

// TODO: 2. Create a function that tests the subtraction
