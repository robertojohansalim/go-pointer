package gopointer

import "testing"

func TestPointerizeInt(t *testing.T) {
	ptrIntVariable := Pointerize(10)

	if ptrIntVariable == nil || *ptrIntVariable != 10 {
		t.Error("Error converting int variable")
	}
}

func TestPointerizeString(t *testing.T) {
	ptrStrVariable := Pointerize("Test String")

	if ptrStrVariable == nil || *ptrStrVariable != "Test String" {
		t.Error("Error converting int variable")
	}

	actualStrVariable := "Test String"

	ptrStrVariable = Pointerize(actualStrVariable)

	if ptrStrVariable == nil || ptrStrVariable == &actualStrVariable {
		t.Error("New pointer expected to create a copy of the actual passed variable")
	}
}
