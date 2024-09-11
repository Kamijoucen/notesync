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

		path, err := cmd.Flags().GetString("f")
		if err != nil {
			Error(cmd, args, err)
		}
		fmt.Println(path)

		output, err := ExecuteCommand("notesync", "server", args...)
		if err != nil {
			Error(cmd, args, err)
		}
		_, _ = fmt.Fprint(os.Stdout, output)
	},
}

func init() {
	serverCmd.Flags().StringP("file", "f", "", "config file")
	rootCmd.AddCommand(serverCmd)
}
