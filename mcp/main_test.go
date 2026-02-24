package main

import (
  "os"
  "testing"
)

func TestMain(t *testing.T) {
  if os.Getenv("SKIP_MAIN_TEST") != "" {
    t.Skip()
  }
  
  originalStartServer := startServer
  defer func() { 
    startServer = originalStartServer
    recover()
  }()
  
  called := false
  startServer = func(port string) error {
    called = true
    if port != "9999" {
      t.Errorf("unexpected port: %s", port)
    }
    panic("exit test")
  }
  
  os.Setenv("PORT", "9999")
  defer os.Unsetenv("PORT")
  
  func() {
    defer func() { recover() }()
    main()
  }()
  
  if !called {
    t.Error("startServer not called")
  }
}

func TestMainDefaultPort(t *testing.T) {
  if os.Getenv("SKIP_MAIN_TEST") != "" {
    t.Skip()
  }
  
  originalStartServer := startServer
  defer func() { 
    startServer = originalStartServer
    recover()
  }()
  
  os.Unsetenv("PORT")
  
  called := false
  startServer = func(port string) error {
    called = true
    if port != "8080" {
      t.Errorf("expected default port 8080, got %s", port)
    }
    panic("exit test")
  }
  
  func() {
    defer func() { recover() }()
    main()
  }()
  
  if !called {
    t.Error("startServer not called")
  }
}
