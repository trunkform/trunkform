package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mark3labs/mcp-go/server"
)

var startHTTPServer = func(port string) error {
	s := server.NewMCPServer("trunkform", "0.1.0")
	NewTrunkformTool().Register(s)

	http := server.NewStreamableHTTPServer(
		s,
		server.WithEndpointPath("/mcp"),
		server.WithHeartbeatInterval(30*time.Second),
	)

	logf.Infof("Starting HTTP server on :"+port+" (endpoint /mcp)", "\x1b[90m")
	return http.Start(":" + port)
}

var startStdioServer = func() error {
	s := server.NewMCPServer("trunkform", "0.1.0")
	NewTrunkformTool().Register(s)

	stdio := server.NewStdioServer(s)
	logf.Infof("Starting stdio server (npx mode)", "\x1b[90m")
	return stdio.Listen(context.Background(), os.Stdin, os.Stdout)
}

func main() {
	logf.Debugf("main: starting trunkform server", "\x1b[90m")
	
	// Check for npx/stdio mode (no args or explicit stdio flag)
	if len(os.Args) > 1 && os.Args[1] == "--http" {
		// HTTP mode
		port := os.Getenv("TRUNKFORM_PORT")
		if port == "" {
			port = "8080"
		}
		logf.Debugf("main: using HTTP mode on port "+port, "\x1b[90m")
		log.Fatal(startHTTPServer(port))
	} else {
		// Stdio mode (default for npx)
		logf.Debugf("main: using stdio mode", "\x1b[90m")
		log.Fatal(startStdioServer())
	}
}
