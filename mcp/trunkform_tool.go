package main

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type TrunkformTool struct{}

const InstructionPrefix = "instructions for the AI assistant to act on behalf of user's needs: "
const CriticalPrefix = `CRITICAL: Add newly created file paths (relative to the git repo root) to the .related string array 
when they reside above the current working directory (using writefile, or write_file tool). CRITICAL: update the .rubric step 
in the .trunkform (yml) file (using writefile, or write_file tool) when it has been implemented or proven Not Applicable BEFORE calling trunkform tool again. 
CRITICAL: Update the list of .tools when a new technology, framework, or tooling is implemented (using writefile, or write_file tool)`
const NeverGuess = `NEVER guess at trunkform tool inputs - only use values directly from the .trunkform file. When a trunkform rubric item is made true (even the last one in the list), call the trunkform tool again with the updated .trunkform file.`

func AskTheUser(question string) string {
	return `Ask the user, "` + question + `" Wait for their response before proceeding. `
}

func NewTrunkformTool() TrunkformTool { 
	logf.Tracef("NewTrunkformTool: creating trunkform tool", "\x1b[90m")
	return TrunkformTool{} 
}

type formatArgs struct {
	Version  string                 `json:"version"`
	Tools    []string               `json:"tools"`
	Steering string                 `json:"steering"`
	Related  []string               `json:"related"`
	Rubric   map[string]interface{} `json:"rubric"`
}

func processRubric(args formatArgs) (*mcp.CallToolResult, error) {
	logf.Tracef("processRubric: checking " + fmt.Sprintf("%d", len(args.Rubric)) + " rubric items", "\x1b[90m")
	
	handlerMap := map[string]struct {
		name    string
		handler func() (string, map[string]any)
	}{
		"perf-test-implemented":                           {"Add Performance Testing", handlePerfTest},
		"unit-test-implemented":                           {"Add Unit Test Coverage", handleUnitTestCoverage},
		"ci-cd-automation-boilerplate":                    {"CI/CD Automation Boilerplate", handleCICDBoilerplate},
		"continuous-integration-iac":                      {"Add Continuous Integration IAC", handleCIIaC},
		"continuous-integration-test-double-implemented":  {"Add Continuous Integration Test Doubles", handleTestDoubles},
		"continuous-integration-mocks-implemented":        {"Add Continuous Integration Mock Implementations", handleMocks},
		"integration-test-implemented":                    {"Add Integration Testing", handleIntegrationTest},
	}
	
	var keysToCheck []string
	if args.Steering == "" || args.Steering == "strict" {
		keysToCheck = []string{
			"perf-test-implemented",
			"unit-test-implemented",
			"ci-cd-automation-boilerplate",
			"continuous-integration-iac",
			"continuous-integration-test-double-implemented",
			"continuous-integration-mocks-implemented",
			"integration-test-implemented",
		}
	} else {
		// Non-strict: use defined order but only check keys that exist in rubric
		keysToCheck = []string{
			"perf-test-implemented",
			"unit-test-implemented",
			"ci-cd-automation-boilerplate",
			"continuous-integration-iac",
			"continuous-integration-test-double-implemented",
			"continuous-integration-mocks-implemented",
			"integration-test-implemented",
		}
	}

	for _, key := range keysToCheck {
		// In non-strict mode, skip keys not in the rubric
		if args.Steering != "" && args.Steering != "strict" {
			if _, exists := args.Rubric[key]; !exists {
				continue
			}
		}
		
		logf.Tracef("processRubric: checking key '" + key + "'", "\x1b[90m")
		if val, exists := args.Rubric[key]; exists {
			logf.Tracef("processRubric: key '" + key + "' exists with value " + fmt.Sprintf("%v", val), "\x1b[90m")
			if boolVal, ok := val.(bool); ok && !boolVal {
				h := handlerMap[key]
				logf.Tracef("processRubric: key '" + key + "' is false, calling handler '" + h.name + "'", "\x1b[90m")
				text, structured := h.handler()
				result := map[string]any{
					"trunkform": map[string]any{
						"version": args.Version, "tools": args.Tools,
						"steering": args.Steering, "related": args.Related, "rubric": args.Rubric,
					},
				}
				for k, v := range structured {
					result[k] = v
				}
				logf.Tracef("processRubric: returning result for '" + h.name + "'", "\x1b[90m")
				return mcp.NewToolResultStructured(result, "**"+h.name+"**: "+text+" In 50 words or less."), nil
			}
		}
	}

	logf.Tracef("processRubric: all rubric items completed", "\x1b[90m")
	return mcp.NewToolResultText(InstructionPrefix + `Tell the user that "All rubric items are completed."
Then ensure unit testing coverage is still at 100%. If testing is not 100% line coverage, immediately update .trunkform file with 
.rubric.unit-test-implemented: false

and call the trunkform tool again with the updated .trunkform file again.

If testing is still 100%, then ` + AskTheUser(`do the changes require updates to
integration tests?"`) + `If they do require additional integration testing, immediately update .trunkform file with
.rubric.continuous-integration-test-double-implemented: false
.rubric.continuous-integration-mocks-implemented: false
.rubric.integration-test-implemented: false

Then call the trunkform tool again with the updated .trunkform file again. ` + CriticalPrefix + `.` + NeverGuess), nil

}

