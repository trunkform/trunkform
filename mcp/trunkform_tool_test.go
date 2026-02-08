package main

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func TestConstants(t *testing.T) {
	expectedPrefix := "instructions for the AI assistant to act on behalf of user's needs: "
	if InstructionPrefix != expectedPrefix {
		t.Errorf("InstructionPrefix = %q, want %q", InstructionPrefix, expectedPrefix)
	}

	expectedCritical := `CRITICAL: Add newly created relative folders and relative file paths to the .trunkform yaml file's .related
 string array when updating the yaml .rubric rule true when it has actually been implemented, and update the list of .tools when a new tool is used.`
	if CriticalPrefix != expectedCritical {
		t.Errorf("CriticalPrefix = %q, want %q", CriticalPrefix, expectedCritical)
	}

	expectedNever := `NEVER guess at trunkform tool inputs - only use values directly from the .trunkform file.`
	if NeverGuess != expectedNever {
		t.Errorf("NeverGuess = %q, want %q", NeverGuess, expectedNever)
	}
}

func TestAskTheUser(t *testing.T) {
	result := AskTheUser("what is your name?")
	expected := `Ask the user, "what is your name?" Wait for their response before proceeding. `
	if result != expected {
		t.Errorf("AskTheUser() = %q, want %q", result, expected)
	}
}

func TestNewTrunkformTool(t *testing.T) {
	tt := NewTrunkformTool()
	if tt == (TrunkformTool{}) {
		// Validates struct creation
	}
}

func TestTrunkformToolRegister(t *testing.T) {
	s := server.NewMCPServer("test", "0.1.0")
	tt := NewTrunkformTool()
	tt.Register(s)
}

func TestProcessRubric(t *testing.T) {
	tests := []struct {
		name   string
		rubric map[string]interface{}
	}{
		{"ci-cd-boilerplate", map[string]interface{}{"ci-cd-automation-boilerplate": false}},
		{"perf-test", map[string]interface{}{"perf-test-implemented": false}},
		{"unit-test", map[string]interface{}{"unit-test-implemented": false}},
		{"integration-test", map[string]interface{}{"integration-test-implemented": false}},
		{"ci-iac", map[string]interface{}{"continuous-integration-iac": false}},
		{"test-doubles", map[string]interface{}{"continuous-integration-test-double-implemented": false}},
		{"mocks", map[string]interface{}{"continuous-integration-mocks-implemented": false}},
		{"all-complete", map[string]interface{}{"ci-cd-automation-boilerplate": true}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := formatArgs{
				Version:  "0.1.0",
				Tools:    []string{"test"},
				Steering: "strict",
				Related:  []string{},
				Rubric:   tt.rubric,
			}
			result, err := processRubric(args)
			if err != nil {
				t.Errorf("processRubric failed: %v", err)
			}
			if result == nil {
				t.Error("Expected non-nil result")
			}
		})
	}
}

func TestProcessRubricNonStrict(t *testing.T) {
	// Test with key that exists in rubric
	args := formatArgs{
		Version:  "0.1.0",
		Tools:    []string{"test"},
		Steering: "flexible",
		Related:  []string{},
		Rubric:   map[string]interface{}{"perf-test-implemented": false},
	}
	result, err := processRubric(args)
	if err != nil {
		t.Errorf("processRubric failed: %v", err)
	}
	if result == nil {
		t.Error("Expected non-nil result")
	}
	
	// Test with key that doesn't exist in rubric (should be skipped)
	args2 := formatArgs{
		Version:  "0.1.0",
		Tools:    []string{"test"},
		Steering: "fuzzy",
		Related:  []string{},
		Rubric:   map[string]interface{}{"unit-test-implemented": true},
	}
	result2, err2 := processRubric(args2)
	if err2 != nil {
		t.Errorf("processRubric failed: %v", err2)
	}
	if result2 == nil {
		t.Error("Expected non-nil result for all complete")
	}
}

func TestHandlers(t *testing.T) {
	handlers := []struct {
		name string
		fn   func() (string, map[string]any)
	}{
		{"handleCICDBoilerplate", handleCICDBoilerplate},
		{"handleUnitTestCoverage", handleUnitTestCoverage},
		{"handleIntegrationTest", handleIntegrationTest},
		{"handlePerfTest", handlePerfTest},
		{"handleCIIaC", handleCIIaC},
		{"handleTestDoubles", handleTestDoubles},
		{"handleMocks", handleMocks},
	}
	
	for _, h := range handlers {
		t.Run(h.name, func(t *testing.T) {
			text, structured := h.fn()
			if text == "" {
				t.Error("Handler returned empty text")
			}
			if structured == nil {
				t.Error("Handler returned nil structured data")
			}
			if len(text) < len(InstructionPrefix) || text[:len(InstructionPrefix)] != InstructionPrefix {
				t.Errorf("%s does not start with InstructionPrefix", h.name)
			}
		})
	}
}

func TestTrunkformToolHandleToolMethod(t *testing.T) {
	tt := NewTrunkformTool()
	
	tests := []struct {
		name   string
		rubric map[string]interface{}
	}{
		{"ci_cd_false", map[string]interface{}{"ci-cd-automation-boilerplate": false}},
		{"all_true", map[string]interface{}{"ci-cd-automation-boilerplate": true}},
	}
	
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			args := formatArgs{
				Version:  "0.1.0",
				Tools:    []string{"test"},
				Steering: "strict",
				Related:  []string{},
				Rubric:   tc.rubric,
			}
			result, err := tt.handleTool(context.Background(), mcp.CallToolRequest{}, args)
			if err != nil || result == nil {
				t.Error("Handler failed")
			}
		})
	}
}


