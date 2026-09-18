//go:build integration || perftest

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"testing"
	"time"
)

// binaryPath resolves the freshly-built binary under test. Set TRUNKFORM_BIN
// to override (the Makefile always points this at the binary it just built,
// so tests always exercise the exact artifact `make build` produced).
func binaryPath() string {
	if bin := os.Getenv("TRUNKFORM_BIN"); bin != "" {
		return bin
	}
	return fmt.Sprintf("./bin/trunkform-%s-%s", runtime.GOOS, runtime.GOARCH)
}

func writeLine(w *bufio.Writer, v map[string]interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if _, err := w.Write(b); err != nil {
		return err
	}
	if err := w.WriteByte('\n'); err != nil {
		return err
	}
	return w.Flush()
}

// freePort asks the OS for an ephemeral port, then releases it immediately
// so a server subprocess can bind it. Small TOCTOU race is acceptable for a
// local test.
func freePort(t *testing.T) int {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to find a free port: %v", err)
	}
	defer func() { _ = l.Close() }()
	return l.Addr().(*net.TCPAddr).Port
}

// waitForPort polls until the given port accepts TCP connections or timeout
// elapses, replacing an arbitrary fixed sleep.
func waitForPort(t *testing.T, port int, timeout time.Duration) {
	deadline := time.Now().Add(timeout)
	addr := "127.0.0.1:" + strconv.Itoa(port)
	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", addr, 100*time.Millisecond)
		if err == nil {
			_ = conn.Close()
			return
		}
		time.Sleep(50 * time.Millisecond)
	}
	t.Fatalf("server did not start listening on %s within %v", addr, timeout)
}

func postJSON(t *testing.T, url string, body map[string]interface{}, sessionID string) *http.Response {
	b, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if sessionID != "" {
		req.Header.Set("Mcp-Session-Id", sessionID)
	}
	resp, err := (&http.Client{Timeout: 5 * time.Second}).Do(req)
	if err != nil {
		t.Fatalf("request to %s failed: %v", url, err)
	}
	return resp
}
