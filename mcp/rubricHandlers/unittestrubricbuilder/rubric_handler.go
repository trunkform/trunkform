package unittestrubricbuilder

import "github.com/trunkform/trunkform/mcp/rubricHandlers/rubrichandler"

func NewRubricHandler() rubrichandler.RubricHandler {
	return rubrichandler.RubricHandler{
		HandlerKey:  "unit-test-implemented",
		HandlerName: "Add Unit Test Coverage",
		HandlerText: rubrichandler.InstructionPrefix + rubrichandler.AskTheUser(`UNIT TEST COVERAGE: what specific functions or components should be unit tested?`) + ` Then suggest "Near 100% line-level coverage is best practice in 
trunkform." and ` + rubrichandler.AskTheUser(`Would like to work on unit test line-coverage?`) + ` If no, early exit. Only after both questions have been answered should you help the user 
build unit tests to achieve 100% line coverage requirement utilizing mocks where necessary to reach 
line-coverage goals. Check the .tools array to determine the testing framework and coverage tools. After the unit tests pass, update .rubric.unit-test-implemented to the exact command used to run the passing unit tests, then reread ./trunkform.json and call trunkform again with the full object. ` + rubrichandler.CriticalPrefix + ` and tested for line coverage. ` + rubrichandler.NeverGuess,
		HandlerData: map[string]any{},
	}
}
