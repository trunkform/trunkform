# TRUNKFORM MCP SERVER (Go, Streamable HTTP)

[![trunkform-release](https://github.com/trunkform/trunkform/actions/workflows/trunkform-release.yml/badge.svg)](https://github.com/trunkform/trunkform/actions/workflows/trunkform-release.yml)

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

1. Make sure `$(go env GOPATH)/bin` is on your `PATH`:

   ```sh
   export PATH="$(go env GOPATH)/bin:$PATH"
   ```

   Add that line to your shell profile (`~/.zshrc`, `~/.bashrc`, etc.) to persist it.

2. Install the `trunkform` binary:

   ```sh
   go install github.com/trunkform/trunkform/trunkform@latest
   ```

   Or download the latest tag's binaries for your OS/arch from [GitHub Releases](https://github.com/trunkform/trunkform/releases) and place it on your `PATH` as `trunkform`.

3. Verify it's installed:

   ```sh
   which trunkform
   ```

   Your MCP client (see configs below) will invoke `trunkform` directly to start the stdio server - you normally won't run it manually.

### Local dev (build from source)

```sh
make start
```

## CLAUDE DESKTOP MCP CONFIG

```json
// ~/.config/Claude/claude_desktop_config.json
{
  "mcpServers": {
    "trunkform": {
      "type": "stdio",
      "command": "trunkform"
    }
  }
}
```

## OPENCODE MCP CONFIG

```json
// opencode.json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {
    "trunkform": {
      "type": "local",
      "command": ["trunkform"],
      "enabled": true
    }
  }
}
```

## COPILOT MCP CONFIG

```json
// ~/.copilot/mcp-config.json
{
  "mcpServers": {
    "trunkform": {
      "type": "stdio",
      "command": "trunkform"
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
      "command": "trunkform"
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
      "command": "trunkform"
    }
  }
}
```

## CODEX MCP CONFIG

```toml
# ~/.codex/config.toml
[mcp_servers.trunkform]
command = "trunkform"
```
