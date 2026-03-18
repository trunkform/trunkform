package perftestrubricbuilder

import "trunkform-mcp/rubricHandlers/rubrichandler"

func NewRubricHandler() rubrichandler.RubricHandler {
	return rubrichandler.RubricHandler{
		HandlerKey:  "perf-test-implemented",
		HandlerName: "Add Performance Testing",
		HandlerText: rubrichandler.InstructionPrefix + rubrichandler.AskTheUser(`PERFORMANCE TESTING: what specific service or application component should be benchmarked?`) + `
Then ` + rubrichandler.AskTheUser(`what language/framework should be used for the performance tests?`) + ` 
Then ` + rubrichandler.AskTheUser(`what are the performance requirements for successful operation (e.g., response
time, throughput, concurrent users)?`) + ` Only after all three questions have been answered should the agentic ai assistant
help the user build the performance testing tool. After the performance test passes, update .rubric.perf-test-implemented to the exact command used to run the passing performance test, then reread ./trunkform.json and call trunkform again with the full object. ` + rubrichandler.CriticalPrefix + ` and tested. ` + rubrichandler.NeverGuess,
		HandlerData: map[string]any{},
	}
}
