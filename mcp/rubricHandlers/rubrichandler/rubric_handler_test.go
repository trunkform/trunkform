package rubrichandler

import "testing"

func TestConstants(t *testing.T) {
	expectedPrefix := "INSTRUCTIONS FOR THE AI ASSISTANT TO ACT ON BEHALF OF USER'S NEEDS: "
	if InstructionPrefix != expectedPrefix {
		t.Errorf("InstructionPrefix = %q, want %q", InstructionPrefix, expectedPrefix)
	}

	expectedCritical := "CRITICAL: Add newly created file paths (relative to the git repo root) to the .related string array \nwhen they reside above the current working directory (using writefile, or write_file tool). CRITICAL: update the .rubric step \nin the ./trunkform.json file (using writefile, or write_file tool) when it has been implemented or proven Not Applicable BEFORE calling trunkform tool again. After editing ./trunkform.json, reread the file and pass the entire trunkform object on the next trunkform tool call.\nRubric items are nullable strings: use null when incomplete, and use a non-null string to record the command used to verify the rubric item. For Not Applicable (N/A), record the value as: cat <<< \"N/A: <reason>\". \nCRITICAL: Update the list of .tools when a new technology, framework, or tooling is implemented (using writefile, or write_file tool), then reread ./trunkform.json and pass the entire trunkform object."
	if CriticalPrefix != expectedCritical {
		t.Errorf("CriticalPrefix = %q, want %q", CriticalPrefix, expectedCritical)
	}

	expectedNever := `NEVER guess at trunkform tool inputs - only use values directly from the ./trunkform.json file. When ./trunkform.json changes, reread it and call the trunkform tool again with the entire updated trunkform object.`
	if NeverGuess != expectedNever {
		t.Errorf("NeverGuess = %q, want %q", NeverGuess, expectedNever)
	}
}

func TestAskTheUser(t *testing.T) {
	result := AskTheUser("what is your name?")
	expected := `Ask the user, "what is your name?" Wait for their response before proceeding. `
	if result != expected {
		t.Errorf("AskTheUser() = %q, want %q", result, expected)
	}
}

func TestRubricHandlerBuild(t *testing.T) {
	handler := RubricHandler{
		HandlerKey:  "test-key",
		HandlerName: "test-name",
		HandlerText: "test-text",
		HandlerData: map[string]any{"k": "v"},
	}

	if handler.Key() != "test-key" {
		t.Fatalf("Key() = %q, want %q", handler.Key(), "test-key")
	}
	if handler.Name() != "test-name" {
		t.Fatalf("Name() = %q, want %q", handler.Name(), "test-name")
	}

	result := handler.Build()
	if result.Text != "test-text" {
		t.Fatalf("Build().Text = %q, want %q", result.Text, "test-text")
	}
	if result.Structured["k"] != "v" {
		t.Fatalf("Build().Structured[k] = %v, want %q", result.Structured["k"], "v")
	}
}
