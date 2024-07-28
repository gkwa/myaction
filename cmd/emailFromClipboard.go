package cmd

import (
	"github.com/gkwa/myaction/core"
	"github.com/spf13/cobra"
)

var emailFromClipboardCmd = &cobra.Command{
	Use:     "email-from-clipboard",
	Short:   "Process email from clipboard",
	Long:    `Fetch email from clipboard, validate it, and process it using Jsonnet.`,
	Aliases: []string{"efclp"},
	Run: func(cmd *cobra.Command, args []string) {
		source := &core.ClipboardEmailSource{}
		core.ProcessEmail(cmd.Context(), source)
	},
}

func init() {
	rootCmd.AddCommand(emailFromClipboardCmd)
}
