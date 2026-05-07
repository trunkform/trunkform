package main

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	rubrichandlers "trunkform-mcp/rubricHandlers"
	"trunkform-mcp/rubricHandlers/rubrichandler"
)

func TestNewTrunkformTool(t *testing.T) {
	_ = NewTrunkformTool()
}

func TestToolDescription(t *testing.T) {
	expected := `
The trunkform tool is a software delivery checklisting tool, often used as a verb; "trunkform."

# USAGE INSTRUCTIONS:

First, read the contents of the ./trunkform.json file on disk and pass it as the argument. If the ./trunkform.json file does not exist, create the ./trunkform.json file with the following schema and default values:

` + "```json" + `
{
  "version": "0.1.0",
  "tools": [],
  "related": [],
  "rubric": {
  }
}
` + "```" + `
`
	if ToolDescription != expected {
		t.Errorf("ToolDescription mismatch:\ngot: %q\nwant: %q", ToolDescription, expected)
	}
}

func TestTrunkformToolRegister(t *testing.T) {
	s := server.NewMCPServer("test", "0.1.0")
	tt := NewTrunkformTool()
	tt.Register(s)

	tool := s.GetTool("trunkform")
	if tool == nil {
		t.Fatal("expected trunkform tool to be registered")
	}

	if len(tool.Tool.InputSchema.Properties) != 4 {
		t.Fatalf("expected 4 top-level parameters, got %d", len(tool.Tool.InputSchema.Properties))
	}

	if len(tool.Tool.InputSchema.Required) != 0 {
		t.Fatalf("expected no required parameters, got %v", tool.Tool.InputSchema.Required)
	}

	// Check that we have the expected properties
	expectedProps := []string{"version", "tools", "related", "rubric"}
	for _, prop := range expectedProps {
		if _, ok := tool.Tool.InputSchema.Properties[prop]; !ok {
			t.Fatalf("expected property %s not found", prop)
		}
	}

	rubricProp, ok := tool.Tool.InputSchema.Properties["rubric"].(map[string]any)
	if !ok {
		t.Fatal("expected rubric property schema")
	}

	additionalProperties, ok := rubricProp["additionalProperties"].(map[string]any)
	if !ok {
		t.Fatal("expected rubric additionalProperties schema")
	}

	types, ok := additionalProperties["type"].([]string)
	if !ok {
		t.Fatalf("expected rubric item types to be []string, got %T", additionalProperties["type"])
	}

	if len(types) != 2 || types[0] != "string" || types[1] != "null" {
		t.Fatalf("expected rubric item types [string null], got %v", types)
	}
}

