package rubrichandlers

import (
	"strings"
	"testing"
	"trunkform-mcp/rubricHandlers/cicdboilerplaterubricbuilder"
	"trunkform-mcp/rubricHandlers/ciiacrubricbuilder"
	"trunkform-mcp/rubricHandlers/integrationtestrubricbuilder"
	"trunkform-mcp/rubricHandlers/mocksrubricbuilder"
	"trunkform-mcp/rubricHandlers/perftestrubricbuilder"
	"trunkform-mcp/rubricHandlers/testdoublesrubricbuilder"
	"trunkform-mcp/rubricHandlers/unittestrubricbuilder"
)

func TestAskTheUser(t *testing.T) {
	result := AskTheUser("test question")
	expected := `Ask the user, "test question" Wait for their response before proceeding. `
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestNewRubricHandlers(t *testing.T) {
	// Test all NewRubricHandler functions to ensure 100% coverage
	handlers := []struct {
		name string
		fn   func() interface{}
	}{
		{"cicdboilerplate", func() interface{} { return cicdboilerplaterubricbuilder.NewRubricHandler() }},
		{"ciiac", func() interface{} { return ciiacrubricbuilder.NewRubricHandler() }},
		{"integrationtest", func() interface{} { return integrationtestrubricbuilder.NewRubricHandler() }},
		{"mocks", func() interface{} { return mocksrubricbuilder.NewRubricHandler() }},
		{"perftest", func() interface{} { return perftestrubricbuilder.NewRubricHandler() }},
		{"testdoubles", func() interface{} { return testdoublesrubricbuilder.NewRubricHandler() }},
		{"unittest", func() interface{} { return unittestrubricbuilder.NewRubricHandler() }},
	}

	for _, h := range handlers {
		t.Run(h.name, func(t *testing.T) {
			handler := h.fn()
			if handler == nil {
				t.Errorf("%s NewRubricHandler returned nil", h.name)
			}
		})
	}
}

func TestDefaultRubricHandlers(t *testing.T) {
	handlers := DefaultRubricHandlers()
	expectations := map[string]string{
		"CI/CD Automation Boilerplate":                    "update .rubric.ci-cd-automation-boilerplate to \"cat <<< \\\"Boilerplate ci-cd automation was created\\\"\"",
		"Add Unit Test Coverage":                          "update .rubric.unit-test-implemented to the exact command used",
		"Add Integration Testing":                         "update .rubric.integration-test-implemented to the exact command used",
		"Add Performance Testing":                         "update .rubric.perf-test-implemented to the exact command used",
		"Add Continuous Integration IAC":                  "update .rubric.continuous-integration-iac to the exact command used",
		"Add Continuous Integration Test Doubles":         "update .rubric.continuous-integration-test-double-implemented to the exact command used",
		"Add Continuous Integration Mock Implementations": "update .rubric.continuous-integration-mocks-implemented to the exact command used",
	}

	if len(handlers) != 7 {
		t.Fatalf("expected 7 handlers, got %d", len(handlers))
	}

	for _, h := range handlers {
		t.Run(h.Name(), func(t *testing.T) {
			result := h.Build()
			if result.Text == "" {
				t.Error("handler returned empty text")
			}
			if result.Structured == nil {
				t.Error("handler returned nil structured data")
			}
			if !strings.HasPrefix(result.Text, InstructionPrefix) {
				t.Errorf("%s does not start with InstructionPrefix", h.Name())
			}
			if !strings.Contains(result.Text, expectations[h.Name()]) {
				t.Errorf("%s should require storing the exact test command", h.Name())
			}
		})
	}
}
