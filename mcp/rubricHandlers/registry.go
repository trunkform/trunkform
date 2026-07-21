// Package rubrichandlers provides the registry and factory for trunkform rubric handlers.
package rubrichandlers

import (
	"github.com/trunkform/trunkform/mcp/rubricHandlers/rubrichandler"
	lintrubricbuilder "github.com/trunkform/trunkform/mcp/rubricHandlers/lintrubricbuilder"
	unittestrubricbuilder "github.com/trunkform/trunkform/mcp/rubricHandlers/unittestrubricbuilder"
	cicdboilerplaterubricbuilder "github.com/trunkform/trunkform/mcp/rubricHandlers/cicdboilerplaterubricbuilder"
	cicdlintingbuilder "github.com/trunkform/trunkform/mcp/rubricHandlers/cicdlintingbuilder"
	cicdunittestingbuilder "github.com/trunkform/trunkform/mcp/rubricHandlers/cicdunittestingbuilder"
	ciiacrubricbuilder "github.com/trunkform/trunkform/mcp/rubricHandlers/ciiacrubricbuilder"
	testdoublesrubricbuilder "github.com/trunkform/trunkform/mcp/rubricHandlers/testdoublesrubricbuilder"
	mocksrubricbuilder "github.com/trunkform/trunkform/mcp/rubricHandlers/mocksrubricbuilder"
	integrationtestrubricbuilder "github.com/trunkform/trunkform/mcp/rubricHandlers/integrationtestrubricbuilder"
	perftestrubricbuilder "github.com/trunkform/trunkform/mcp/rubricHandlers/perftestrubricbuilder"
	steeringrubricbuilder "github.com/trunkform/trunkform/mcp/rubricHandlers/steeringrubricbuilder"
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
		cicdlintingbuilder.NewRubricHandler(),
		cicdunittestingbuilder.NewRubricHandler(),
		ciiacrubricbuilder.NewRubricHandler(),
		testdoublesrubricbuilder.NewRubricHandler(),
		mocksrubricbuilder.NewRubricHandler(),
		integrationtestrubricbuilder.NewRubricHandler(),
		perftestrubricbuilder.NewRubricHandler(),
	}
}
