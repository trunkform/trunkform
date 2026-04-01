# Changelog

## 0.1.1-rc1 2026-03-31

### Added

- added a release binary action which triggers on new tags.
- cleaned up makefile for deterministic releases
- added tests which cover the order of operations by which the tool should test the rubric items.

### Changed

- fixed the order that the mcp server returns when you're ready to test each rubric item.

## 2026-03-11

### Changed

🔴 Migration steps:

Convert YAML to JSON and replace rubric booleans:

```bash
yq -o=json \
  'del(.steering) | 
  .rubric |= with_entries(
    select(.value == false) .value = null | 
    select(.value == true) .value = "exit 1"
  )' .trunkform > trunkform.json
rm .trunkform
```

- Renamed `mcp/infra/prod/` to `mcp/infra/cd/`
- Moved `mcp/infra/remote_state.hcl` to repo root `remote_state.hcl`
- Updated `trunkform_bench_test.go` for new `trunkform` wrapper
- Makefile: new `message-install` and `start` targets, `all` now prompts instead of auto-installing
  - Makefile `make test` uses port 8081 to avoid conflicts with running server on 8080
- Linted with `gofmt` (backlogged linting rubric handler)
- [`trunkform_tool.go`](trunkform_tool.go) accepts trunkform.json content directly as individual parameters (removed formatArgs wrapper)
- Guidance strings in [`trunkform_tool.go`](/Users/Rich/Documents/source/trunkform/trunkform/mcp/trunkform_tool.go) were updated to say:
  - edit `./trunkform.json`
  - reread the file
  - call the tool again with the full updated object
- Rubric values were changed from booleans to nullable strings.
  - `null` means incomplete
  - non-null string means complete
  - N/A wording now says: `For Not Applicable (N/A), record a non-null string which echo's the reason given why it is N/A.`
- [`trunkformSchema`](/Users/Rich/Documents/source/trunkform/trunkform/mcp/trunkform_tool.go) now uses `map[string]*string` for `Rubric`.
- [`processRubric`](/Users/Rich/Documents/source/trunkform/trunkform/mcp/trunkform_tool.go) now treats:
  - missing or `null` rubric entry as incomplete
  - non-null rubric entry as complete
- Steering was removed from the trunkformSchema since the tool's internal logic now fully handles the rubric state and completion path. Will implement custom rubric steering logic handling later.
- JSON schema for rubric additional properties was updated from `boolean` to `["string", "null"]`.
- Completion/reset instructions in the tool text were updated from `false` to `null`.
- `handlePerfTest` was updated first to tell the AI to store the exact passing test command in `.rubric.perf-test-implemented`.
- Then the rest of the rubric handlers were updated to include the same “store the exact command used” language:
  - `handleCICDBoilerplate`
  - `handleUnitTestCoverage`
  - `handleIntegrationTest`
  - `handleCIIaC`
  - `handleTestDoubles`
  - `handleMocks`
- The “all complete” completion-path language was changed so it now:
  - instructs execution of the commands stored in `./trunkform.json .rubric` in `keysToCheck` order
  - then uses the final completion phrase `all rubric items tested and complete. yolo.`
  - and the execution-intro text now says:
    `Inform the user the Rubric is complete but needs to be tested, then execute every command stored in ./trunkform.json .rubric for the keys in this exact order`
  - and if one of those commands fails, the AI must ask the user whether they want remediation suggestions before the rubric item is marked `null`
- [`trunkform_tool_test.go`](/Users/Rich/Documents/source/trunkform/trunkform/mcp/trunkform_tool_test.go) was updated throughout to match the nullable-string rubric model and the revised language.
- `TestHandlers` now asserts each handler includes the “update `.rubric...` to the exact command used” wording.
- The completion-path test assertion was updated to match the new execution-intro wording.
- [`Makefile`](/Users/Rich/Documents/source/trunkform/trunkform/mcp/Makefile) integration payload was updated earlier to use nested `trunkform` arguments, and perf/integration invocations now use `null` for incomplete rubric entries where relevant.
- [`trunkform_bench_test.go`](/Users/Rich/Documents/source/trunkform/trunkform/mcp/trunkform_bench_test.go) was updated to send `null` for incomplete `perf-test-implemented`.
- [`trunkform.json`](/Users/Rich/Documents/source/trunkform/trunkform/mcp/trunkform.json) currently stores string rubric commands, not booleans.
- all rubric entries in `trunkform.json` were updated to store the exact command used to complete that rubric item, instead of just `true`.
- One stored rubric command was corrected:
  - `continuous-integration-iac` is now `cat ../.github/workflows/mcp.yml | grep 'terragrunt run --all apply'`
- [`Makefile`](/Users/Rich/Documents/source/trunkform/trunkform/mcp/Makefile) unit-test output was cleaned up to stop printing the raw statement-coverage line and instead report the enforced function-level coverage result directly.
- The completion-path coverage prompt in [`trunkform_tool.go`](/Users/Rich/Documents/source/trunkform/trunkform/mcp/trunkform_tool.go) was revised to ask:
  - `is less than 100% line coverage acceptable for these changes, or is the test command missing a coverage flag?`
- The completion path in [`trunkform_tool.go`](/Users/Rich/Documents/source/trunkform/trunkform/mcp/trunkform_tool.go) now asks whether integration tests need updates before instructing the AI to run all rubric verification commands.
- An existing malformed quote in the completion-path integration-test prompt was fixed in [`trunkform_tool.go`](/Users/Rich/Documents/source/trunkform/trunkform/mcp/trunkform_tool.go).
- [`trunkform_tool_test.go`](/Users/Rich/Documents/source/trunkform/trunkform/mcp/trunkform_tool_test.go) now asserts the revised coverage prompt in the `all-complete` path.
- [`trunkform_tool_test.go`](/Users/Rich/Documents/source/trunkform/trunkform/mcp/trunkform_tool_test.go) now asserts that the integration-test update question appears before rubric command execution in the `all-complete` path.
- [`trunkform_tool_test.go`](/Users/Rich/Documents/source/trunkform/trunkform/mcp/trunkform_tool_test.go) also checks the new measurable-coverage wording in the completion path.

### Added

- new copilot test workflow in `.github/workflows/copilot.yml` which:
  - starts the MCP server
  - sets up copilot config to point to the MCP server
  - runs copilot with the trunkform tool and checks for a rocket emoji in the output log to confirm it ran successfully against the MCP server
