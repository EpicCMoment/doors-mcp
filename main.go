package main

import (
	"context"
	"fmt"
	"os"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"

	backend "github.com/EpicCMoment/doors-backend"
)

func main() {
	cfgPath := "config.json"
	if len(os.Args) > 1 {
		cfgPath = os.Args[1]
	}
	cfg, err := backend.LoadConfig(cfgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "config error: %v\n", err)
		os.Exit(1)
	}

	d, err := backend.Init(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: backend init failed: %v (tools will report an error)\n", err)
		d = nil
	}

	server := mcp.NewServer(&mcp.Implementation{Name: "doors-mcp", Version: "0.1.0"}, nil)
	registerTools(server, d)
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
