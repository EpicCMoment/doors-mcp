# doors-mcp

MCP server (stdio) wrapping `doors-backend` for AI agent access to DOORS 9.7.

## Facts an agent needs

- Go module `github.com/EpicCMoment/doors-mcp`, sibling module `doors-backend` is pulled in via a local `replace` => `../doors-backend`. Builds succeed only when both directories exist side by side.
- Config is JSON (`config.json` in cwd, or pass path as arg), parsed with stdlib `encoding/json` — no yaml lib. `Config` defaults live in `doors-backend.DefaultConfig()`; see keys in `config.json`.
- MCP transport: stdio via official `modelcontextprotocol/go-sdk`.
- Tool surface: `list_items`, `list_modules`, `get_module`, `get_requirements`, `get_baselines`, `traverse_hierarchy` (see `tools.go`). `ping` intentionally not exposed — `backend.Init` runs it as a connectivity self-test.

## Commands

```sh
go vet ./...
go test ./...
go build ./...
./doors-mcp config.json
```
