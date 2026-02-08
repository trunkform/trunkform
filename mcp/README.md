# TRUNKFORM MCP SERVER (Go, Streamable HTTP)

# SETUP

Add the Go bin directory to your PATH .bashrc:

```sh
export PATH="$(go env GOPATH)/bin:$PATH"
```

make the Makefile, add the trunkform mcp server to your kiro agents config, and run the server and kiro-cli:

```bash
make
jq '.mcpServers = (.mcpServers // {}) + {"trunkform":{"url":"http://localhost:8080/mcp"}}' ~/.kiro/agents/default.json > /tmp/default.json && mv /tmp/default.json ~/.kiro/agents/default.json
trunkform-mcp & kiro-cli
```
