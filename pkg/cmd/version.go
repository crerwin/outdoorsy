package cmd

import (
	"fmt"

	"github.com/crerwin/outdoorsy/pkg/outdoorsy"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the Outdoorsy version",
	Long:  "Displays the Outdoorsy version",
	Run:   versionRun,
}

func versionRun(cmd *cobra.Command, args []string) {
	fmt.Printf("Outdoorsy version %v\n", outdoorsy.Version)
}
