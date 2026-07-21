package integrationtestrubricbuilder

import "github.com/trunkform/trunkform/mcp/rubricHandlers/rubrichandler"

func NewRubricHandler() rubrichandler.RubricHandler {
	return rubrichandler.RubricHandler{
		HandlerKey:  "integration-test-implemented",
		HandlerName: "Add Integration Testing",
		HandlerText: rubrichandler.InstructionPrefix + rubrichandler.AskTheUser(`INTEGRATION TESTING: What additional services need to be integration tested?`) + `
Then help the user implement integration tests that verify service interactions. After the integration tests pass, update .rubric.integration-test-implemented to the exact command used to run the passing integration tests, then reread ./trunkform.json and call trunkform again with the full object. ` + rubrichandler.CriticalPrefix + ` and tested locally or in 
CI where possible. ` + rubrichandler.NeverGuess,
		HandlerData: map[string]any{},
	}
}
