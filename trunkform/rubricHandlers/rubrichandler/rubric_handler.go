// Package rubrichandler defines the interface and types for trunkform rubric handlers.
package rubrichandler

const InstructionPrefix = "INSTRUCTIONS FOR THE AI ASSISTANT TO ACT ON BEHALF OF USER'S NEEDS: "
const CriticalPrefix = `CRITICAL: Add newly created file paths (relative to the git repo root) to the .related string array 
when they reside above the current working directory (using writefile, or write_file tool). CRITICAL: update the .rubric step 
in the ./trunkform.json file (using writefile, or write_file tool) when it has been implemented or proven Not Applicable BEFORE calling trunkform tool again. After editing ./trunkform.json, reread the file and pass the entire trunkform object on the next trunkform tool call.
Rubric items are nullable strings: use null when incomplete, and use a non-null string to record the command used to verify the rubric item. For Not Applicable (N/A), record the value as: cat <<< "N/A: <reason>". 
CRITICAL: Update the list of .tools when a new technology, framework, or tooling is implemented (using writefile, or write_file tool), then reread ./trunkform.json and pass the entire trunkform object.`
const NeverGuess = `NEVER guess at trunkform tool inputs - only use values directly from the ./trunkform.json file. When ./trunkform.json changes, reread it and call the trunkform tool again with the entire updated trunkform object.`

// Result contains the text and structured data returned by a rubric handler.
type Result struct {
	Text       string
	Structured map[string]any
}

// RubricHandler implements a single rubric checklist item handler.
type RubricHandler struct {
	HandlerKey  string
	HandlerName string
	HandlerText string
	HandlerData map[string]any
}

func (h RubricHandler) Key() string {
	return h.HandlerKey
}

func (h RubricHandler) Name() string {
	return h.HandlerName
}

func (h RubricHandler) Build() Result {
	return Result{
		Text:       h.HandlerText,
		Structured: h.HandlerData,
	}
}

func AskTheUser(question string) string {
	return `Ask the user, "` + question + `" Wait for their response before proceeding. `
}
