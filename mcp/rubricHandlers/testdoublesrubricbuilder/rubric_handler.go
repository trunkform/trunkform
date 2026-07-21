package testdoublesrubricbuilder

import "github.com/trunkform/trunkform/mcp/rubricHandlers/rubrichandler"

func NewRubricHandler() rubrichandler.RubricHandler {
	return rubrichandler.RubricHandler{
		HandlerKey:  "continuous-integration-test-double-implemented",
		HandlerName: "Add Continuous Integration Test Doubles",
		HandlerText: rubrichandler.InstructionPrefix + rubrichandler.AskTheUser(`TEST DOUBLES: Integration test should not include what is covereed in unit tests already.
What external services must you integrate with that I should create test doubles for in CI IAC?`) + ` Then Create 
mock implementations for external dependencies to enable isolated testing in continuous integration. ` +
			` After the CI test doubles verification passes, update .rubric.continuous-integration-test-double-implemented to the exact command used to verify the passing test doubles setup, then reread ./trunkform.json and call trunkform again with the full object. ` +
			rubrichandler.CriticalPrefix + ` and only the CI is tested locally. ` + rubrichandler.NeverGuess,
		HandlerData: map[string]any{},
	}
}