func TestProcessRubric(t *testing.T) {
	tests := []struct {
		name     string
		rubric   map[string]*string
		contains string
	}{
		{"steering-test", map[string]*string{
			"steering-implemented": nil,
		}, "Add Steering"},
		{"lint-test", map[string]*string{
			"steering-implemented": strPtr("cat ./AGENTS.md"),
			"lint-implemented":     nil,
		}, "Add Linting"},
		{"unit-test", map[string]*string{
			"steering-implemented":  strPtr("cat ./AGENTS.md"),
			"lint-implemented":      strPtr("make"),
			"unit-test-implemented": nil,
		}, "Add Unit Test Coverage"},
		{"ci-cd-boilerplate", map[string]*string{
			"steering-implemented":         strPtr("cat ./AGENTS.md"),
			"lint-implemented":             strPtr("make"),
			"unit-test-implemented":        strPtr("make unit-test"),
			"ci-cd-automation-boilerplate": nil,
		}, "CI/CD Automation Boilerplate"},
		{"ci-cd-linting", map[string]*string{
			"steering-implemented":         strPtr("cat ./AGENTS.md"),
			"lint-implemented":             strPtr("make"),
			"unit-test-implemented":        strPtr("make unit-test"),
			"ci-cd-automation-boilerplate": strPtr("cat ../.github/workflows/mcp.yml | grep -- 'max-parallel: 1'"),
		}, "CI/CD Linting Automation"},
		{"ci-cd-unit-testing", map[string]*string{
			"steering-implemented":         strPtr("cat ./AGENTS.md"),
			"lint-implemented":             strPtr("make"),
			"unit-test-implemented":        strPtr("make unit-test"),
			"ci-cd-automation-boilerplate": strPtr("cat ../.github/workflows/mcp.yml | grep -- 'max-parallel: 1'"),
			"ci-cd-linting":                strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
		}, "CI/CD Unit Testing Automation"},
		{"perf-test", map[string]*string{
			"steering-implemented":                           strPtr("cat ./AGENTS.md"),
			"lint-implemented":                               strPtr("make"),
			"unit-test-implemented":                          strPtr("make unit-test"),
			"ci-cd-automation-boilerplate":                   strPtr("cat ../.github/workflows/mcp.yml | grep -- 'max-parallel: 1'"),
			"ci-cd-linting":                                  strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"ci-cd-unit-testing":                             strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"continuous-integration-iac":                     strPtr("cat ../.github/workflows/mcp.yml | grep -- 'terragrunt run --all apply'"),
			"continuous-integration-test-double-implemented": strPtr("make local-int-test-http"),
			"continuous-integration-mocks-implemented":       strPtr("make local-int-test-http"),
			"integration-test-implemented":                   strPtr("make local-int-test-http"),
			"perf-test-implemented":                          nil,
		}, "Add Performance Testing"},
		{"integration-test", map[string]*string{
			"steering-implemented":                           strPtr("cat ./AGENTS.md"),
			"lint-implemented":                               strPtr("make"),
			"unit-test-implemented":                          strPtr("make unit-test"),
			"ci-cd-automation-boilerplate":                   strPtr("cat ../.github/workflows/mcp.yml | grep -- 'max-parallel: 1'"),
			"ci-cd-linting":                                  strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"ci-cd-unit-testing":                             strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"continuous-integration-iac":                     strPtr("cat ../.github/workflows/mcp.yml | grep -- 'terragrunt run --all apply'"),
			"continuous-integration-test-double-implemented": strPtr("make local-int-test-http"),
			"continuous-integration-mocks-implemented":       strPtr("make local-int-test-http"),
			"integration-test-implemented":                   nil,
		}, "Add Integration Testing"},
		{"ci-iac", map[string]*string{
			"steering-implemented":         strPtr("cat ./AGENTS.md"),
			"lint-implemented":             strPtr("make"),
			"unit-test-implemented":        strPtr("make unit-test"),
			"ci-cd-automation-boilerplate": strPtr("cat ../.github/workflows/mcp.yml | grep -- 'max-parallel: 1'"),
			"ci-cd-linting":                strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"ci-cd-unit-testing":           strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"continuous-integration-iac":   nil,
		}, "Add Continuous Integration IAC"},
		{"test-doubles", map[string]*string{
			"steering-implemented":                           strPtr("cat ./AGENTS.md"),
			"lint-implemented":                               strPtr("make"),
			"unit-test-implemented":                          strPtr("make unit-test"),
			"ci-cd-automation-boilerplate":                   strPtr("cat ../.github/workflows/mcp.yml | grep -- 'max-parallel: 1'"),
			"ci-cd-linting":                                  strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"ci-cd-unit-testing":                             strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"continuous-integration-iac":                     strPtr("cat ../.github/workflows/mcp.yml | grep -- 'terragrunt run --all apply'"),
			"continuous-integration-test-double-implemented": nil,
		}, "Add Continuous Integration Test Doubles"},
		{"mocks", map[string]*string{
			"steering-implemented":                           strPtr("cat ./AGENTS.md"),
			"lint-implemented":                               strPtr("make"),
			"unit-test-implemented":                          strPtr("make unit-test"),
			"ci-cd-automation-boilerplate":                   strPtr("cat ../.github/workflows/mcp.yml | grep -- 'max-parallel: 1'"),
			"ci-cd-linting":                                  strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"ci-cd-unit-testing":                             strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"continuous-integration-iac":                     strPtr("cat ../.github/workflows/mcp.yml | grep -- 'terragrunt run --all apply'"),
			"continuous-integration-test-double-implemented": strPtr("make local-int-test-http"),
			"continuous-integration-mocks-implemented":       nil,
		}, "Add Continuous Integration Mock Implementations"},
		{"all-complete", map[string]*string{
			"steering-implemented":                           strPtr("cat ./AGENTS.md"),
			"perf-test-implemented":                          strPtr("make local-perf-test"),
			"lint-implemented":                               strPtr("make"),
			"unit-test-implemented":                          strPtr("make unit-test"),
			"ci-cd-automation-boilerplate":                   strPtr("cat ../.github/workflows/mcp.yml | grep -- 'max-parallel: 1'"),
			"ci-cd-linting":                                  strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"ci-cd-unit-testing":                             strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"continuous-integration-iac":                     strPtr("cat ../.github/workflows/mcp.yml | grep -- 'terragrunt run --all apply'"),
			"continuous-integration-test-double-implemented": strPtr("make local-int-test-http"),
			"continuous-integration-mocks-implemented":       strPtr("make local-int-test-http"),
			"integration-test-implemented":                   strPtr("make local-int-test-http"),
		}, `all rubric items tested and complete. yolo 🚀`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := trunkformSchema{
				Version: "0.1.0",
				Tools:   []string{"test"},
				Related: []string{},
				Rubric:  tt.rubric,
			}
			result, err := processRubric(args)
			if err != nil {
				t.Errorf("processRubric failed: %v", err)
			}
			if result == nil {
				t.Fatal("Expected non-nil result")
			}
			resultText := result.Content[0].(mcp.TextContent).Text
			if !contains(resultText, tt.contains) {
				t.Errorf("result text %q does not contain %q", resultText, tt.contains)
			}
			if tt.name == "all-complete" && !contains(resultText, "If they do not require additional integration testing, then explicitly execute every command stored in ./trunkform.json .rubric for the keys in this exact order:") {
				t.Errorf("result text %q does not contain ordered rubric execution instructions", resultText)
			}
				if tt.name == "all-complete" && !contains(resultText, "would you like suggestions on how to remediate the failing test?") {
					t.Errorf("result text %q does not contain remediation prompt instructions", resultText)
				}
				if tt.name == "all-complete" && !contains(resultText, `1. steering-implemented\n2. lint-implemented\n3. unit-test-implemented\n4. ci-cd-automation-boilerplate\n5. ci-cd-linting\n6. ci-cd-unit-testing\n7. continuous-integration-iac\n8. continuous-integration-test-double-implemented\n9. continuous-integration-mocks-implemented\n10. integration-test-implemented\n11. perf-test-implemented`) {
					t.Errorf("result text %q does not contain the expected rubric execution order", resultText)
				}
				if tt.name == "all-complete" && !contains(resultText, `Ask the user, \"is less than 100% line coverage acceptable for these changes, or is the test command missing a coverage flag?\" Wait for their response before proceeding. `) {
					t.Errorf("result text %q does not contain coverage prompt instructions", resultText)
				}
			if tt.name == "all-complete" && !contains(resultText, "Ask the user, \\\"do the changes\\nrequire updates to integration tests?\\\" Wait for their response before proceeding. ") {
				t.Errorf("result text %q does not contain integration test prompt instructions", resultText)
			}
			if tt.name == "all-complete" {
				integrationPrompt := "Ask the user, \\\"do the changes\\nrequire updates to integration tests?\\\" Wait for their response before proceeding."
				commandExecution := "If they do not require additional integration testing, then explicitly execute every command stored in ./trunkform.json .rubric for the keys in this exact order:"
				if strings.Index(resultText, integrationPrompt) > strings.Index(resultText, commandExecution) {
					t.Errorf("result text %q does not place the integration test prompt before rubric execution", resultText)
				}
			}
		})
	}
}

