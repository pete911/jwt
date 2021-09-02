package cmd

import (
	"github.com/spf13/cobra"
)

var (
	silentFlag bool

	RootCmd = &cobra.Command{
		Use:   "jwt",
		Short: "jwt utilities",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			initLog(silentFlag)
		},
	}
)

func init() {
	RootCmd.PersistentFlags().BoolVar(&silentFlag, "silent", false, "suppress logs (warn, debug, ...)")
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(decodeCmd)
	RootCmd.AddCommand(encodeCmd)
}
