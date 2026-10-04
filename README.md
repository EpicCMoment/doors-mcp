# doors-mcp

MCP server (stdio) wrapping `doors-backend` for AI agent access to DOORS 9.7.

## Status

Stub: `main.go` currently loads a JSON config and reports what it loaded. The
MCP tool bindings are planned (official `modelcontextprotocol/go-sdk`).

## Run

```sh
go build ./...
./doors-mcp config.json
```

`config.json` — see the committed sample. Config keys map to
`doors-backend.Config`.
