# TRUNKFORM MCP SERVER (Go, Streamable HTTP)

## Architecture

```
rubricHandlers/
├── rubrichandler/              # Base interface
├── perftestrubricbuilder/      # First of seven handler implementations
├── ...
└── registry.go                 # Handler registration
```

# SETUP

```sh
make start
```

## COPILOT MCP CONFIG

```json
// ~/.copilot/mcp-config.json
{
  "mcpServers": {
    "trunkform": {
      "type": "http",
      "url": "http://localhost:8080/mcp"
    }
  }
}
```

## KIRO MCP CONFIG

```json
// ~/.kiro/agents/default.json
{
  "mcpServers": {
    "trunkform": {
      "url": "http://localhost:8080/mcp"
    }
  }
}
```

## GEMINI MCP CONFIG

```json
// ~/.gemini/settings.json
{
  "mcpServers": {
    "trunkform": {
      "httpUrl": "http://localhost:8080/mcp"
    }
  }
}
```

## CODEX MCP CONFIG

```toml
# ~/.codex/config.toml
[mcp_servers.trunkform]
enabled = true
url = "http://localhost:8080/mcp"
```
