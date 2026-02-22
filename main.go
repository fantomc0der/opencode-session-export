package main

import (
	"fmt"
	"os"

	"github.com/fantomc0der/opencode-session-export/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
