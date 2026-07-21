package lintrubricbuilder

import "github.com/trunkform/trunkform/mcp/rubricHandlers/rubrichandler"

func NewRubricHandler() rubrichandler.RubricHandler {
	return rubrichandler.RubricHandler{
		HandlerKey:  "lint-implemented",
		HandlerName: "Add Linting",
		HandlerText: rubrichandler.InstructionPrefix + rubrichandler.AskTheUser(`LINTING: is linting configured for this project?`) + ` If no, help the user set up a linter appropriate for the language/framework in .tools. ` + rubrichandler.AskTheUser(`What is the command to run the linter?`) + ` Run the lint command and fix any issues. After linting passes, update .rubric.lint-implemented to the exact command used to run the passing lint check, then reread ./trunkform.json and call trunkform again with the full object. ` + rubrichandler.CriticalPrefix + ` ` + rubrichandler.NeverGuess,
		HandlerData: map[string]any{},
	}
}
