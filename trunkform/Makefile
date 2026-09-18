.PHONY: all deps mcp-go-upgrade build unit-test test local-int-test-stdio local-int-test-http local-perf-test-http install clean

MAKE_LOGLEVEL ?= 4
MAKE_ALL ?=

all: deps lint build test message-install

test: deps lint build unit-test local-int-test-stdio local-int-test-http local-perf-test-http

clean:
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== CLEAN ===\033[0m"; fi
	@go clean -cache
	@rm -f coverage.out coverage_*.out *.html
	@rm -rf bin/ .gotmp/ .gocache/
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Clean completed\033[0m"; fi

deps: clean
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== DEPENDENCIES ===\033[0m"; fi
	@go mod download
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Dependencies ready\033[0m"; fi
	@echo

mcp-go-upgrade: clean
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== MCP-GO UPGRADE ===\033[0m"; fi
	@go get github.com/mark3labs/mcp-go@latest
	@go mod tidy
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ mcp-go upgraded and module files tidied\033[0m"; fi
	@echo

lint: deps
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== LINT ===\033[0m"; fi
	@go tool golangci-lint run
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Lint passed\033[0m"; fi
	@echo

build: deps
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== BUILD ===\033[0m"; fi
	@mkdir -p bin
	@if [ -n "$(MAKE_ALL)" ]; then \
		GOOS=linux GOARCH=amd64 go build -o bin/trunkform-linux-amd64 .; \
		GOOS=linux GOARCH=arm64 go build -o bin/trunkform-linux-arm64 .; \
		GOOS=darwin GOARCH=amd64 go build -o bin/trunkform-darwin-amd64 .; \
		GOOS=darwin GOARCH=arm64 go build -o bin/trunkform-darwin-arm64 .; \
		GOOS=windows GOARCH=amd64 go build -o bin/trunkform-windows-amd64.exe .; \
		GOOS=windows GOARCH=arm64 go build -o bin/trunkform-windows-arm64.exe .; \
	else \
		go build -o bin/trunkform-$$(go env GOOS)-$$(go env GOARCH) .; \
	fi
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Binaries built\033[0m"; fi
	@shasum -a 256 bin/* > bin/SHA256SUMS
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Checksums generated\033[0m"; fi
	@echo

unit-test: build
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== UNIT TESTS ===\033[0m"; fi
	@PKGS=$$(go list ./... | tr '\n' ',' | sed 's/,$$//'); \
	go test ./... -coverpkg=$$PKGS -coverprofile=coverage.out -v 2>&1 | tee /tmp/test_output.txt | sed \
	  -e '/^=== RUN/d' \
	  -e '/^=== PAUSE/d' \
	  -e '/^=== CONT/d' \
	  -e '/^coverage:/d' \
	  -e '/coverage: .* of statements/d' \
	  -e '/^[[:space:]]*.*coverage: .* of statements$$/d' \
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
	  -e "s|^\\(ok.*trunkform.*\\)|$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \\x1b[90m\\1\\x1b[0m|" \
	  -e "s|^\\(FAIL[[:space:]].*trunkform.*\\)|$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \\x1b[90m\\1\\x1b[0m|"
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
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[90mAll functions meet the enforced 100.0% coverage requirement\033[0m"; fi
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Unit tests passed\033[0m"; fi
	@echo

local-int-test-stdio: build
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== INTEGRATION TEST (stdio, default transport) ===\033[0m"; fi
	@TRUNKFORM_BIN=./bin/trunkform-$$(go env GOOS)-$$(go env GOARCH) go test -tags integration -run=TestStdioIntegration -v .
	@echo

local-int-test-http: build
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== INTEGRATION TEST (http, secondary transport) ===\033[0m"; fi
	@TRUNKFORM_BIN=./bin/trunkform-$$(go env GOOS)-$$(go env GOARCH) go test -tags integration -run=TestHTTPIntegration -v .
	@echo

local-perf-test-http: build
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== PERFORMANCE TEST (http, primary - shared server capacity) ===\033[0m"; fi
	@TRUNKFORM_BIN=./bin/trunkform-$$(go env GOOS)-$$(go env GOARCH) go test -tags perftest -run=TestTrunkformLoadTestHTTP -v .
	@echo

message-install:
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== READY TO INSTALL ===\033[0m"; fi
	@echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ All tests passed. Run 'make install' to install the MCP binary globally.\033[0m"
	@echo

install: build
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[34m=== INSTALL ===\033[0m"; fi
	@GOBIN_PATH=$$(go env GOPATH)/bin; \
	if ! echo "$$PATH" | grep -q "$$GOBIN_PATH"; then \
		echo "$$(date '+%Y/%m/%d %H:%M:%S') [ERROR] \033[31m✗ Error: $$GOBIN_PATH not in PATH\033[0m"; \
		echo "$$(date '+%Y/%m/%d %H:%M:%S') [WARN] \033[33mAdd this to your shell profile: export PATH=\"\$$(go env GOPATH)/bin:\$$PATH\"\033[0m"; \
		exit 1; \
	fi
	@GOBIN_PATH=$$(go env GOPATH)/bin; \
	cp bin/trunkform-$$(go env GOOS)-$$(go env GOARCH) "$$GOBIN_PATH/trunkform.tmp" && \
	mv "$$GOBIN_PATH/trunkform.tmp" "$$GOBIN_PATH/trunkform"
	@if [ $(MAKE_LOGLEVEL) -ge 4 ]; then \
		echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ Installed globally: $$(go env GOPATH)/bin/trunkform\033[0m"; \
		echo "$$(date '+%Y/%m/%d %H:%M:%S') [INFO] \033[32m✓ You can now run 'trunkform' from anywhere or 'make start'\033[0m"; \
	fi

start: build
	@echo "Starting trunkform MCP server on port 8080 (set TRUNKFORM_PORT to use a different port)..."
	@./bin/trunkform-$$(go env GOOS)-$$(go env GOARCH) --http
