package ciiacrubricbuilder

import "trunkform-mcp/rubricHandlers/rubrichandler"

func NewRubricHandler() rubrichandler.RubricHandler {
	return rubrichandler.RubricHandler{
		HandlerKey:  "continuous-integration-iac",
		HandlerName: "Add Continuous Integration IAC",
		HandlerText: rubrichandler.InstructionPrefix + rubrichandler.AskTheUser(`CI IAC: what language should be used for infrastructure as code (Terraform,
Terragrunt, CloudFormation, etc.)?`) + ` Then help the user add infrastructure as code to their CICD pipeline 
using their chosen language in GitHub Actions workflow for automated infrastructure deployment. ` +
			` After the CI IaC verification passes, update .rubric.continuous-integration-iac to the exact command used to verify the passing CI IaC configuration, then reread ./trunkform.json and call trunkform again with the full object. ` +
			rubrichandler.CriticalPrefix + ` and the infrastucture has been tested in CI.` + rubrichandler.NeverGuess,
		HandlerData: map[string]any{},
	}
}
