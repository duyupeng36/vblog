package cmd

import (
	"vblog/cmd/start"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "vblog",
	Short: "vblog project back-end",
	Long: `vblog is a project cli
	
	vblog start -t file -f etc/config.toml
	`,
}

func init() {
	rootCmd.AddCommand(start.Cmd)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
