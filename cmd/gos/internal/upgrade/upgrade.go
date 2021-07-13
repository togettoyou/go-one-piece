package upgrade

import (
	"fmt"
	"github.com/togettoyou/go-one-server/cmd/gos/internal/base"

	"github.com/spf13/cobra"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the gos tools",
	Long:  "Upgrade the gos tools. Example: gos upgrade",
	Run:   Run,
}

// Run upgrade the gos tools.
func Run(cmd *cobra.Command, args []string) {
	err := base.GoGet(
		"github.com/togettoyou/go-one-server/cmd/gos",
	)
	if err != nil {
		fmt.Println(err)
	}
}
