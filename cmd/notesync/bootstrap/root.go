package bootstrap

import (
	"errors"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "notesync",
	Short: "notesync is a tool to sync notes",
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unrecognized command"))
	},
}

func Execute() {
	_ = rootCmd.Execute()
}
