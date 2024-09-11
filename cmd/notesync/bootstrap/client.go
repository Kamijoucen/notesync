package bootstrap

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client is a tool to manage the client",

	Run: func(cmd *cobra.Command, args []string) {
		output, err := ExecuteCommand("notesync", "client", args...)
		if err != nil {
			Error(cmd, args, err)
		}
		_, _ = fmt.Fprint(os.Stdout, output)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
