package main

import (
	"log"
	"os"
	"time"

	"github.com/mark3labs/mcp-go/server"
)

var startServer = func(port string) error {
	s := server.NewMCPServer("trunkform", "0.1.0")
	NewTrunkformTool().Register(s)
	
	http := server.NewStreamableHTTPServer(
		s,
		server.WithEndpointPath("/mcp"),
		server.WithHeartbeatInterval(30*time.Second),
	)
	
	logf.Infof("Listening on :"+port+" (endpoint /mcp)", "\x1b[90m")
	return http.Start(":" + port)
}

func main() {
	logf.Debugf("main: starting trunkform-mcp server", "\x1b[90m")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logf.Debugf("main: using port "+port, "\x1b[90m")
	log.Fatal(startServer(port))
}