func TestProcessRubricMissingKeys(t *testing.T) {
	args := trunkformSchema{
		Version: "0.1.0",
		Tools:   []string{"test"},
		Related: []string{},
		Rubric: map[string]*string{
			"steering-implemented":  strPtr("cat ./AGENTS.md"),
			"perf-test-implemented": strPtr("make local-perf-test"),
		},
	}
	result, err := processRubric(args)
	if err != nil {
		t.Errorf("processRubric failed: %v", err)
	}
	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	resultText := result.Content[0].(mcp.TextContent).Text
	if !contains(resultText, "Add Linting") {
		t.Errorf("result text %q does not contain %q", resultText, "Add Linting")
	}

	args2 := trunkformSchema{
		Version: "0.1.0",
		Tools:   []string{"test"},
		Related: []string{},
		Rubric: map[string]*string{
			"steering-implemented":                           strPtr("cat ./AGENTS.md"),
			"lint-implemented":                               strPtr("make"),
			"unit-test-implemented":                          strPtr("make unit-test"),
			"ci-cd-automation-boilerplate":                   strPtr("cat ../.github/workflows/mcp.yml | grep -- 'max-parallel: 1'"),
			"ci-cd-linting":                                  strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"ci-cd-unit-testing":                             strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"continuous-integration-iac":                     strPtr("cat ../.github/workflows/mcp.yml | grep -- 'terragrunt run --all apply'"),
			"continuous-integration-test-double-implemented": strPtr("make local-int-test-http"),
			"continuous-integration-mocks-implemented":       strPtr("make local-int-test-http"),
			"integration-test-implemented":                   strPtr("make local-int-test-http"),
		},
	}
	result2, err2 := processRubric(args2)
	if err2 != nil {
		t.Errorf("processRubric failed: %v", err2)
	}
	if result2 == nil {
		t.Fatal("Expected non-nil result for missing leading key")
	}
	resultText2 := result2.Content[0].(mcp.TextContent).Text
	if !contains(resultText2, "Add Performance Testing") {
		t.Errorf("result text %q does not contain %q", resultText2, "Add Performance Testing")
	}
}

