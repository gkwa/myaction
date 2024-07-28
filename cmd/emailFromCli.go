package cmd

import (
	"github.com/gkwa/myaction/core"
	"github.com/spf13/cobra"
)

var emailFromCliCmd = &cobra.Command{
	Use:     "email-from-cli [email]",
	Short:   "Process email from command line argument",
	Long:    `Fetch email from command line argument, validate it, and process it using Jsonnet.`,
	Aliases: []string{"efcli"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		source := &core.CliEmailSource{Email: args[0]}
		core.ProcessEmail(cmd.Context(), source)
	},
}

func init() {
	rootCmd.AddCommand(emailFromCliCmd)
}
