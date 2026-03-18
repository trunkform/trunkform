# Code Review - Actionable Tasks Not Covered by CHANGELOG.md

## Major Refactoring (Not Documented)

### Extract rubricHandlers Package
- [x] Add `mcp/rubricHandlers/` package to git with `git add mcp/rubricHandlers/`
- [x] Document the new package structure in `mcp/README.md` (7 handler subdirectories + registry)
- [x] Add godoc comments to `rubricHandlers/registry.go` explaining the handler registration pattern
- [x] Document the `RubricHandler` interface in `rubricHandlers/rubrichandler/rubric_handler.go`

### Schema Changes (Breaking)
- [x] Update CHANGELOG.md to document removal of `steering` field from schema
- [x] Update CHANGELOG.md to document change from YAML (`.trunkform`) to JSON (`./trunkform.json`)
- [x] Update CHANGELOG.md to document nested `trunkform` argument structure (breaking API change)
- [x] Add migration script or instructions for users with existing `.trunkform` YAML files

## New Files to Track

### Configuration Files
- [ ] Add `mcp/trunkform.json` to git: `git add mcp/trunkform.json`
- [ ] Add `cfn/trunkform.json` to git: `git add cfn/trunkform.json`
- [ ] Add `trunkform.ai/trunkform.json` to git: `git add trunkform.ai/trunkform.json`

### Infrastructure
- [ ] Add `mcp/infra/localstack/` to git: `git add mcp/infra/localstack/`
- [ ] Add `mcp/infra/modules/mcp/providers.tf` to git: `git add mcp/infra/modules/mcp/providers.tf`
- [ ] Add `mcp/infra/modules/persistent-resources/` to git: `git add mcp/infra/modules/persistent-resources/`
- [ ] Add `mcp/infra/persistent-resources-ci/` to git: `git add mcp/infra/persistent-resources-ci/`

### Renamed/Moved Files
- [x] Update CHANGELOG.md to document `mcp/infra/prod/` renamed to `mcp/infra/cd/`
- [x] Update CHANGELOG.md to document `mcp/infra/remote_state.hcl` moved to repo root
- [x] Update any documentation referencing the old `prod` directory name
- [x] Verify all terragrunt references use the new paths

## Test Coverage Gaps

### Performance Test Updates
- [x] Document in CHANGELOG.md that `trunkform_bench_test.go` was updated for new `trunkform` wrapper
- [x] Add assertions for actual performance metrics (N/A - local test, production metrics TBD)
- [x] Document expected performance baseline in test comments (N/A - baseline varies by infrastructure)
- [x] Consider adding performance regression detection (N/A - will spec based on production needs)

### Handler Tests
- [x] Verify all 7 rubric handlers in `rubricHandlers/` have corresponding tests
- [x] Add integration test that exercises the full handler registry
- [x] Test handler ordering and dependencies

## Makefile Changes Not in Changelog

### Target Changes
- [x] Document Makefile changes: new `message-install` and `start` targets, `all` now prompts instead of auto-installing

### Port Management
- [x] Document hardcoded port 8081 for tests in CHANGELOG.md
- [x] Document why tests use 8081 instead of 8080 (avoid conflicts with running server)

## Code Quality Issues

### Whitespace Changes
- [x] Run `gofmt` to ensure consistent formatting across all Go files

### Magic Numbers
- [x] Replace hardcoded `keysToCheck` array with `DefaultRubricHandlers()` list

### Error Handling
- [x] Add error handling for `json.Marshal` in `trunkform_bench_test.go` (N/A - safe to ignore in load test)
- [x] Add error handling for `http.NewRequest` in `trunkform_bench_test.go` (N/A - safe to ignore in load test)

## Cleanup Tasks

### Ignore Files
- [x] Add `cfn/coverage/` to `.gitignore`
- [x] Add `mcp/.gocache/` to `.gitignore`
- [x] Verify `coverage.out` is already in `.gitignore`

## Testing Before Commit

### Integration Tests
- [x] Test with copilot to ensure MCP integration still works
- [x] Verify `trunkform.json` files are read correctly by the tool

### Infrastructure Tests
- [ ] Run `terragrunt plan` in all infrastructure directories
- [ ] Verify localstack setup works for local testing