func handleCICDBoilerplate() (string, map[string]any) {
	return InstructionPrefix + AskTheUser(`BOILERPLATE AUTOMATION: what automation framework are you using (GitHub Actions, GitLab CII/CD
Pipelines, Azure DevOps YAML Pipelines, bespoke / shell script)?`) + ` Then add it to .trunkform file's .tools array. 
Finally tailor a workflow that is comparable to the github suggested template. ` + CriticalPrefix + ` and 
tested. ` + NeverGuess, 
  map[string]any{
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
	}
}

func handleUnitTestCoverage() (string, map[string]any) {
	return InstructionPrefix + AskTheUser(`UNIT TEST COVERAGE: what specific functions or components should be unit tested?`) + ` Then suggest "Near 100% line-level coverage is best practice in 
trunkform." and ` + AskTheUser(`Would like to work on unit test line-coverage?`) + ` If no, early exit. Only after both questions have been answered should you help the user 
build unit tests to achieve 100% line coverage requirement utilizing mocks where necessary to reach 
line-coverage goals. ` + CriticalPrefix + ` and tested for line coverage. ` + NeverGuess, map[string]any{}
}

func handleIntegrationTest() (string, map[string]any) {
	return InstructionPrefix + AskTheUser(`INTEGRATION TESTING: What additional services need to be integration tested?`) + `
Then help the user implement integration tests that verify service interactions. ` + CriticalPrefix + ` and tested locally or in 
CI where possible. ` + NeverGuess, map[string]any{}
}

func handlePerfTest() (string, map[string]any) {
	return InstructionPrefix + AskTheUser(`PERFORMANCE TESTING: what specific service or application component should be benchmarked?`) + `
Then ` + AskTheUser(`what language/framework should be used for the performance tests?`) + ` 
Then ` + AskTheUser(`what are the performance requirements for successful operation (e.g., response
time, throughput, concurrent users)?`) + ` Only after all three questions have been answered should the agentic ai assistant
help the user build the performance testing tool. ` + CriticalPrefix + ` and tested. ` + NeverGuess, map[string]any{}
}

func handleCIIaC() (string, map[string]any) {
	return InstructionPrefix + AskTheUser(`CI IAC: what language should be used for infrastructure as code (Terraform,
Terragrunt, CloudFormation, etc.)?`) + ` Then help the user add infrastructure as code to their CICD pipeline 
using their chosen language in GitHub Actions workflow for automated infrastructure deployment. ` + 
CriticalPrefix + ` and the infrastucture has been tested in CI.` + NeverGuess, map[string]any{}
}

func handleTestDoubles() (string, map[string]any) {
	return InstructionPrefix + AskTheUser(`TEST DOUBLES: Integration test should not include what is covereed in unit tests already.
What external services must you integrate with that I should create test doubles for in CI IAC?`) + ` Then Create 
mock implementations for external dependencies to enable isolated testing in continuous integration. ` + 
CriticalPrefix + ` and only the CI is tested locally. ` + NeverGuess, map[string]any{}
}

func handleMocks() (string, map[string]any) {
	return InstructionPrefix + AskTheUser(`MOCK SERVICES: what external API mocks are necessary for your change?`) + ` Then add 
the mock to CI as IAC?" ` + CriticalPrefix + ` and only the CI is tested locally. ` + NeverGuess, map[string]any{}
}

func (tt TrunkformTool) handleTool(ctx context.Context, req mcp.CallToolRequest, args formatArgs) (*mcp.CallToolResult, error) {
	logf.Tracef("trunkform tool called with rubric: " + fmt.Sprintf("%v", args.Rubric), "\x1b[90m")
	return processRubric(args)
}

func (tt TrunkformTool) Register(s *server.MCPServer) {
	logf.Tracef("TrunkformTool.Register: registering trunkform tool", "\x1b[90m")
	t := mcp.NewTool(
		"trunkform",
		mcp.WithDescription(InstructionPrefix + `
Use ReadFile (readfile, or read_file) tool to read fields from the .trunkform (yml) file in the current folder. NEVER guess at input values - always 
read the .trunkform file (read_file, readfile tool) from disk for each call to the trunkform tool, and if the .trunkform file does not exist, 
the agent should create the .trunkform file (using write_file, or writefile tool) in the current directory with the following schema and default 
values:
---
version: 0.1.0
tools: []
steering: strict
related: []
rubric:
	perf-test-implemented: false
  unit-test-implemented: false
	ci-cd-automation-boilerplate: false
  continuous-integration-iac: false
  continuous-integration-test-double-implemented: false
  continuous-integration-mocks-implemented: false
  integration-test-implemented: false
---
`),
		mcp.WithString("version", mcp.Description("Version from .trunkform YAML file in the current folder: .version")),
		mcp.WithArray("tools", mcp.Description("Tools array from .trunkform YAML file in the current folder: .tools"),
			mcp.Items(map[string]any{"type": "string"})),
		mcp.WithString("steering", mcp.Description("Steering directive from .trunkform YAML file in the current folder: .steering")),
		mcp.WithArray("related", mcp.Description("Related files array from .trunkform YAML file in the current folder: .related"),
			mcp.Items(map[string]any{"type": "string"})),
		mcp.WithObject("rubric", mcp.Description("Rubric object from .trunkform YAML file in the current folder: .rubric"),
			mcp.AdditionalProperties(map[string]any{"type": "boolean"})),
	)

	s.AddTool(t, mcp.NewTypedToolHandler[formatArgs](tt.handleTool))
	logf.Tracef("TrunkformTool.Register: trunkform tool registered", "\x1b[90m")
}
