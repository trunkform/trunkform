# TRUNKFORM MCP SERVER (Go, Streamable HTTP)

[![mcp-release](https://github.com/trunkform/trunkform/actions/workflows/mcp-release.yml/badge.svg)](https://github.com/trunkform/trunkform/actions/workflows/mcp-release.yml)

## Architecture

It's just a server which uses the mcp-go library to handle incoming MCP requests. The MCP server is designed to be modular, allowing for easy addition of new rubric handlers. The directory structure is as follows:

```
rubricHandlers/
├── rubrichandler/rubric_handler.go     # interface
├── lintrubricbuilder/rubric_handler.go # First of many handler implementations
├── ...
└── registry.go                         # Handler registration
logf.go                                 # Logging utility
main.go                                 # Server setup and routing
trunkform_tool.go                       # 
```

# SETUP

```sh
go install github.com/trunkform/trunkform/mcp@latest
```

Or download the latest tag's binaries for your OS/arch from [GitHub Releases](https://github.com/trunkform/trunkform/releases).

### Local dev (build from source)

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
