package steeringrubricbuilder

import "testing"

func TestNewRubricHandler(t *testing.T) {
	handler := NewRubricHandler()
	if handler.Key() == "" {
		t.Error("handler key should not be empty")
	}
}
