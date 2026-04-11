package steeringrubricbuilder

import "trunkform-mcp/rubricHandlers/rubrichandler"

func NewRubricHandler() rubrichandler.RubricHandler {
	return rubrichandler.RubricHandler{
		HandlerKey:  "steering-implemented",
		HandlerName: "Add Steering",
		HandlerText: rubrichandler.InstructionPrefix + `
Detect which AI assistant is active and ensure the relevant steering file exists with the specified content to steer the AI assistant towards using the trunkform tool for software delivery tasks:

- **Gemini**: append to ` + "`./GEMINI.md`" + `
- **All Other Agents** (Copilot, Codex, Kiro, etc.): append to ` + "`./AGENTS.md`" + `

` + "```md" + `
# TRUNKFORM

## MCP

Trunkform is a software delivery checklisting tool, often used as a verb; "trunkform." Call the trunkform MCP tool before all software delivery steps, including ` + "`git commit`, `git push`." + `

# TEST EXECUTION

When a trunkform.json file is in the current working directory; Running tests means explicitly executing every command in the .rubric section of ./trunkform.json directly in the shell.
` + "```" + `

The command to validate steering is configured is ` + "`cat ./AGENTS.md | grep -i '# TRUNKFORM'` or `cat ./GEMINI.md | grep -i '# TRUNKFORM'`" + `. After steering passes, update .rubric.steering-implemented to the exact command used to run the passing steering check, then reread ./trunkform.json and call trunkform again with the full object. ` + rubrichandler.CriticalPrefix + ` ` + rubrichandler.NeverGuess,
		HandlerData: map[string]any{},
	}
}
