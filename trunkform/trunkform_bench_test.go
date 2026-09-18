//go:build perftest

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"testing"
	"time"
)

// TestTrunkformLoadTestHTTP load-tests a Streamable HTTP MCP server
// (secondary transport, intended for a future hosted deployment). By
// default it spawns its own instance of the freshly-built binary on an
// OS-assigned free port, so it never conflicts with anything else running
// locally. Set TRUNKFORM_MCP_PERF_TEST_URL to the full /mcp endpoint of an
// already-running server (local or hosted) to load-test that instead -
// no local binary is spawned in that case. Requests are paced evenly across
// testDuration via a ticker, so this measures whether the server sustains a
// fixed concurrent request rate without falling over - not raw maximum
// throughput.
//
// Note: this is HTTP-specific because it's the only transport with a shared
// server to contend over. Stdio is 1-process-per-client by design, so
// there's no analogous "capacity" to load-test.
func TestTrunkformLoadTestHTTP(t *testing.T) {
	mcpURL := os.Getenv("TRUNKFORM_MCP_PERF_TEST_URL")
	if mcpURL == "" {
		bin := binaryPath()
		if _, err := os.Stat(bin); err != nil {
			t.Fatalf("binary under test not found at %q (set TRUNKFORM_BIN or run `make build` first): %v", bin, err)
		}

		portNum := freePort(t)
		port := strconv.Itoa(portNum)

		cmd := exec.Command(bin, "--http")
		cmd.Env = append(os.Environ(), "TRUNKFORM_PORT="+port)
		cmd.Stderr = io.Discard
		cmd.Stdout = io.Discard

		if err := cmd.Start(); err != nil {
			t.Fatalf("failed to start binary: %v", err)
		}
		defer func() {
			_ = cmd.Process.Kill()
			_ = cmd.Wait()
		}()

		waitForPort(t, portNum, 5*time.Second)
		mcpURL = "http://127.0.0.1:" + port + "/mcp"
		logf.Infof("Load-testing local instance at "+mcpURL, "\x1b[90m")
	} else {
		logf.Infof("Load-testing remote instance at "+mcpURL, "\x1b[90m")
	}

	totalRequests := 10000
	testDuration := 30 * time.Second

	var wg sync.WaitGroup
	var successCount, errorCount int64
	var mu sync.Mutex

	start := time.Now()

	ticker := time.NewTicker(testDuration / time.Duration(totalRequests))
	defer ticker.Stop()

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			client := &http.Client{Timeout: 5 * time.Second}

			// Initialize session
			initPayload := map[string]interface{}{
				"jsonrpc": "2.0",
				"id":      1,
				"method":  "initialize",
				"params": map[string]interface{}{
					"protocolVersion": "2025-03-26",
					"capabilities":    map[string]interface{}{},
					"clientInfo":      map[string]interface{}{"name": "loadtest", "version": "1.0.0"},
				},
			}
			body, _ := json.Marshal(initPayload)
			req, _ := http.NewRequest("POST", mcpURL, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				mu.Lock()
				errorCount++
				mu.Unlock()
				return
			}
			sessionID := resp.Header.Get("Mcp-Session-Id")
			_ = resp.Body.Close()

			// Call trunkform tool
			trunkformPayload := map[string]interface{}{
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
			}
			body, _ = json.Marshal(trunkformPayload)
			req, _ = http.NewRequest("POST", mcpURL, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Mcp-Session-Id", sessionID)

			resp, err = client.Do(req)
			mu.Lock()
			if err != nil {
				errorCount++
			} else {
				successCount++
				_ = resp.Body.Close()
			}
			mu.Unlock()
		}()

		<-ticker.C
	}

	wg.Wait()
	elapsed := time.Since(start)

	logf.Infof(fmt.Sprintf("Load test completed in %v", elapsed), "\x1b[90m")
	logf.Infof(fmt.Sprintf("Successful requests: %d", successCount), "\x1b[90m")
	logf.Infof(fmt.Sprintf("Failed requests: %d", errorCount), "\x1b[90m")
	logf.Infof(fmt.Sprintf("Actual RPS: %.2f", float64(successCount)/elapsed.Seconds()), "\x1b[90m")

	minSuccess := int64(float64(totalRequests) * 0.95)
	if successCount < minSuccess {
		t.Fatalf("Too many failures: %d/%d succeeded", successCount, totalRequests)
	}

	actualRPS := float64(successCount) / elapsed.Seconds()
	if actualRPS < 300 {
		t.Fatalf("Throughput too low: %.2f RPS (expected >= 300)", actualRPS)
	}

	logf.Infof("✓ Performance test (http) completed", "\x1b[32m")
}
