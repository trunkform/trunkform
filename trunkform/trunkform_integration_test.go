//go:build integration

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"
)

// TestStdioIntegration spawns the built binary over stdio - the default
// transport - and drives it through the same initialize -> initialized ->
// tools/list -> tools/call handshake a real MCP client (OpenCode, Claude,
// etc.) would perform, then asserts the tool response contains the expected
// instructions text.
func TestStdioIntegration(t *testing.T) {
	logf.Infof("Running MCP tests over stdio...", "\x1b[33m")

	bin := binaryPath()
	if _, err := os.Stat(bin); err != nil {
		t.Fatalf("binary under test not found at %q (set TRUNKFORM_BIN or run `make build` first): %v", bin, err)
	}

	cmd := exec.Command(bin)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatalf("failed to open stdin pipe: %v", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatalf("failed to open stdout pipe: %v", err)
	}
	cmd.Stderr = io.Discard

	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start binary: %v", err)
	}
	defer func() {
		_ = stdin.Close()
		_ = cmd.Wait()
	}()

	w := bufio.NewWriter(stdin)
	scanner := bufio.NewScanner(stdout)

	if err := writeLine(w, map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "initialize",
		"params": map[string]interface{}{
			"protocolVersion": "2025-03-26",
			"capabilities":    map[string]interface{}{},
			"clientInfo":      map[string]interface{}{"name": "test", "version": "1.0"},
		},
	}); err != nil {
		t.Fatalf("failed to write initialize request: %v", err)
	}
	if !scanner.Scan() {
		t.Fatalf("no response to initialize request")
	}

	if err := writeLine(w, map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "notifications/initialized",
		"params":  map[string]interface{}{},
	}); err != nil {
		t.Fatalf("failed to write initialized notification: %v", err)
	}

	if err := writeLine(w, map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      2,
		"method":  "tools/list",
		"params":  map[string]interface{}{},
	}); err != nil {
		t.Fatalf("failed to write tools/list request: %v", err)
	}
	if !scanner.Scan() {
		t.Fatalf("no response to tools/list request")
	}

	if err := writeLine(w, map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      3,
		"method":  "tools/call",
		"params": map[string]interface{}{
			"name": "trunkform",
			"arguments": map[string]interface{}{
				"version": "0.1.0",
				"tools":   []string{"test"},
				"related": []string{},
				"rubric": map[string]interface{}{
					"perf-test-implemented": nil,
				},
			},
		},
	}); err != nil {
		t.Fatalf("failed to write tools/call request: %v", err)
	}
	if !scanner.Scan() {
		t.Fatalf("no response to tools/call request")
	}

	response := scanner.Text()
	if !strings.Contains(response, "INSTRUCTIONS FOR THE AI ASSISTANT TO ACT ON BEHALF OF USER'S NEEDS:") {
		t.Fatalf("trunkform tool response missing expected instructions text: %s", response)
	}

	logf.Infof("✓ Trunkform tool test passed", "\x1b[32m")
}

// TestHTTPIntegration spawns the built binary in --http mode - the
// secondary transport, intended for a future hosted deployment - and drives
// it through the same initialize -> initialized -> tools/call handshake a
// real MCP HTTP client would perform, then asserts the tool response
// contains the expected instructions text.
func TestHTTPIntegration(t *testing.T) {
	logf.Infof("Running MCP tests over http...", "\x1b[33m")

	bin := binaryPath()
	if _, err := os.Stat(bin); err != nil {
		t.Fatalf("binary under test not found at %q (set TRUNKFORM_BIN or run `make build` first): %v", bin, err)
	}

	port := freePort(t)
	cmd := exec.Command(bin, "--http")
	cmd.Env = append(os.Environ(), "TRUNKFORM_PORT="+strconv.Itoa(port))
	cmd.Stderr = io.Discard
	cmd.Stdout = io.Discard

	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start binary: %v", err)
	}
	defer func() {
		_ = cmd.Process.Kill()
		_ = cmd.Wait()
	}()

	waitForPort(t, port, 5*time.Second)

	url := fmt.Sprintf("http://127.0.0.1:%d/mcp", port)

	initResp := postJSON(t, url, map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "initialize",
		"params": map[string]interface{}{
			"protocolVersion": "2025-03-26",
			"capabilities":    map[string]interface{}{},
			"clientInfo":      map[string]interface{}{"name": "test", "version": "1.0"},
		},
	}, "")
	sessionID := initResp.Header.Get("Mcp-Session-Id")
	_ = initResp.Body.Close()
	if sessionID == "" {
		t.Fatalf("no Mcp-Session-Id header in initialize response")
	}

	notifyResp := postJSON(t, url, map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "notifications/initialized",
		"params":  map[string]interface{}{},
	}, sessionID)
	_ = notifyResp.Body.Close()

	callResp := postJSON(t, url, map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      2,
		"method":  "tools/call",
		"params": map[string]interface{}{
			"name": "trunkform",
			"arguments": map[string]interface{}{
				"version": "0.1.0",
				"tools":   []string{"test"},
				"related": []string{},
				"rubric": map[string]interface{}{
					"perf-test-implemented": nil,
				},
			},
		},
	}, sessionID)
	defer func() { _ = callResp.Body.Close() }()

	body, err := io.ReadAll(callResp.Body)
	if err != nil {
		t.Fatalf("failed to read tools/call response body: %v", err)
	}

	response := string(body)
	if !strings.Contains(response, "INSTRUCTIONS FOR THE AI ASSISTANT TO ACT ON BEHALF OF USER'S NEEDS:") {
		t.Fatalf("trunkform tool response missing expected instructions text: %s", response)
	}

	logf.Infof("✓ Trunkform tool test passed", "\x1b[32m")
}
