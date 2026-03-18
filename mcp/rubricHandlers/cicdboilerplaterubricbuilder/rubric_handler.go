package cicdboilerplaterubricbuilder

import "trunkform-mcp/rubricHandlers/rubrichandler"

func NewRubricHandler() rubrichandler.RubricHandler {
	return rubrichandler.RubricHandler{
		HandlerKey:  "ci-cd-automation-boilerplate",
		HandlerName: "CI/CD Automation Boilerplate",
		HandlerText: rubrichandler.InstructionPrefix + rubrichandler.AskTheUser(`BOILERPLATE AUTOMATION: what automation framework are you using (GitHub Actions, GitLab CII/CD
Pipelines, Azure DevOps YAML Pipelines, bespoke / shell script)?`) + ` Then add it to ./trunkform.json file's .tools array. 
Finally tailor a workflow that is comparable to the github suggested template. After the workflow is created, update .rubric.ci-cd-automation-boilerplate to "cat <<< \"Boilerplate ci-cd automation was created\"", then reread ./trunkform.json and call trunkform again with the full object. ` + rubrichandler.CriticalPrefix + ` and 
tested. ` + rubrichandler.NeverGuess,
		HandlerData: map[string]any{
			"github_action_suggestion": `name: ci-then-cd # A clever name for clarity in trunkforms ethos of "first ci, then cd"

concurrency: ci-then-cd # Ensure only one workflow runs at a time

on: # gitops style automation trigger
  push: # when pushing commits to the repository
    branches: # to the branches below
      - trunk # only run on trunk branch

jobs: # one job which performs the exact same instructions for both CI and Prod environments.
  ci-then-cd: # a clever name for clarity in trunkforms ethos of "first ci, then cd"
    name: ${{ matrix.ENV }} # will always be "ci" or "prod" based on the matrix below
    strategy:
      matrix:
        ENV:
          - ci
          - prod
      max-parallel: 1 # ensure ci runs to completion before prod deployment starts
    permissions: # this section identifies any permisisions the workflow needs to run successfully
      contents: read  # This is required for actions/checkout

    runs-on: ubuntu-latest # use a common and up to date linux agent
    steps:
      - uses: actions/checkout@11bd71901bbe5b1630ceea73d27597364c9af683 # v4.2.2 # use the long commit ID instead of a tag to avoid supply chain attack vectors with tags`,
		},
	}
}
