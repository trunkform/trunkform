package main

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	rubrichandlers "trunkform-mcp/rubricHandlers"
)

type TrunkformTool struct{}

const InstructionPrefix = rubrichandlers.InstructionPrefix
const CriticalPrefix = rubrichandlers.CriticalPrefix
const NeverGuess = rubrichandlers.NeverGuess
const ToolDescription = `
The trunkform tool is a software delivery checklisting tool, often used as a verb; "trunkform."
USAGE INSTRUCTIONS: 
When a user says "trunkform" it means to call this mcp tool, first reading the contents of the ./trunkform.json 
file on disk and passing it as the argument. If the ./trunkform.json file does not exist, create the ./trunkform.json 
file with the following schema and default values:

~~~
{
  "version": "0.1.0",
  "tools": [],
  "related": [],
  "rubric": {
	}
}
~~~
`

func AskTheUser(question string) string {
	return rubrichandlers.AskTheUser(question)
}

func NewTrunkformTool() TrunkformTool {
	logf.Tracef("NewTrunkformTool: creating trunkform tool", "\x1b[90m")
	return TrunkformTool{}
}

type trunkformSchema struct {
	Version string             `json:"version"`
	Tools   []string           `json:"tools"`
	Related []string           `json:"related"`
	Rubric  map[string]*string `json:"rubric"`
}

func processRubric(args trunkformSchema) (*mcp.CallToolResult, error) {
	logf.Tracef("processRubric: checking "+fmt.Sprintf("%d", len(args.Rubric))+" rubric items", "\x1b[90m")

	handlers := rubrichandlers.DefaultRubricHandlers()

	for _, handler := range handlers {
		key := handler.Key()
		logf.Tracef("processRubric: checking key '"+key+"'", "\x1b[90m")
		val, exists := args.Rubric[key]
		if exists {
			logf.Tracef("processRubric: key '"+key+"' exists with value "+fmt.Sprintf("%v", val), "\x1b[90m")
			if val != nil {
				continue
			}
		}

		logf.Tracef("processRubric: key '"+key+"' is missing or incomplete, calling handler '"+handler.Name()+"'", "\x1b[90m")
		resultData := handler.Build()
		result := map[string]any{
			"name":         handler.Name(),
			"instructions": resultData.Text,
			"trunkform": map[string]any{
				"version": args.Version, "tools": args.Tools,
				"related": args.Related, "rubric": args.Rubric,
			},
		}
		for k, v := range resultData.Structured {
			result[k] = v
		}
		logf.Tracef("processRubric: returning result for '"+handler.Name()+"'", "\x1b[90m")
		return mcp.NewToolResultStructuredOnly(result), nil
	}

	logf.Tracef("processRubric: all rubric items completed", "\x1b[90m")
	return mcp.NewToolResultStructuredOnly(map[string]any{"instructions": InstructionPrefix + `Inform the user the Rubric is complete but needs to be tested. ` + AskTheUser(`do the changes
require updates to integration tests?`) + `If they do require additional integration testing, immediately update ./trunkform.json file so it sets
.rubric.continuous-integration-test-double-implemented: null, .rubric.continuous-integration-mocks-implemented: null, .rubric.integration-test-implemented: null,
read ./trunkform.json, and call the trunkform tool again with the entire updated trunkform json object.

If they do not require additional integration testing, then execute every command stored in ./trunkform.json .rubric for the keys in this exact order:
1. perf-test-implemented
2. unit-test-implemented
3. ci-cd-automation-boilerplate
4. continuous-integration-iac
5. continuous-integration-test-double-implemented
6. continuous-integration-mocks-implemented
7. integration-test-implemented
If any command fails, ` + AskTheUser(`would you like suggestions on how to remediate the failing test before the rubric item is marked null?`) + `If the user says no, immediately update that rubric item to null in ./trunkform.json, read ./trunkform.json, and call the trunkform tool again with the entire updated trunkform object.

If unit-testcoverage is 100%, has no measurable coverage ("Unknown%" means "no measureable coverage"), then tell the user that "all rubric items tested and complete. yolo", 
otherwise` + AskTheUser(`is less than 100% line coverage acceptable for these changes, or is the test command missing a coverage flag?`) + `If the user says no, immediately update ./trunkform.json file so it contains 
.rubric.unit-test-implemented: null, read ./trunkform.json, and call the trunkform tool again with the entire updated trunkform json object. ` + CriticalPrefix + ` ` + NeverGuess}), nil
}

func (tt TrunkformTool) handleTool(ctx context.Context, req mcp.CallToolRequest, args trunkformSchema) (*mcp.CallToolResult, error) {
	if args.Version == "" && len(args.Tools) == 0 && len(args.Related) == 0 && len(args.Rubric) == 0 {
		return mcp.NewToolResultStructuredOnly(map[string]any{"instructions": InstructionPrefix + ToolDescription + NeverGuess}), nil
	}
	logf.Tracef("trunkform tool called with rubric: "+fmt.Sprintf("%v", args.Rubric), "\x1b[90m")
	return processRubric(args)
}

func (tt TrunkformTool) Register(s *server.MCPServer) {
	logf.Tracef("TrunkformTool.Register: registering trunkform tool", "\x1b[90m")
	t := mcp.NewTool(
		"trunkform",
		mcp.WithDescription(ToolDescription),
		mcp.WithString("version", mcp.Description("Version from ./trunkform.json file in the current folder at: .version")),
		mcp.WithArray("tools", mcp.Description("Tools array from ./trunkform.json file in the current folder at: .tools"),
			mcp.Items(map[string]any{"type": "string"})),
		mcp.WithArray("related", mcp.Description("Related files array from ./trunkform.json file in the current folder at: .related"),
			mcp.Items(map[string]any{"type": "string"})),
		mcp.WithObject("rubric", mcp.Description("Rubric object from ./trunkform.json file in the current folder at: .rubric"),
			mcp.AdditionalProperties(map[string]any{"type": []string{"string", "null"}})),
	)

	s.AddTool(t, mcp.NewTypedToolHandler[trunkformSchema](tt.handleTool))
	logf.Tracef("TrunkformTool.Register: trunkform tool registered", "\x1b[90m")
}
