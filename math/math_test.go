package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add(2, 2) // Arrange and Act
	if total != 4 {    // Assertion
		t.Errorf("Sum was incorrect, Actual: %d, Expected: %d", total, 4)
	}
	t.Log("running TestAdd")
}
