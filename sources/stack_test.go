package main

import (
	"testing"
)

func Test_Push(t *testing.T) {
	stack := NewStack()

	input := 1
	expected := 1

	stack.Push(input)
	actual := stack.s[len(stack.s)-1]

	if expected != actual {
		t.Fatalf("Input: %v, Expected: %v, Actual: %v", input, expected, actual)
	}
}

func Test_PeekEmpty1(t *testing.T) {
	stack := NewStack()

	expected_value, expected_error := 0, &StackError{"empty stack"}
	actual_value, actual_error := stack.Peek()

	expected_len := 0
	actual_len := len(stack.s)

	if expected_value != actual_value {
		t.Fatalf("Expected: %v, Actual: %v", expected_value, actual_value)
	}

	if actual_error == nil || actual_error.Error() != expected_error.msg {
		t.Fatalf("Expected: %v, Actual: %v", expected_error, actual_error)
	}

	if expected_len != actual_len {
		t.Fatalf("Expected: %v, Actual: %v", expected_error, actual_error)
	}
}

func Test_PeekEmpty2(t *testing.T) {
	stack := NewStack()

	expected_value, expected_error := 0, &StackError{"another error msg"}
	actual_value, actual_error := stack.Peek()

	expected_len := 0
	actual_len := len(stack.s)

	if expected_value != actual_value {
		t.Fatalf("Expected: %v, Actual: %v", expected_value, actual_value)
	}

	if actual_error == nil || actual_error.Error() == expected_error.msg {
		t.Fatalf("Expected: %v, Actual: %v", expected_error, actual_error)
	}

	if expected_len != actual_len {
		t.Fatalf("Expected: %v, Actual: %v", expected_error, actual_error)
	}
}

func Test_PeekValid(t *testing.T) {
	stack := NewStack()
	stack.s = []int{1}

	expected_value := 1
	actual_value, actual_error := stack.Peek()

	expected_len := 1
	actual_len := len(stack.s)

	if expected_value != actual_value {
		t.Fatalf("Expected: %v, Actual: %v", expected_value, actual_value)
	}

	if actual_error != nil {
		t.Fatalf("Expected: %v, Actual: %v", nil, actual_error)
	}

	if expected_len != actual_len {
		t.Fatalf("Expected: %v, Actual: %v", expected_len, actual_len)
	}
}

func Test_PopEmpty1(t *testing.T) {
	stack := NewStack()

	expected_value, expected_error := 0, &StackError{"empty stack"}
	actual_value, actual_error := stack.Pop()

	if expected_value != actual_value {
		t.Fatalf("Expected: %v, Actual: %v", expected_value, actual_value)
	}

	if actual_error == nil || actual_error.Error() != expected_error.msg {
		t.Fatalf("Expected: %v, Actual: %v", expected_error, actual_error)
	}
}

func Test_PopEmpty2(t *testing.T) {
	stack := NewStack()

	expected_value, expected_error := 0, &StackError{"another stack error msg"}
	actual_value, actual_error := stack.Pop()

	if expected_value != actual_value {
		t.Fatalf("Expected: %v, Actual: %v", expected_value, actual_value)
	}

	if actual_error == nil || actual_error.Error() == expected_error.msg {
		t.Fatalf("Expected: %v, Actual: %v", expected_error, actual_error)
	}
}

func Test_PopValid(t *testing.T) {
	stack := NewStack()
	stack.s = []int{1}

	expected_value := 1
	actual_value, actual_error := stack.Pop()

	if expected_value != actual_value {
		t.Fatalf("Expected: %v, Actual: %v", expected_value, actual_value)
	}

	if actual_error != nil {
		t.Fatalf("Expected: %v, Actual: %v", nil, actual_error)
	}
}

func Test_Size(t *testing.T) {
	stack := NewStack()
	stack.s = []int{1}

	expected_value := 1
	actual_value := stack.Size()

	if expected_value != actual_value {
		t.Fatalf("Expected: %v, Actual: %v", expected_value, actual_value)
	}

	if expected_value != actual_value {
		t.Fatalf("Expected: %v, Actual: %v", expected_value, actual_value)
	}
}

func Test_TopEmpty1(t *testing.T) {
	stack := NewStack()

	expected_value, expected_error := 0, &StackError{"empty stack"}
	actual_value, actual_error := stack.Top()

	expected_len := 0
	actual_len := len(stack.s)

	if expected_value != actual_value {
		t.Fatalf("Expected: %v, Actual: %v", expected_value, actual_value)
	}

	if actual_error == nil || actual_error.Error() != expected_error.msg {
		t.Fatalf("Expected: %v, Actual: %v", expected_error, actual_error)
	}

	if expected_len != actual_len {
		t.Fatalf("Expected: %v, Actual: %v", expected_error, actual_error)
	}
}

func Test_TopEmpty2(t *testing.T) {
	stack := NewStack()

	expected_value, expected_error := 0, &StackError{"another error msg"}
	actual_value, actual_error := stack.Top()

	expected_len := 0
	actual_len := len(stack.s)

	if expected_value != actual_value {
		t.Fatalf("Expected: %v, Actual: %v", expected_value, actual_value)
	}

	if actual_error == nil || actual_error.Error() == expected_error.msg {
		t.Fatalf("Expected: %v, Actual: %v", expected_error, actual_error)
	}

	if expected_len != actual_len {
		t.Fatalf("Expected: %v, Actual: %v", expected_error, actual_error)
	}
}

func Test_TopValid(t *testing.T) {
	stack := NewStack()
	stack.s = []int{1}

	expected_value := 1
	actual_value, actual_error := stack.Top()

	expected_len := 1
	actual_len := len(stack.s)

	if expected_value != actual_value {
		t.Fatalf("Expected: %v, Actual: %v", expected_value, actual_value)
	}

	if actual_error != nil {
		t.Fatalf("Expected: %v, Actual: %v", nil, actual_error)
	}

	if expected_len != actual_len {
		t.Fatalf("Expected: %v, Actual: %v", expected_len, actual_len)
	}
}
