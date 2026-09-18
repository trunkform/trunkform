package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	if os.Getenv("SKIP_MAIN_TEST") != "" {
		t.Skip()
	}

	originalStartHTTPServer := startHTTPServer
	defer func() {
		startHTTPServer = originalStartHTTPServer
		_ = recover()
	}()

	called := false
	startHTTPServer = func(port string) error {
		called = true
		if port != "9999" {
			t.Errorf("unexpected port: %s", port)
		}
		panic("exit test")
	}

	_ = os.Setenv("TRUNKFORM_PORT", "9999")
	defer func() { _ = os.Unsetenv("TRUNKFORM_PORT") }()
	
	// Set --http flag to use HTTP mode
	originalArgs := os.Args
	os.Args = []string{"trunkform", "--http"}
	defer func() { os.Args = originalArgs }()

	func() {
		defer func() { _ = recover() }()
		main()
	}()

	if !called {
		t.Error("startHTTPServer not called")
	}
}

func TestMainDefaultPort(t *testing.T) {
	if os.Getenv("SKIP_MAIN_TEST") != "" {
		t.Skip()
	}

	originalStartHTTPServer := startHTTPServer
	defer func() {
		startHTTPServer = originalStartHTTPServer
		_ = recover()
	}()

	_ = os.Unsetenv("TRUNKFORM_PORT")

	called := false
	startHTTPServer = func(port string) error {
		called = true
		if port != "8080" {
			t.Errorf("expected default port 8080, got %s", port)
		}
		panic("exit test")
	}

	// Set --http flag to use HTTP mode
	originalArgs := os.Args
	os.Args = []string{"trunkform", "--http"}
	defer func() { os.Args = originalArgs }()

	func() {
		defer func() { _ = recover() }()
		main()
	}()

	if !called {
		t.Error("startHTTPServer not called")
	}
}

func TestMainDefaultStdioMode(t *testing.T) {
	if os.Getenv("SKIP_MAIN_TEST") != "" {
		t.Skip()
	}

	originalStartStdioServer := startStdioServer
	defer func() {
		startStdioServer = originalStartStdioServer
		_ = recover()
	}()

	called := false
	startStdioServer = func() error {
		called = true
		panic("exit test")
	}

	// No --http flag means stdio mode (default)
	originalArgs := os.Args
	os.Args = []string{"trunkform"}
	defer func() { os.Args = originalArgs }()

	func() {
		defer func() { _ = recover() }()
		main()
	}()

	if !called {
		t.Error("startStdioServer not called")
	}
}
