.PHONY: all deps build unit-test test local-int-test-http local-perf-test install kill

MAKE_LOGLEVEL ?= 4

all: deps build test message-install

test: deps build unit-test local-int-test-http local-perf-test 

kill:
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== KILL (port 8080) ===\033[0m"; fi
	-@lsof -ti:8080 | xargs kill 2>/dev/null || true
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mKill completed\033[0m"; fi

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

unit-test:
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

local-int-test-http:
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== INTEGRATION TEST ===\033[0m"; fi
	@PIDS=$$(lsof -ti:8081 2>/dev/null); \
	if [ -n "$$PIDS" ]; then \
		if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mKilling existing processes on port 8081...\033[0m"; fi; \
		echo "$$PIDS" | xargs kill 2>/dev/null || true; \
	else \
		if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mPort 8081 is clean\033[0m"; fi; \
	fi
	@set -e; \
	if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mStarting server...\033[0m"; fi; \
	TRUNKFORM_PORT=8081 TRUNKFORM_LOGLEVEL=$(MAKE_LOGLEVEL) ./bin/trunkform-mcp-$$(go env GOOS)-$$(go env GOARCH) & \
	PID=$$!; \
	cleanup(){ \
		kill $$PID 2>/dev/null || true; \
		if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mKilled server on port 8081\033[0m"; fi; \
	}; \
	trap cleanup EXIT; \
	sleep 1; \
	if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mRunning MCP tests...\033[0m"; fi; \
	SID=$$(curl -s -D - -H 'Content-Type: application/json' -X POST \
	  -d '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-03-26","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}}}' \
	  http://127.0.0.1:8081/mcp | awk '/^[Mm]cp-[Ss]ession-[Ii]d:/{gsub(/\r/,"");gsub(/^[^:]*: */,"");print}'); \
	curl -s -H "Mcp-Session-Id: $$SID" -H 'Content-Type: application/json' -X POST \
	  -d '{"jsonrpc":"2.0","method":"notifications/initialized","params":{}}' http://127.0.0.1:8081/mcp >/dev/null; \
	if [ $(MAKE_LOGLEVEL) -ge 5 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [DEBUG] \x1b[90m$$(curl -s -H "Mcp-Session-Id: $$SID" -H 'Content-Type: application/json' -X POST \
	  -d '{"jsonrpc":"2.0","id":2,"method":"prompts/list","params":{}}' http://127.0.0.1:8081/mcp)\x1b[0m"; fi; \
	if [ $(MAKE_LOGLEVEL) -ge 5 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [DEBUG] \x1b[90m$$(curl -s -H "Mcp-Session-Id: $$SID" -H 'Content-Type: application/json' -X POST \
	  -d '{"jsonrpc":"2.0","id":3,"method":"tools/list","params":{}}' http://127.0.0.1:8081/mcp)\x1b[0m"; fi; \
	if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mCall trunkform tool w/ perf-test-implemented:false...\033[0m"; fi; \
	RESPONSE=$$(curl -s -H "Mcp-Session-Id: $$SID" -H 'Content-Type: application/json' -X POST \
	  -d '{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"trunkform","arguments":{"version":"0.1.0","tools":["test"],"steering":"strict","related":[],"rubric":{"perf-test-implemented":false}}}}' http://127.0.0.1:8081/mcp); \
	if echo "$$RESPONSE" | grep -q "INSTRUCTIONS FOR THE AI ASSISTANT TO ACT ON BEHALF OF USER'S NEEDS:"; then \
		if [ $(MAKE_LOGLEVEL) -ge 5 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [DEBUG] \x1b[90mResponse: $$RESPONSE\x1b[0m"; fi; \
		if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Trunkform tool test passed\033[0m"; fi; \
	else \
		if [ $(MAKE_LOGLEVEL) -ge 5 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [DEBUG] \x1b[90mResponse: $$RESPONSE\x1b[0m"; fi; \
		echo "$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \033[31m✗ Trunkform tool test failed - expected text not found\033[0m"; \
		exit 1; \
	fi; \
	if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Integration test passed\033[0m"; fi
	@echo

local-perf-test:
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== PERFORMANCE TEST ===\033[0m"; fi
	@PIDS=$$(lsof -ti:8081 2>/dev/null); \
	if [ -n "$$PIDS" ]; then \
		if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mKilling existing processes on port 8081...\033[0m"; fi; \
		echo "$$PIDS" | xargs kill 2>/dev/null || true; \
	else \
		if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mPort 8081 is clean\033[0m"; fi; \
	fi
	@set -e; \
	if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mStarting server...\033[0m"; fi; \
	TRUNKFORM_PORT=8081 TRUNKFORM_LOGLEVEL=$(MAKE_LOGLEVEL) ./bin/trunkform-mcp-$$(go env GOOS)-$$(go env GOARCH) & \
	PID=$$!; \
	cleanup(){ \
		kill $$PID 2>/dev/null || true; \
		if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mKilled server on port 8081\033[0m"; fi; \
	}; \
	trap cleanup EXIT; \
	sleep 1; \
	if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[33mRunning performance test...\033[0m"; fi; \
	TRUNKFORM_PORT=8081 go test -run=TestTrunkformLoadTest -v 2>&1 | tee /tmp/test_output.txt | sed \
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

message-install:
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== READY TO INSTALL ===\033[0m"; fi
	@echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ All tests passed. Run 'make install' to install the MCP binary globally.\033[0m"
	@echo

install: deps build kill
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== INSTALL ===\033[0m"; fi
	@GOBIN_PATH=$$(go env GOPATH)/bin; \
	if ! echo "$$PATH" | grep -q "$$GOBIN_PATH"; then \
		echo "$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \033[31m✗ Error: $$GOBIN_PATH not in PATH\033[0m"; \
		echo "$$(date '+%Y/%m/%d %H:%M:%S') [WARN] \033[33mAdd this to your shell profile: export PATH=\"\$$(go env GOPATH)/bin:\$$PATH\"\033[0m"; \
		exit 1; \
	fi
	@cp bin/trunkform-mcp-$$(go env GOOS)-$$(go env GOARCH) $$(go env GOPATH)/bin/trunkform-mcp
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then \
		echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Installed globally: $$(go env GOPATH)/bin/trunkform-mcp\033[0m"; \
		echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ You can now run 'trunkform-mcp' from anywhere or 'make start'\033[0m"; \
	fi

start:
	@echo "Starting trunkform MCP server on port 8080..."
	@trunkform-mcp &
