package main

import (
	"github.com/togettoyou/go-one-server/cmd/gos/internal/swag"
	"github.com/togettoyou/go-one-server/cmd/gos/internal/upgrade"
	"log"

	"github.com/spf13/cobra"
	"github.com/togettoyou/go-one-server/cmd/gos/internal/project"
	"github.com/togettoyou/go-one-server/cmd/gos/internal/run"
)

var (
	version = "v1.0.2"

	rootCmd = &cobra.Command{
		Use:     "gos",
		Short:   "gos: An elegant toolkit for Go services.",
		Long:    `gos: An elegant toolkit for Go services.`,
		Version: version,
	}
)

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(run.CmdRun)
	rootCmd.AddCommand(swag.CmdSwag)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
