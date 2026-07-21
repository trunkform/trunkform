//go:build perftest

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"
)

func TestTrunkformLoadTest(t *testing.T) {
	// Test 10,000 requests in 30 seconds = ~333 RPS
	totalRequests := 10000
	testDuration := 30 * time.Second

	port := os.Getenv("TRUNKFORM_PORT")
	if port == "" {
		port = "8080"
	}

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
			req, _ := http.NewRequest("POST", "http://localhost:"+port+"/mcp", bytes.NewBuffer(body))
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
			req, _ = http.NewRequest("POST", "http://localhost:"+port+"/mcp", bytes.NewBuffer(body))
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
}
