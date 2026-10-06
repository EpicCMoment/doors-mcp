package main

import (
	"context"
	"encoding/json"
	"errors"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"

	backend "github.com/EpicCMoment/doors-backend"
)

var errBackendNotRunning = errors.New("backend not initialized (DOORS failed to start)")

// pingArgs / modulePathArgs are the small typed inputs for each tool.
type emptyArgs struct{}

type pathArgs struct {
	Path string `json:"path" jsonschema:"full DOORS path, e.g. /MyProject/Folder"`
}

type moduleArgs struct {
	Path       string `json:"path" jsonschema:"full DOORS path of the module"`
	BaselineID string `json:"baseline,omitempty" jsonschema:"optional baseline name; empty means the current view"`
}

type modulePathArgs struct {
	ModulePath string `json:"modulePath" jsonschema:"full DOORS path of the module"`
	BaselineID string `json:"baseline,omitempty" jsonschema:"optional baseline name; empty means the current view"`
}

func textResult(v any) (*mcp.CallToolResult, any, error) {
	raw, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: string(raw)}},
	}, nil, nil
}

// registerTools wires each MCP tool to the backend singleton.
func registerTools(server *mcp.Server, d *backend.DoorsController) {
	mcp.AddTool(server, &mcp.Tool{Name: "ping", Description: "smoke-test the DXL server"},
		func(ctx context.Context, req *mcp.CallToolRequest, args emptyArgs) (*mcp.CallToolResult, any, error) {
	if d == nil { return nil, nil, errBackendNotRunning }
			raw, err := d.Dxl().ExecTemplate("ping", nil)
			if err != nil {
				return nil, nil, err
			}
			return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: string(raw)}}}, nil, nil
		})

	mcp.AddTool(server, &mcp.Tool{Name: "list_items", Description: "list projects/folders/modules directly under a path"},
		func(ctx context.Context, req *mcp.CallToolRequest, args pathArgs) (*mcp.CallToolResult, any, error) {
	if d == nil { return nil, nil, errBackendNotRunning }
			items, err := d.Modules().ListItems(args.Path)
			if err != nil {
				return nil, nil, err
			}
			return textResult(items)
		})

	mcp.AddTool(server, &mcp.Tool{Name: "list_modules", Description: "list modules directly under a path"},
		func(ctx context.Context, req *mcp.CallToolRequest, args pathArgs) (*mcp.CallToolResult, any, error) {
	if d == nil { return nil, nil, errBackendNotRunning }
			mods, err := d.Modules().ListModules(args.Path)
			if err != nil {
				return nil, nil, err
			}
			return textResult(mods)
		})

	mcp.AddTool(server, &mcp.Tool{Name: "get_module", Description: "fetch a module's metadata"},
		func(ctx context.Context, req *mcp.CallToolRequest, args moduleArgs) (*mcp.CallToolResult, any, error) {
	if d == nil { return nil, nil, errBackendNotRunning }
			m, err := d.Modules().GetModule(args.Path, args.BaselineID)
			if err != nil {
				return nil, nil, err
			}
			return textResult(m)
		})

	mcp.AddTool(server, &mcp.Tool{Name: "get_requirements", Description: "fetch requirements of a module"},
		func(ctx context.Context, req *mcp.CallToolRequest, args modulePathArgs) (*mcp.CallToolResult, any, error) {
	if d == nil { return nil, nil, errBackendNotRunning }
			reqs, err := d.Modules().GetRequirements(args.ModulePath, args.BaselineID)
			if err != nil {
				return nil, nil, err
			}
			return textResult(reqs)
		})

	mcp.AddTool(server, &mcp.Tool{Name: "get_baselines", Description: "list a module's baselines"},
		func(ctx context.Context, req *mcp.CallToolRequest, args modulePathArgs) (*mcp.CallToolResult, any, error) {
	if d == nil { return nil, nil, errBackendNotRunning }
			bs, err := d.Modules().GetBaselines(args.ModulePath)
			if err != nil {
				return nil, nil, err
			}
			return textResult(bs)
		})

	mcp.AddTool(server, &mcp.Tool{Name: "traverse_hierarchy", Description: "flat requirement outline with parent links"},
		func(ctx context.Context, req *mcp.CallToolRequest, args modulePathArgs) (*mcp.CallToolResult, any, error) {
	if d == nil { return nil, nil, errBackendNotRunning }
			h, err := d.Modules().TraverseHierarchy(args.ModulePath, args.BaselineID)
			if err != nil {
				return nil, nil, err
			}
			return textResult(h)
		})

}
