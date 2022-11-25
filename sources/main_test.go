package main

import (
	"testing"
)

func Test_sumSubarrayMins(t *testing.T) {
	input := []int{3, 1, 2, 4}
	expected := 17
	actual := sumSubarrayMins(input)
	if expected != actual {
		t.Fatalf("Input: %v, Expected: %v, Actual: %v", input, expected, actual)
	}
}
