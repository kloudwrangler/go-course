package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd1(t *testing.T) {
	result := Add(2, 3)
	assert.Equal(t, 5, result, "Addition failed")
}

func TestMultiply1(t *testing.T) {
	result := Multiply(2, 3)
	assert.Equal(t, 6, result, "Multiplication failed")
}

func TestMultWithoutAssert(t *testing.T) {
	got := Multiply(2, 3)
	if got != 6 {
		t.Errorf("Multiply(2,3) = %d; want 6", got)
	}
}

// TODO: Create a function that tests the subtraction
func TestSubtract1(t *testing.T) {
	result := Subtract(3, 2)
	assert.Equal(t, 1, result, "Subtraction Failed")
}

// TODO: Create a function that tests the subtraction without the use of assert
func TestSubtractWithoutAssert(t *testing.T) {
	result := Subtract(3, 2)
	if result != 1 {
		t.Errorf("Subtract(3, 2) = %d; want 1", result)
	}
}
