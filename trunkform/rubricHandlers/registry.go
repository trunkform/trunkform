// Package rubrichandlers provides the registry and factory for trunkform rubric handlers.
package rubrichandlers

import (
	"github.com/trunkform/trunkform/trunkform/rubricHandlers/rubrichandler"
	lintrubricbuilder "github.com/trunkform/trunkform/trunkform/rubricHandlers/lintrubricbuilder"
	unittestrubricbuilder "github.com/trunkform/trunkform/trunkform/rubricHandlers/unittestrubricbuilder"
	cicdboilerplaterubricbuilder "github.com/trunkform/trunkform/trunkform/rubricHandlers/cicdboilerplaterubricbuilder"
	cicdlintingbuilder "github.com/trunkform/trunkform/trunkform/rubricHandlers/cicdlintingbuilder"
	cicdunittestingbuilder "github.com/trunkform/trunkform/trunkform/rubricHandlers/cicdunittestingbuilder"
	ciiacrubricbuilder "github.com/trunkform/trunkform/trunkform/rubricHandlers/ciiacrubricbuilder"
	testdoublesrubricbuilder "github.com/trunkform/trunkform/trunkform/rubricHandlers/testdoublesrubricbuilder"
	mocksrubricbuilder "github.com/trunkform/trunkform/trunkform/rubricHandlers/mocksrubricbuilder"
	integrationtestrubricbuilder "github.com/trunkform/trunkform/trunkform/rubricHandlers/integrationtestrubricbuilder"
	perftestrubricbuilder "github.com/trunkform/trunkform/trunkform/rubricHandlers/perftestrubricbuilder"
	steeringrubricbuilder "github.com/trunkform/trunkform/trunkform/rubricHandlers/steeringrubricbuilder"
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
