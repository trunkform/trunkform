package cicdunittestingbuilder

import "github.com/trunkform/trunkform/mcp/rubricHandlers/rubrichandler"

func NewRubricHandler() rubrichandler.RubricHandler {
	return rubrichandler.RubricHandler{
		HandlerKey:  "ci-cd-unit-testing",
		HandlerName: "CI/CD Unit Testing Automation",
		HandlerText: rubrichandler.InstructionPrefix + `UNIT TESTING AUTOMATION: check the ./trunkform.json file's
.rubric.unit-test-implemented command exists in continuous integration, based on where it was implemented in
.rubric.ci-cd-automation-boilerplate, update .rubric.ci-cd-unit-testing to the exact command used to grep for
the unit testing command in automation, then reread ./trunkform.json and call trunkform again with the full
object. ` + rubrichandler.CriticalPrefix + rubrichandler.NeverGuess,
		HandlerData: map[string]any{},
	}
}