func TestTrunkformToolHandleToolMethod(t *testing.T) {
	tt := NewTrunkformTool()

	tests := []struct {
		name   string
		rubric map[string]*string
	}{
		{"ci_cd_false", map[string]*string{"ci-cd-automation-boilerplate": nil}},
		{"all_true", map[string]*string{"ci-cd-automation-boilerplate": strPtr("make")}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			args := trunkformSchema{
				Version: "0.1.0",
				Tools:   []string{"test"},
				Related: []string{},
				Rubric:  tc.rubric,
			}
			result, err := tt.handleTool(context.Background(), mcp.CallToolRequest{}, args)
			if err != nil || result == nil {
				t.Error("Handler failed")
			}
		})
	}
}

func TestHandleToolEarlyExit(t *testing.T) {
	tt := NewTrunkformTool()
	result, _ := tt.handleTool(context.Background(), mcp.CallToolRequest{}, trunkformSchema{})
	expected := map[string]any{"instructions": InstructionPrefix + ToolDescription + NeverGuess}
	jsonBytes, _ := json.Marshal(expected)
	if result.Content[0].(mcp.TextContent).Text != string(jsonBytes) {
		t.Error("Early exit content mismatch")
	}
}

func TestToolsArrayCheck(t *testing.T) {
	var lintTestHandler rubrichandler.RubricHandler
	for _, handler := range rubrichandlers.DefaultRubricHandlers() {
		if handler.Key() == "lint-implemented" {
			lintTestHandler = handler
			break
		}
	}
	if lintTestHandler.Key() != "lint-implemented" {
		t.Fatal("expected lint strategy")
	}
	if !contains(lintTestHandler.Build().Text, "help the user set up a linter appropriate for the language") {
		t.Error("handleLintCoverage should check tools array")
	}

	args := trunkformSchema{
		Version: "0.1.0",
		Tools:   []string{"go"},
		Related: []string{},
		Rubric: map[string]*string{
			"steering-implemented":                           strPtr("cat ./AGENTS.md"),
			"lint-implemented":                               strPtr("make"),
			"unit-test-implemented":                          strPtr("make unit-test"),
			"ci-cd-automation-boilerplate":                   strPtr("cat ../.github/workflows/mcp.yml | grep -- 'max-parallel: 1'"),
			"ci-cd-linting":                                  strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"ci-cd-unit-testing":                             strPtr("cat ../.github/workflows/mcp.yml | grep -- 'run: make'"),
			"continuous-integration-iac":                     strPtr("cat ../.github/workflows/mcp.yml | grep -- 'terragrunt run --all apply'"),
			"continuous-integration-test-double-implemented": strPtr("make local-int-test-http"),
			"continuous-integration-mocks-implemented":       strPtr("make local-int-test-http"),
			"integration-test-implemented":                   strPtr("make local-int-test-http"),
			"perf-test-implemented":                          strPtr("make local-perf-test"),
		},
	}
	result, _ := processRubric(args)
	resultText := result.Content[0].(mcp.TextContent).Text
	if !contains(resultText, `If unit-testcoverage is 100%, has no measurable coverage (\"Unknown%\" means \"no measureable coverage\"), then tell the user that \"all rubric items tested and complete. yolo 🚀\"`) {
		t.Error("processRubric completion should include the updated coverage guidance")
	}
	if !contains(resultText, `Ask the user, \"is less than 100% line coverage acceptable for these changes, or is the test command missing a coverage flag?\" Wait for their response before proceeding.`) {
		t.Error("processRubric completion should include the updated coverage question")
	}
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func strPtr(s string) *string {
	return &s
}
