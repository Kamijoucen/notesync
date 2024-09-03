package bootstrap

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server is a tool to manage the server",

	Run: func(cmd *cobra.Command, args []string) {
		output, err := ExecuteCommand("notesync", "server1", args...)
		if err != nil {
			Error(cmd, args, err)
		}
		_, _ = fmt.Fprint(os.Stdout, output)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func execute() {

}
