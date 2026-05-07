package cicdlintingbuilder

import "trunkform-mcp/rubricHandlers/rubrichandler"

func NewRubricHandler() rubrichandler.RubricHandler {
	return rubrichandler.RubricHandler{
		HandlerKey:  "ci-cd-linting",
		HandlerName: "CI/CD Linting Automation",
		HandlerText: rubrichandler.InstructionPrefix + `LINTING AUTOMATION: check the ./trunkform.json file's
.rubric.lint-implemented command exists in continuous integration, based on where it was implemented in
.rubric.ci-cd-automation-boilerplate, update .rubric.ci-cd-linting to the exact command used to grep for
the linting command in automation, then reread ./trunkform.json and call trunkform again with the full
object. ` + rubrichandler.CriticalPrefix + rubrichandler.NeverGuess,
		HandlerData: map[string]any{},
	}
}
