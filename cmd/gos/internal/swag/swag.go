package swag

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/togettoyou/go-one-server/cmd/gos/internal/base"
	"os"
	"os/exec"
)

// CmdSwag generate documentation command.
var CmdSwag = &cobra.Command{
	Use:   "swag",
	Short: "generate documentation",
	Long:  "generate documentation. Example: gos swag",
	Run:   Swag,
}

// Swag generate documentation.
func Swag(cmd *cobra.Command, args []string) {
	_, err := exec.LookPath("swag")
	if err != nil {
		err := base.GoGet("github.com/swaggo/swag/cmd/swag")
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err.Error())
			return
		}
	}

	fd := exec.Command("swag", "init")
	fd.Stdout = os.Stdout
	fd.Stderr = os.Stderr
	if err := fd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err.Error())
		return
	}
	return
}
