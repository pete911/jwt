package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Version = "dev"

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "jwt cli version",
		Run:   versionCmdRun,
	}
)

func versionCmdRun(_ *cobra.Command, _ []string) {
	fmt.Println(Version)
}
