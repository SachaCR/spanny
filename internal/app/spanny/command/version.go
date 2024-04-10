package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print Spanny version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Spanny: v%s\n", version)
	},
}
