package cicdunittestingbuilder

import "testing"

func TestNewRubricHandler(t *testing.T) {
	handler := NewRubricHandler()
	if handler.Key() == "" {
		t.Error("handler key should not be empty")
	}
	if handler.Name() == "" {
		t.Error("handler name should not be empty")
	}
}
