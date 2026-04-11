// Package rubrichandlers provides the registry and factory for trunkform rubric handlers.
package rubrichandlers

import (
	"trunkform-mcp/rubricHandlers/rubrichandler"
	lintrubricbuilder "trunkform-mcp/rubricHandlers/lintrubricbuilder"
	unittestrubricbuilder "trunkform-mcp/rubricHandlers/unittestrubricbuilder"
	cicdboilerplaterubricbuilder "trunkform-mcp/rubricHandlers/cicdboilerplaterubricbuilder"
	ciiacrubricbuilder "trunkform-mcp/rubricHandlers/ciiacrubricbuilder"
	testdoublesrubricbuilder "trunkform-mcp/rubricHandlers/testdoublesrubricbuilder"
	mocksrubricbuilder "trunkform-mcp/rubricHandlers/mocksrubricbuilder"
	integrationtestrubricbuilder "trunkform-mcp/rubricHandlers/integrationtestrubricbuilder"
	perftestrubricbuilder "trunkform-mcp/rubricHandlers/perftestrubricbuilder"
	steeringrubricbuilder "trunkform-mcp/rubricHandlers/steeringrubricbuilder"
)

const InstructionPrefix = rubrichandler.InstructionPrefix
const CriticalPrefix = rubrichandler.CriticalPrefix
const NeverGuess = rubrichandler.NeverGuess

func AskTheUser(question string) string {
	return rubrichandler.AskTheUser(question)
}

// DefaultRubricHandlers returns the ordered list of rubric handlers for the trunkform workflow.
func DefaultRubricHandlers() []rubrichandler.RubricHandler {
	return []rubrichandler.RubricHandler{
		steeringrubricbuilder.NewRubricHandler(),
		lintrubricbuilder.NewRubricHandler(),
		unittestrubricbuilder.NewRubricHandler(),
		cicdboilerplaterubricbuilder.NewRubricHandler(),
		ciiacrubricbuilder.NewRubricHandler(),
		testdoublesrubricbuilder.NewRubricHandler(),
		mocksrubricbuilder.NewRubricHandler(),
		integrationtestrubricbuilder.NewRubricHandler(),
		perftestrubricbuilder.NewRubricHandler(),
	}
}
