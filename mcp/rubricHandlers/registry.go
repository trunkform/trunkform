// Package rubrichandlers provides the registry and factory for trunkform rubric handlers.
package rubrichandlers

import (
	cicdboilerplaterubricbuilder "trunkform-mcp/rubricHandlers/cicdboilerplaterubricbuilder"
	ciiacrubricbuilder "trunkform-mcp/rubricHandlers/ciiacrubricbuilder"
	integrationtestrubricbuilder "trunkform-mcp/rubricHandlers/integrationtestrubricbuilder"
	mocksrubricbuilder "trunkform-mcp/rubricHandlers/mocksrubricbuilder"
	perftestrubricbuilder "trunkform-mcp/rubricHandlers/perftestrubricbuilder"
	"trunkform-mcp/rubricHandlers/rubrichandler"
	testdoublesrubricbuilder "trunkform-mcp/rubricHandlers/testdoublesrubricbuilder"
	unittestrubricbuilder "trunkform-mcp/rubricHandlers/unittestrubricbuilder"
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
		unittestrubricbuilder.NewRubricHandler(),
		cicdboilerplaterubricbuilder.NewRubricHandler(),
		ciiacrubricbuilder.NewRubricHandler(),
		testdoublesrubricbuilder.NewRubricHandler(),
		mocksrubricbuilder.NewRubricHandler(),
		integrationtestrubricbuilder.NewRubricHandler(),
		perftestrubricbuilder.NewRubricHandler(),
	}
}
