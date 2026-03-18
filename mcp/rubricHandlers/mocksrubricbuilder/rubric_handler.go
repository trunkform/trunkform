package mocksrubricbuilder

import "trunkform-mcp/rubricHandlers/rubrichandler"

func NewRubricHandler() rubrichandler.RubricHandler {
	return rubrichandler.RubricHandler{
		HandlerKey:  "continuous-integration-mocks-implemented",
		HandlerName: "Add Continuous Integration Mock Implementations",
		HandlerText: rubrichandler.InstructionPrefix + rubrichandler.AskTheUser(`MOCK SERVICES: what external API mocks are necessary for your change?`) + ` Then add 
the mock to CI as IAC?" After the mock-services verification passes, update .rubric.continuous-integration-mocks-implemented to the exact command used to verify the passing mock-services setup, then reread ./trunkform.json and call trunkform again with the full object. ` + rubrichandler.CriticalPrefix + ` and only the CI is tested locally. ` + rubrichandler.NeverGuess,
		HandlerData: map[string]any{},
	}
}
