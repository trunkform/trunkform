.PHONY: all deps build test local-int-test-http local-perf-test install kill

PORT ?= 8080
MCP  ?= http://127.0.0.1:$(PORT)/mcp
MAKE_LOGLEVEL ?= 4

all: deps build test local-int-test-http local-perf-test install

kill:
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== KILL ===\033[0m"; fi
	@pkill -f trunkform-mcp || true
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then pkill -f trunkform-mcp > /dev/null 2>&1 || echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mNo trunkform-mcp processes found\033[0m"; fi

deps:
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== DEPENDENCIES ===\033[0m"; fi
	@go get github.com/mark3labs/mcp-go@latest
	@go mod tidy
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Dependencies updated\033[0m"; fi
	@echo

build:
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== BUILD ===\033[0m"; fi
	@mkdir -p bin
	@GOOS=linux GOARCH=amd64 go build -o bin/trunkform-mcp-linux-amd64 .
	@GOOS=linux GOARCH=arm64 go build -o bin/trunkform-mcp-linux-arm64 .
	@GOOS=darwin GOARCH=amd64 go build -o bin/trunkform-mcp-darwin-amd64 .
	@GOOS=darwin GOARCH=arm64 go build -o bin/trunkform-mcp-darwin-arm64 .
	@GOOS=windows GOARCH=amd64 go build -o bin/trunkform-mcp-windows-amd64.exe .
	@GOOS=windows GOARCH=arm64 go build -o bin/trunkform-mcp-windows-arm64.exe .
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Binaries built for all platforms\033[0m"; fi
	@echo

test:
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== UNIT TESTS ===\033[0m"; fi
	@go test -skip 'TestTrunkformLoadTest' -coverprofile=coverage.out -v 2>&1 | tee /tmp/test_output.txt | sed \
	  -e '/^=== RUN/d' \
	  -e '/^=== PAUSE/d' \
	  -e '/^=== CONT/d' \
	  -e '/^exit status 1$$/d' \
	  -e 's/^\(#.*\)/\x1b[90m\1\x1b[0m/' \
	  -e 's/^\(.*\.go:[0-9]*:.*\)/\x1b[90m\1\x1b[0m/' \
	  -e "s|^\\(--- FAIL:.*\\)|$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \\x1b[90m\\1\\x1b[0m|" \
	  -e "s|^\\(--- PASS:.*\\)|$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \\x1b[90m\\1\\x1b[0m|" \
	  -e "s|^\\(    --- FAIL:.*\\)|$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \\x1b[90m\\1\\x1b[0m|" \
	  -e "s|^\\(    --- PASS:.*\\)|$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \\x1b[90m\\1\\x1b[0m|" \
	  -e "s|^PASS$$|$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \\x1b[90mPASS\\x1b[0m|" \
	  -e "s|^FAIL$$|$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \\x1b[90mFAIL\\x1b[0m|" \
	  -e "s|^\\(coverage:.*\\)|$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \\x1b[90m\\1\\x1b[0m|" \
	  -e "s|^\\(ok.*trunkform-mcp.*\\)|$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \\x1b[90m\\1\\x1b[0m|" \
	  -e "s|^\\(FAIL[[:space:]].*trunkform-mcp.*\\)|$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \\x1b[90m\\1\\x1b[0m|"
	@if grep -q "^FAIL$$" /tmp/test_output.txt; then \
		echo "$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \033[31m✗ Unit tests failed\033[0m"; \
		exit 1; \
	fi
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then \
		go tool cover -func=coverage.out | grep -v "total:" | while read line; do \
			echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[90m$$line\033[0m"; \
		done; \
	fi
	@if go tool cover -func=coverage.out | grep -v "total:" | grep -v "100.0%" > /dev/null; then \
		echo "$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \033[31m✗ Unit tests failed - Not all functions have 100% coverage\033[0m"; \
		exit 1; \
	fi
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Unit tests passed\033[0m"; fi
	@echo

local-int-test-http: kill
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== INTEGRATION TEST ===\033[0m"; fi
	@set -e; \
	if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mStarting server...\033[0m"; fi; \
	TRUNKFORM_LOGLEVEL=$(MAKE_LOGLEVEL) ./bin/trunkform-mcp-$$(go env GOOS)-$$(go env GOARCH) & \
	PID=$$!; \
	cleanup(){ kill $$PID 2>/dev/null || true; }; \
	trap cleanup EXIT; \
	sleep 1; \
	if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mRunning MCP tests...\033[0m"; fi; \
	SID=$$(curl -s -D - -H 'Content-Type: application/json' -X POST \
	  -d '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-03-26","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}}}' \
	  $(MCP) | awk '/^[Mm]cp-[Ss]ession-[Ii]d:/{gsub(/\r/,"");gsub(/^[^:]*: */,"");print}'); \
	curl -s -H "Mcp-Session-Id: $$SID" -H 'Content-Type: application/json' -X POST \
	  -d '{"jsonrpc":"2.0","method":"notifications/initialized","params":{}}' $(MCP) >/dev/null; \
	if [ $(MAKE_LOGLEVEL) -ge 5 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [DEBUG] \x1b[90m$$(curl -s -H "Mcp-Session-Id: $$SID" -H 'Content-Type: application/json' -X POST \
	  -d '{"jsonrpc":"2.0","id":2,"method":"prompts/list","params":{}}' $(MCP))\x1b[0m"; fi; \
	if [ $(MAKE_LOGLEVEL) -ge 5 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [DEBUG] \x1b[90m$$(curl -s -H "Mcp-Session-Id: $$SID" -H 'Content-Type: application/json' -X POST \
	  -d '{"jsonrpc":"2.0","id":3,"method":"tools/list","params":{}}' $(MCP))\x1b[0m"; fi; \
	if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mCall trunkform tool w/ perf-test-implemented:false...\033[0m"; fi; \
	RESPONSE=$$(curl -s -H "Mcp-Session-Id: $$SID" -H 'Content-Type: application/json' -X POST \
	  -d '{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"trunkform","arguments":{"version":"0.1.0","tools":["test"],"steering":"strict","related":[],"rubric":{"perf-test-implemented":false}}}}' $(MCP)); \
	if echo "$$RESPONSE" | grep -q "instructions for the AI assistant to act on behalf of user's needs:"; then \
		if [ $(MAKE_LOGLEVEL) -ge 5 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [DEBUG] \x1b[90mResponse: $$RESPONSE\x1b[0m"; fi; \
		if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Trunkform tool test passed\033[0m"; fi; \
	else \
		if [ $(MAKE_LOGLEVEL) -ge 5 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [DEBUG] \x1b[90mResponse: $$RESPONSE\x1b[0m"; fi; \
		echo "$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \033[31m✗ Trunkform tool test failed - expected text not found\033[0m"; \
		exit 1; \
	fi; \
	if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Integration test passed\033[0m"; fi
	@echo

local-perf-test: kill
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== PERFORMANCE TEST ===\033[0m"; fi
	@set -e; \
	if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mStarting server...\033[0m"; fi; \
	TRUNKFORM_LOGLEVEL=$(MAKE_LOGLEVEL) ./bin/trunkform-mcp-$$(go env GOOS)-$$(go env GOARCH) & \
	PID=$$!; \
	cleanup(){ kill $$PID 2>/dev/null || true; }; \
	trap cleanup EXIT; \
	sleep 1; \
	if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mRunning performance test...\033[0m"; fi; \
	go test -run=TestTrunkformLoadTest -v 2>&1 | tee /tmp/test_output.txt | sed \
	  -e '/^=== RUN/d' \
	  -e '/^exit status 1$$/d' \
	  -e 's/^\(#.*\)/\x1b[90m\1\x1b[0m/' \
	  -e 's/^\(.*\.go:[0-9]*:.*\)/\x1b[90m\1\x1b[0m/' \
	  -e "s|^\\(--- FAIL:.*\\)|$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \\x1b[90m\\1\\x1b[0m|" \
	  -e "s|^\\(--- PASS:.*\\)|$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \\x1b[90m\\1\\x1b[0m|" \
	  -e "s|^PASS$$|$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \\x1b[90mPASS\\x1b[0m|" \
	  -e "s|^FAIL$$|$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \\x1b[90mFAIL\\x1b[0m|" \
	  -e "s|^\\(ok.*trunkform-mcp.*\\)|$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \\x1b[90m\\1\\x1b[0m|" \
	  -e "s|^\\(FAIL[[:space:]].*trunkform-mcp.*\\)|$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \\x1b[90m\\1\\x1b[0m|"; \
	if grep -q "^PASS$$" /tmp/test_output.txt; then \
		if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Performance test completed\033[0m"; fi; \
	else \
		echo "$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \033[31m✖︎ Performance test failed\033[0m"; \
		exit 1; \
	fi

install: deps build
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== INSTALL ===\033[0m"; fi
	@GOBIN_PATH=$$(go env GOPATH)/bin; \
	if ! echo "$$PATH" | grep -q "$$GOBIN_PATH"; then \
		echo "$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \033[31m✗ Error: $$GOBIN_PATH not in PATH\033[0m"; \
		echo "$$(date '+%Y/%m/%d %H:%M:%S') [WARN] \033[33mAdd this to your shell profile: export PATH=\"\$$(go env GOPATH)/bin:\$$PATH\"\033[0m"; \
		exit 1; \
	fi
	@cp bin/trunkform-mcp-$$(go env GOOS)-$$(go env GOARCH) $$(go env GOPATH)/bin/trunkform-mcp
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Installed globally: $$(go env GOPATH)/bin/trunkform-mcp\033[0m"; fi


