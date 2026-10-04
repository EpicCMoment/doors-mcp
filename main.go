package main

import (
	"fmt"
	"os"

	backend "gitlab.com/ariffil/doors-backend"
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
	fmt.Fprintf(os.Stderr, "loaded config: doorsPath=%q host=%s port=%d\n", cfg.DoorsPath, cfg.Host, cfg.Port)
}
