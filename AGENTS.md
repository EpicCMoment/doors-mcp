# doors-mcp

MCP server stub wrapping `doors-backend`.

## Facts an agent needs

- Go module `github.com/EpicCMoment/doors-mcp`, sibling module `doors-backend` is pulled in via a local `replace` => `../doors-backend`. Builds succeed only when both directories exist side by side.
- Config is JSON (`config.json` in cwd, or pass path as arg), parsed with stdlib `encoding/json` — no yaml lib. `Config` defaults live in `doors-backend.DefaultConfig()`; see keys in `config.json`.
- Transports planned: stdio via official `modelcontextprotocol/go-sdk` (not yet added); current `main.go` only loads config.
- When adding MCP tools, they should call the backend singleton: `doors-backend.Init(cfg)` then `controller.Modules()`.

## Commands

```sh
go vet ./...
go test ./...
go build ./...
./doors-mcp config.json
```
