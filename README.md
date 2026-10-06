# doors-mcp

MCP server (stdio) wrapping `doors-backend` for AI agent access to DOORS 9.7.

## Status

Implemented: `main.go` loads a JSON config, starts the backend (`backend.Init`,
which also pings DOORS), and serves MCP tools on stdio. `ping` is *not* an
MCP tool — it's a backend self-test run at startup.

## Tools

- `list_items(path)` — projects/folders/modules directly under a path
- `list_modules(path)` — modules directly under a path (wraps `list_items`)
- `get_module(path, baseline?)`
- `get_requirements(modulePath, baseline?)`
- `get_baselines(modulePath)`
- `traverse_hierarchy(modulePath, baseline?)` — flat outline with parent links

## Run

```sh
go build ./...
./doors-mcp config.json
```

`config.json` — see the committed sample. Config keys map to
`doors-backend.Config`.
